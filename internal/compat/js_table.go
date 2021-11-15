// This file was automatically generated by "compat-table.js"

package compat

type Engine uint8

const (
	Chrome Engine = iota
	Edge
	ES
	Firefox
	IOS
	Node
	Safari
)

func (e Engine) String() string {
	switch e {
	case Chrome:
		return "chrome"
	case Edge:
		return "edge"
	case ES:
		return "es"
	case Firefox:
		return "firefox"
	case IOS:
		return "ios"
	case Node:
		return "node"
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
	BigInt
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
	Const
	DefaultArgument
	Destructuring
	DynamicImport
	ExponentOperator
	ExportStarAs
	ForAwait
	ForOf
	Generator
	Hashbang
	ImportAssertions
	ImportMeta
	Let
	LogicalAssignment
	NestedRestBinding
	NewTarget
	NullishCoalescing
	ObjectAccessors
	ObjectExtensions
	ObjectRestSpread
	OptionalCatchBinding
	OptionalChain
	RestArgument
	TemplateLiteral
	TopLevelAwait
	UnicodeEscapes
)

func (features JSFeature) Has(feature JSFeature) bool {
	return (features & feature) != 0
}

var jsTable = map[JSFeature]map[Engine]v{
	ArbitraryModuleNamespaceNames: {
		Chrome:  v{90, 0, 0},
		Firefox: v{87, 0, 0},
		Node:    v{16, 0, 0},
	},
	ArraySpread: {
		Chrome:  v{46, 0, 0},
		Edge:    v{13, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{36, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{5, 0, 0},
		Safari:  v{10, 0, 0},
	},
	Arrow: {
		Chrome:  v{49, 0, 0},
		Edge:    v{13, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{45, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 0, 0},
	},
	AsyncAwait: {
		Chrome:  v{55, 0, 0},
		Edge:    v{15, 0, 0},
		ES:      v{2017, 0, 0},
		Firefox: v{52, 0, 0},
		IOS:     v{11, 0, 0},
		Node:    v{7, 6, 0},
		Safari:  v{11, 0, 0},
	},
	AsyncGenerator: {
		Chrome:  v{63, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2018, 0, 0},
		Firefox: v{57, 0, 0},
		IOS:     v{12, 0, 0},
		Node:    v{10, 0, 0},
		Safari:  v{12, 0, 0},
	},
	BigInt: {
		Chrome:  v{67, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2020, 0, 0},
		Firefox: v{68, 0, 0},
		IOS:     v{14, 0, 0},
		Node:    v{10, 4, 0},
		Safari:  v{14, 0, 0},
	},
	Class: {
		Chrome:  v{49, 0, 0},
		Edge:    v{13, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{45, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 0, 0},
	},
	ClassField: {
		Chrome:  v{73, 0, 0},
		Edge:    v{79, 0, 0},
		Firefox: v{69, 0, 0},
		IOS:     v{14, 0, 0},
		Node:    v{12, 0, 0},
		Safari:  v{14, 0, 0},
	},
	ClassPrivateAccessor: {
		Chrome:  v{84, 0, 0},
		Edge:    v{84, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{14, 6, 0},
		Safari:  v{15, 0, 0},
	},
	ClassPrivateBrandCheck: {
		Chrome:  v{91, 0, 0},
		Edge:    v{91, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{16, 9, 0},
		Safari:  v{15, 0, 0},
	},
	ClassPrivateField: {
		Chrome:  v{84, 0, 0},
		Edge:    v{84, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{14, 6, 0},
		Safari:  v{14, 1, 0},
	},
	ClassPrivateMethod: {
		Chrome:  v{84, 0, 0},
		Edge:    v{84, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{14, 6, 0},
		Safari:  v{15, 0, 0},
	},
	ClassPrivateStaticAccessor: {
		Chrome:  v{84, 0, 0},
		Edge:    v{84, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{14, 6, 0},
		Safari:  v{15, 0, 0},
	},
	ClassPrivateStaticField: {
		Chrome:  v{74, 0, 0},
		Edge:    v{79, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{12, 0, 0},
		Safari:  v{14, 1, 0},
	},
	ClassPrivateStaticMethod: {
		Chrome:  v{84, 0, 0},
		Edge:    v{84, 0, 0},
		Firefox: v{90, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{14, 6, 0},
		Safari:  v{15, 0, 0},
	},
	ClassStaticBlocks: {
		Chrome: v{91, 0, 0},
		Node:   v{16, 11, 0},
	},
	ClassStaticField: {
		Chrome:  v{73, 0, 0},
		Edge:    v{79, 0, 0},
		Firefox: v{75, 0, 0},
		IOS:     v{15, 0, 0},
		Node:    v{12, 0, 0},
		Safari:  v{14, 1, 0},
	},
	Const: {
		Chrome:  v{49, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{51, 0, 0},
		IOS:     v{11, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{11, 0, 0},
	},
	DefaultArgument: {
		Chrome:  v{49, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{53, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 0, 0},
	},
	Destructuring: {
		Chrome:  v{51, 0, 0},
		Edge:    v{18, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{53, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 5, 0},
		Safari:  v{10, 0, 0},
	},
	DynamicImport: {
		Chrome:  v{63, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{67, 0, 0},
		IOS:     v{11, 0, 0},
		Node:    v{13, 2, 0},
		Safari:  v{11, 1, 0},
	},
	ExponentOperator: {
		Chrome:  v{52, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2016, 0, 0},
		Firefox: v{52, 0, 0},
		IOS:     v{10, 3, 0},
		Node:    v{7, 0, 0},
		Safari:  v{10, 1, 0},
	},
	ExportStarAs: {
		Chrome:  v{72, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2020, 0, 0},
		Firefox: v{80, 0, 0},
		Node:    v{12, 0, 0},
	},
	ForAwait: {
		Chrome:  v{63, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2018, 0, 0},
		Firefox: v{57, 0, 0},
		IOS:     v{12, 0, 0},
		Node:    v{10, 0, 0},
		Safari:  v{12, 0, 0},
	},
	ForOf: {
		Chrome:  v{51, 0, 0},
		Edge:    v{15, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{53, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 5, 0},
		Safari:  v{10, 0, 0},
	},
	Generator: {
		Chrome:  v{50, 0, 0},
		Edge:    v{13, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{53, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 0, 0},
	},
	Hashbang: {
		Chrome:  v{74, 0, 0},
		Edge:    v{79, 0, 0},
		Firefox: v{67, 0, 0},
		IOS:     v{13, 4, 0},
		Node:    v{12, 0, 0},
		Safari:  v{13, 1, 0},
	},
	ImportAssertions: {
		Chrome: v{91, 0, 0},
	},
	ImportMeta: {
		Chrome:  v{64, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2020, 0, 0},
		Firefox: v{62, 0, 0},
		IOS:     v{12, 0, 0},
		Node:    v{10, 4, 0},
		Safari:  v{11, 1, 0},
	},
	Let: {
		Chrome:  v{49, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{51, 0, 0},
		IOS:     v{11, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{11, 0, 0},
	},
	LogicalAssignment: {
		Chrome:  v{85, 0, 0},
		Edge:    v{85, 0, 0},
		ES:      v{2021, 0, 0},
		Firefox: v{79, 0, 0},
		IOS:     v{14, 0, 0},
		Node:    v{15, 0, 0},
		Safari:  v{14, 0, 0},
	},
	NestedRestBinding: {
		Chrome:  v{49, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2016, 0, 0},
		Firefox: v{47, 0, 0},
		IOS:     v{10, 3, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 1, 0},
	},
	NewTarget: {
		Chrome:  v{46, 0, 0},
		Edge:    v{14, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{41, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{5, 0, 0},
		Safari:  v{10, 0, 0},
	},
	NullishCoalescing: {
		Chrome:  v{80, 0, 0},
		Edge:    v{80, 0, 0},
		ES:      v{2020, 0, 0},
		Firefox: v{72, 0, 0},
		IOS:     v{13, 4, 0},
		Node:    v{14, 0, 0},
		Safari:  v{13, 1, 0},
	},
	ObjectAccessors: {
		Chrome:  v{5, 0, 0},
		Edge:    v{12, 0, 0},
		ES:      v{5, 0, 0},
		Firefox: v{2, 0, 0},
		IOS:     v{6, 0, 0},
		Node:    v{0, 10, 0},
		Safari:  v{3, 1, 0},
	},
	ObjectExtensions: {
		Chrome:  v{44, 0, 0},
		Edge:    v{12, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{34, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{4, 0, 0},
		Safari:  v{10, 0, 0},
	},
	ObjectRestSpread: {
		ES:      v{2018, 0, 0},
		Firefox: v{55, 0, 0},
		IOS:     v{11, 3, 0},
		Safari:  v{11, 1, 0},
	},
	OptionalCatchBinding: {
		Chrome:  v{66, 0, 0},
		Edge:    v{79, 0, 0},
		ES:      v{2019, 0, 0},
		Firefox: v{58, 0, 0},
		IOS:     v{11, 3, 0},
		Node:    v{10, 0, 0},
		Safari:  v{11, 1, 0},
	},
	OptionalChain: {
		Chrome:  v{91, 0, 0},
		Edge:    v{91, 0, 0},
		ES:      v{2020, 0, 0},
		Firefox: v{74, 0, 0},
		IOS:     v{13, 4, 0},
		Node:    v{16, 9, 0},
		Safari:  v{13, 1, 0},
	},
	RestArgument: {
		Chrome:  v{47, 0, 0},
		Edge:    v{12, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{43, 0, 0},
		IOS:     v{10, 0, 0},
		Node:    v{6, 0, 0},
		Safari:  v{10, 0, 0},
	},
	TemplateLiteral: {
		Chrome:  v{41, 0, 0},
		Edge:    v{13, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{34, 0, 0},
		IOS:     v{9, 0, 0},
		Node:    v{4, 0, 0},
		Safari:  v{9, 0, 0},
	},
	TopLevelAwait: {
		Chrome:  v{89, 0, 0},
		Edge:    v{89, 0, 0},
		Firefox: v{89, 0, 0},
		Node:    v{14, 8, 0},
		Safari:  v{15, 0, 0},
	},
	UnicodeEscapes: {
		Chrome:  v{44, 0, 0},
		Edge:    v{12, 0, 0},
		ES:      v{2015, 0, 0},
		Firefox: v{53, 0, 0},
		IOS:     v{9, 0, 0},
		Node:    v{4, 0, 0},
		Safari:  v{9, 0, 0},
	},
}

// Return all features that are not available in at least one environment
func UnsupportedJSFeatures(constraints map[Engine][]int) (unsupported JSFeature) {
	for feature, engines := range jsTable {
		for engine, version := range constraints {
			if minVersion, ok := engines[engine]; !ok || compareVersions(minVersion, version) > 0 {
				unsupported |= feature
			}
		}
	}
	return
}
