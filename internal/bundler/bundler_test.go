package bundler

// Bundling test results are stored in snapshot files, located in the
// "snapshots" directory. This allows test results to be updated easily without
// manually rewriting all of the expected values. To update the tests run
// "UPDATE_SNAPSHOTS=1 make test" and commit the updated values. Make sure to
// inspect the diff to ensure the expected values are valid.

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/evanw/esbuild/internal/compat"
	"github.com/evanw/esbuild/internal/config"
	"github.com/evanw/esbuild/internal/fs"
	"github.com/evanw/esbuild/internal/logging"
	"github.com/evanw/esbuild/internal/resolver"
	"github.com/kylelemons/godebug/diff"
)

func es(version int) compat.Feature {
	return compat.UnsupportedFeatures(map[compat.Engine][]int{
		compat.ES: {version},
	})
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		stringA := fmt.Sprintf("%v", a)
		stringB := fmt.Sprintf("%v", b)
		if strings.Contains(stringA, "\n") {
			t.Fatal(diff.Diff(stringB, stringA))
		} else {
			t.Fatalf("%s != %s", a, b)
		}
	}
}

func assertLog(t *testing.T, msgs []logging.Msg, expected string) {
	text := ""
	for _, msg := range msgs {
		text += msg.String(logging.StderrOptions{}, logging.TerminalInfo{})
	}
	assertEqual(t, text, expected)
}

func hasErrors(msgs []logging.Msg) bool {
	for _, msg := range msgs {
		if msg.Kind == logging.Error {
			return true
		}
	}
	return false
}

type bundled struct {
	files              map[string]string
	entryPaths         []string
	expectedScanLog    string
	expectedCompileLog string
	options            config.Options
}

type suite struct {
	name               string
	path               string
	mutex              sync.Mutex
	expectedSnapshots  map[string]string
	generatedSnapshots map[string]string
}

func (s *suite) expectBundled(t *testing.T, args bundled) {
	testName := t.Name()
	t.Run("", func(t *testing.T) {
		fs := fs.MockFS(args.files)
		args.options.ExtensionOrder = []string{".tsx", ".ts", ".jsx", ".js", ".json"}
		if args.options.AbsOutputFile != "" {
			args.options.AbsOutputDir = path.Dir(args.options.AbsOutputFile)
		}
		log := logging.NewDeferLog()
		resolver := resolver.NewResolver(fs, log, args.options)
		bundle := ScanBundle(log, fs, resolver, args.entryPaths, args.options)
		msgs := log.Done()
		assertLog(t, msgs, args.expectedScanLog)

		// Stop now if there were any errors during the scan
		if hasErrors(msgs) {
			return
		}

		log = logging.NewDeferLog()
		args.options.OmitRuntimeForTests = true
		results := bundle.Compile(log, args.options)
		msgs = log.Done()
		assertLog(t, msgs, args.expectedCompileLog)

		// Stop now if there were any errors during the compile
		if hasErrors(msgs) {
			return
		}

		// Don't include source maps in results since they are just noise. Source
		// map validity is tested separately in a test that uses Mozilla's source
		// map parsing library.
		generated := ""
		for _, result := range results {
			if !strings.HasSuffix(result.AbsPath, ".map") {
				if generated != "" {
					generated += "\n"
				}
				generated += fmt.Sprintf("---------- %s ----------\n%s", result.AbsPath, string(result.Contents))
			}
		}
		s.compareSnapshot(t, testName, generated)
	})
}

const snapshotsDir = "snapshots"
const snapshotSplitter = "\n================================================================================\n"

var globalTestMutex sync.Mutex
var globalSuites map[*suite]bool
var globalUpdateSnapshots bool

func (s *suite) compareSnapshot(t *testing.T, testName string, generated string) {
	// Initialize the test suite during the first test
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.path == "" {
		s.path = snapshotsDir + "/snapshots_" + s.name + ".txt"
		s.generatedSnapshots = make(map[string]string)
		s.expectedSnapshots = make(map[string]string)
		if contents, err := ioutil.ReadFile(s.path); err == nil {
			for _, part := range strings.Split(string(contents), snapshotSplitter) {
				if newline := strings.IndexByte(part, '\n'); newline != -1 {
					key := part[:newline]
					value := part[newline+1:]
					s.expectedSnapshots[key] = value
				} else {
					s.expectedSnapshots[part] = ""
				}
			}
		}
		globalTestMutex.Lock()
		defer globalTestMutex.Unlock()
		if globalSuites == nil {
			globalSuites = make(map[*suite]bool)
		}
		globalSuites[s] = true
		_, globalUpdateSnapshots = os.LookupEnv("UPDATE_SNAPSHOTS")
	}

	// Check against the stored snapshot if present
	s.generatedSnapshots[testName] = generated
	if !globalUpdateSnapshots {
		if expected, ok := s.expectedSnapshots[testName]; ok {
			assertEqual(t, generated, expected)
		} else {
			t.Fatalf("No snapshot saved for %s\n%s", testName, generated)
		}
	}
}

func (s *suite) updateSnapshots() {
	os.Mkdir(snapshotsDir, 0755)
	keys := make([]string, 0, len(s.generatedSnapshots))
	for key := range s.generatedSnapshots {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	contents := ""
	for i, key := range keys {
		if i > 0 {
			contents += snapshotSplitter
		}
		contents += fmt.Sprintf("%s\n%s", key, s.generatedSnapshots[key])
	}
	if err := ioutil.WriteFile(s.path, []byte(contents), 0644); err != nil {
		panic(err)
	}
}

func (s *suite) validateSnapshots() bool {
	isValid := true
	for key := range s.expectedSnapshots {
		if _, ok := s.generatedSnapshots[key]; !ok {
			if isValid {
				fmt.Printf("%s\n", s.path)
			}
			fmt.Printf("    No test found for snapshot %s\n", key)
			isValid = false
		}
	}
	return isValid
}

func TestMain(m *testing.M) {
	code := m.Run()
	if globalSuites != nil {
		if globalUpdateSnapshots {
			for s := range globalSuites {
				s.updateSnapshots()
			}
		} else {
			for s := range globalSuites {
				if !s.validateSnapshots() {
					code = 1
				}
			}
		}
	}
	os.Exit(code)
}
