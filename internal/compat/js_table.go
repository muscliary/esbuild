// This file was automatically generated by "compat-table.js"

package compat

type Engine uint8

const (
	Chrome Engine = iota
	Deno
	Edge
	ES
	Firefox
	Hermes
	IE
	IOS
	Node
	Opera
	Rhino
	Safari
)

func (e Engine) String() string {
	switch e {
	case Chrome:
		return "chrome"
	case Deno:
		return "deno"
	case Edge:
		return "edge"
	case ES:
		return "es"
	case Firefox:
		return "firefox"
	case Hermes:
		return "hermes"
	case IE:
		return "ie"
	case IOS:
		return "ios"
	case Node:
		return "node"
	case Opera:
		return "opera"
	case Rhino:
		return "rhino"
	case Safari:
		return "safari"
	}
	return ""
}

type JSFeature uint64

const (
	ArbitraryModuleNamespaceNames JSFeature = 1 << iota
	ArraySpread
	Arrow
	AsyncAwait
	AsyncGenerator
	Bigint
	Class
	ClassField
	ClassPrivateAccessor
	ClassPrivateBrandCheck
	ClassPrivateField
	ClassPrivateMethod
	ClassPrivateStaticAccessor
	ClassPrivateStaticField
	ClassPrivateStaticMethod
	ClassStaticBlocks
	ClassStaticField
	ConstAndLet
	Decorators
	DefaultArgument
	Destructuring
	DynamicImport
	ExponentOperator
	ExportStarAs
	ForAwait
	ForOf
	FunctionOrClassPropertyAccess
	Generator
	Hashbang
	ImportAssertions
	ImportMeta
	InlineScript
	LogicalAssignment
	NestedRestBinding
	NewTarget
	NodeColonPrefixImport
	NodeColonPrefixRequire
	NullishCoalescing
	ObjectAccessors
	ObjectExtensions
	ObjectRestSpread
	OptionalCatchBinding
	OptionalChain
	RegexpDotAllFlag
	RegexpLookbehindAssertions
	RegexpMatchIndices
	RegexpNamedCaptureGroups
	RegexpSetNotation
	RegexpStickyAndUnicodeFlags
	RegexpUnicodePropertyEscapes
	RestArgument
	TemplateLiteral
	TopLevelAwait
	TypeofExoticObjectIsObject
	UnicodeEscapes
)

var StringToJSFeature = map[string]JSFeature{
	"arbitrary-module-namespace-names":  ArbitraryModuleNamespaceNames,
	"array-spread":                      ArraySpread,
	"arrow":                             Arrow,
	"async-await":                       AsyncAwait,
	"async-generator":                   AsyncGenerator,
	"bigint":                            Bigint,
	"class":                             Class,
	"class-field":                       ClassField,
	"class-private-accessor":            ClassPrivateAccessor,
	"class-private-brand-check":         ClassPrivateBrandCheck,
	"class-private-field":               ClassPrivateField,
	"class-private-method":              ClassPrivateMethod,
	"class-private-static-accessor":     ClassPrivateStaticAccessor,
	"class-private-static-field":        ClassPrivateStaticField,
	"class-private-static-method":       ClassPrivateStaticMethod,
	"class-static-blocks":               ClassStaticBlocks,
	"class-static-field":                ClassStaticField,
	"const-and-let":                     ConstAndLet,
	"decorators":                        Decorators,
	"default-argument":                  DefaultArgument,
	"destructuring":                     Destructuring,
	"dynamic-import":                    DynamicImport,
	"exponent-operator":                 ExponentOperator,
	"export-star-as":                    ExportStarAs,
	"for-await":                         ForAwait,
	"for-of":                            ForOf,
	"function-or-class-property-access": FunctionOrClassPropertyAccess,
	"generator":                         Generator,
	"hashbang":                          Hashbang,
	"import-assertions":                 ImportAssertions,
	"import-meta":                       ImportMeta,
	"inline-script":                     InlineScript,
	"logical-assignment":                LogicalAssignment,
	"nested-rest-binding":               NestedRestBinding,
	"new-target":                        NewTarget,
	"node-colon-prefix-import":          NodeColonPrefixImport,
	"node-colon-prefix-require":         NodeColonPrefixRequire,
	"nullish-coalescing":                NullishCoalescing,
	"object-accessors":                  ObjectAccessors,
	"object-extensions":                 ObjectExtensions,
	"object-rest-spread":                ObjectRestSpread,
	"optional-catch-binding":            OptionalCatchBinding,
	"optional-chain":                    OptionalChain,
	"regexp-dot-all-flag":               RegexpDotAllFlag,
	"regexp-lookbehind-assertions":      RegexpLookbehindAssertions,
	"regexp-match-indices":              RegexpMatchIndices,
	"regexp-named-capture-groups":       RegexpNamedCaptureGroups,
	"regexp-set-notation":               RegexpSetNotation,
	"regexp-sticky-and-unicode-flags":   RegexpStickyAndUnicodeFlags,
	"regexp-unicode-property-escapes":   RegexpUnicodePropertyEscapes,
	"rest-argument":                     RestArgument,
	"template-literal":                  TemplateLiteral,
	"top-level-await":                   TopLevelAwait,
	"typeof-exotic-object-is-object":    TypeofExoticObjectIsObject,
	"unicode-escapes":                   UnicodeEscapes,
}

func (features JSFeature) Has(feature JSFeature) bool {
	return (features & feature) != 0
}

func (features JSFeature) ApplyOverrides(overrides JSFeature, mask JSFeature) JSFeature {
	return (features & ^mask) | (overrides & mask)
}

var jsTable = map[JSFeature]map[Engine][]versionRange{
	ArbitraryModuleNamespaceNames: {
		Chrome:  {{start: v{90, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{87, 0, 0}}},
		Node:    {{start: v{16, 0, 0}}},
	},
	ArraySpread: {
		Chrome:  {{start: v{46, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{36, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Arrow: {
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	AsyncAwait: {
		Chrome:  {{start: v{55, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2017, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{7, 6, 0}}},
		Opera:   {{start: v{42, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	AsyncGenerator: {
		Chrome:  {{start: v{63, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	Bigint: {
		Chrome:  {{start: v{67, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{68, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Opera:   {{start: v{54, 0, 0}}},
		Rhino:   {{start: v{1, 7, 14}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	Class: {
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{45, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ClassField: {
		Chrome:  {{start: v{73, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{69, 0, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	ClassPrivateAccessor: {
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateBrandCheck: {
		Chrome:  {{start: v{91, 0, 0}}},
		Deno:    {{start: v{1, 9, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{16, 4, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateField: {
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateMethod: {
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticAccessor: {
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassPrivateStaticField: {
		Chrome:  {{start: v{74, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ClassPrivateStaticMethod: {
		Chrome:  {{start: v{84, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{84, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{90, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 6, 0}}},
		Opera:   {{start: v{70, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	ClassStaticBlocks: {
		Chrome:  {{start: v{91, 0, 0}}},
		Edge:    {{start: v{94, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{93, 0, 0}}},
		Node:    {{start: v{16, 11, 0}}},
		Opera:   {{start: v{80, 0, 0}}},
	},
	ClassStaticField: {
		Chrome:  {{start: v{73, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{75, 0, 0}}},
		IOS:     {{start: v{14, 5, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
		Safari:  {{start: v{14, 1, 0}}},
	},
	ConstAndLet: {
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{51, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{11, 0, 0}}},
	},
	Decorators: {},
	DefaultArgument: {
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Destructuring: {
		Chrome:  {{start: v{51, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{18, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	DynamicImport: {
		Chrome:  {{start: v{63, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		IOS:     {{start: v{11, 0, 0}}},
		Node:    {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{13, 2, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	ExponentOperator: {
		Chrome:  {{start: v{52, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{52, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{7, 0, 0}}},
		Opera:   {{start: v{39, 0, 0}}},
		Rhino:   {{start: v{1, 7, 14}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	ExportStarAs: {
		Chrome:  {{start: v{72, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{80, 0, 0}}},
		Node:    {{start: v{12, 0, 0}}},
		Opera:   {{start: v{60, 0, 0}}},
	},
	ForAwait: {
		Chrome:  {{start: v{63, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{57, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{50, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	ForOf: {
		Chrome:  {{start: v{51, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{15, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 5, 0}}},
		Opera:   {{start: v{38, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	FunctionOrClassPropertyAccess: {
		Chrome:  {{start: v{0, 0, 0}}},
		Deno:    {{start: v{0, 0, 0}}},
		Edge:    {{start: v{0, 0, 0}}},
		ES:      {{start: v{0, 0, 0}}},
		Firefox: {{start: v{0, 0, 0}}},
		Hermes:  {{start: v{0, 0, 0}}},
		IE:      {{start: v{0, 0, 0}}},
		IOS:     {{start: v{0, 0, 0}}},
		Node:    {{start: v{0, 0, 0}}},
		Opera:   {{start: v{0, 0, 0}}},
		Rhino:   {{start: v{0, 0, 0}}},
		Safari:  {{start: v{16, 3, 0}}},
	},
	Generator: {
		Chrome:  {{start: v{50, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{37, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	Hashbang: {
		Chrome:  {{start: v{74, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		Firefox: {{start: v{67, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{12, 5, 0}}},
		Opera:   {{start: v{62, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ImportAssertions: {
		Chrome: {{start: v{91, 0, 0}}},
		Node:   {{start: v{16, 14, 0}}},
	},
	ImportMeta: {
		Chrome:  {{start: v{64, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{62, 0, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{10, 4, 0}}},
		Opera:   {{start: v{51, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	InlineScript: {},
	LogicalAssignment: {
		Chrome:  {{start: v{85, 0, 0}}},
		Deno:    {{start: v{1, 2, 0}}},
		Edge:    {{start: v{85, 0, 0}}},
		ES:      {{start: v{2021, 0, 0}}},
		Firefox: {{start: v{79, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{14, 0, 0}}},
		Node:    {{start: v{15, 0, 0}}},
		Opera:   {{start: v{71, 0, 0}}},
		Safari:  {{start: v{14, 0, 0}}},
	},
	NestedRestBinding: {
		Chrome:  {{start: v{49, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2016, 0, 0}}},
		Firefox: {{start: v{47, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 3, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{36, 0, 0}}},
		Safari:  {{start: v{10, 1, 0}}},
	},
	NewTarget: {
		Chrome:  {{start: v{46, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{14, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{41, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{5, 0, 0}}},
		Opera:   {{start: v{33, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	NodeColonPrefixImport: {
		Node: {{start: v{12, 20, 0}, end: v{13, 0, 0}}, {start: v{14, 13, 1}}},
	},
	NodeColonPrefixRequire: {
		Node: {{start: v{14, 18, 0}, end: v{15, 0, 0}}, {start: v{16, 0, 0}}},
	},
	NullishCoalescing: {
		Chrome:  {{start: v{80, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{80, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{72, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{14, 0, 0}}},
		Opera:   {{start: v{67, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	ObjectAccessors: {
		Chrome:  {{start: v{5, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{5, 0, 0}}},
		Firefox: {{start: v{2, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IE:      {{start: v{9, 0, 0}}},
		IOS:     {{start: v{6, 0, 0}}},
		Node:    {{start: v{0, 4, 0}}},
		Opera:   {{start: v{10, 10, 0}}},
		Rhino:   {{start: v{1, 7, 13}}},
		Safari:  {{start: v{3, 1, 0}}},
	},
	ObjectExtensions: {
		Chrome:  {{start: v{44, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	ObjectRestSpread: {
		Chrome:  {{start: v{60, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{55, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{8, 3, 0}}},
		Opera:   {{start: v{47, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalCatchBinding: {
		Chrome:  {{start: v{66, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2019, 0, 0}}},
		Firefox: {{start: v{58, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{53, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	OptionalChain: {
		Chrome:  {{start: v{91, 0, 0}}},
		Deno:    {{start: v{1, 9, 0}}},
		Edge:    {{start: v{91, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{74, 0, 0}}},
		Hermes:  {{start: v{0, 12, 0}}},
		IOS:     {{start: v{13, 4, 0}}},
		Node:    {{start: v{16, 1, 0}}},
		Opera:   {{start: v{77, 0, 0}}},
		Safari:  {{start: v{13, 1, 0}}},
	},
	RegexpDotAllFlag: {
		Chrome:  {{start: v{62, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{8, 10, 0}}},
		Opera:   {{start: v{49, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	RegexpLookbehindAssertions: {
		Chrome:  {{start: v{62, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		Node:    {{start: v{8, 10, 0}}},
		Opera:   {{start: v{49, 0, 0}}},
	},
	RegexpMatchIndices: {
		Chrome:  {{start: v{90, 0, 0}}},
		Edge:    {{start: v{90, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{88, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Opera:   {{start: v{76, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	RegexpNamedCaptureGroups: {
		Chrome:  {{start: v{64, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{51, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	RegexpSetNotation: {},
	RegexpStickyAndUnicodeFlags: {
		Chrome:  {{start: v{50, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{46, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{12, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{37, 0, 0}}},
		Safari:  {{start: v{12, 0, 0}}},
	},
	RegexpUnicodePropertyEscapes: {
		Chrome:  {{start: v{64, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{79, 0, 0}}},
		ES:      {{start: v{2018, 0, 0}}},
		Firefox: {{start: v{78, 0, 0}}},
		IOS:     {{start: v{11, 3, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{51, 0, 0}}},
		Safari:  {{start: v{11, 1, 0}}},
	},
	RestArgument: {
		Chrome:  {{start: v{47, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{43, 0, 0}}},
		IOS:     {{start: v{10, 0, 0}}},
		Node:    {{start: v{6, 0, 0}}},
		Opera:   {{start: v{34, 0, 0}}},
		Safari:  {{start: v{10, 0, 0}}},
	},
	TemplateLiteral: {
		Chrome:  {{start: v{41, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{13, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{34, 0, 0}}},
		IOS:     {{start: v{9, 0, 0}}},
		Node:    {{start: v{10, 0, 0}}},
		Opera:   {{start: v{28, 0, 0}}},
		Safari:  {{start: v{9, 0, 0}}},
	},
	TopLevelAwait: {
		Chrome:  {{start: v{89, 0, 0}}},
		Edge:    {{start: v{89, 0, 0}}},
		ES:      {{start: v{2022, 0, 0}}},
		Firefox: {{start: v{89, 0, 0}}},
		IOS:     {{start: v{15, 0, 0}}},
		Node:    {{start: v{14, 8, 0}}},
		Opera:   {{start: v{75, 0, 0}}},
		Safari:  {{start: v{15, 0, 0}}},
	},
	TypeofExoticObjectIsObject: {
		Chrome:  {{start: v{0, 0, 0}}},
		Deno:    {{start: v{0, 0, 0}}},
		Edge:    {{start: v{0, 0, 0}}},
		ES:      {{start: v{2020, 0, 0}}},
		Firefox: {{start: v{0, 0, 0}}},
		Hermes:  {{start: v{0, 0, 0}}},
		IOS:     {{start: v{0, 0, 0}}},
		Node:    {{start: v{0, 0, 0}}},
		Opera:   {{start: v{0, 0, 0}}},
		Rhino:   {{start: v{0, 0, 0}}},
		Safari:  {{start: v{0, 0, 0}}},
	},
	UnicodeEscapes: {
		Chrome:  {{start: v{44, 0, 0}}},
		Deno:    {{start: v{1, 0, 0}}},
		Edge:    {{start: v{12, 0, 0}}},
		ES:      {{start: v{2015, 0, 0}}},
		Firefox: {{start: v{53, 0, 0}}},
		Hermes:  {{start: v{0, 7, 0}}},
		IOS:     {{start: v{9, 0, 0}}},
		Node:    {{start: v{4, 0, 0}}},
		Opera:   {{start: v{31, 0, 0}}},
		Safari:  {{start: v{9, 0, 0}}},
	},
}

// Return all features that are not available in at least one environment
func UnsupportedJSFeatures(constraints map[Engine][]int) (unsupported JSFeature) {
	for feature, engines := range jsTable {
		if feature == InlineScript {
			continue // This is purely user-specified
		}
		for engine, version := range constraints {
			if versionRanges, ok := engines[engine]; !ok || !isVersionSupported(versionRanges, version) {
				unsupported |= feature
			}
		}
	}
	return
}
