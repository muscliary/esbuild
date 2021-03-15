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

type JSFeature uint64

const (
	ArraySpread JSFeature = 1 << iota
	Arrow
	AsyncAwait
	AsyncGenerator
	BigInt
	Class
	ClassField
	ClassPrivateAccessor
	ClassPrivateField
	ClassPrivateMethod
	ClassPrivateStaticAccessor
	ClassPrivateStaticField
	ClassPrivateStaticMethod
	ClassStaticField
	Const
	DefaultArgument
	Destructuring
	ExponentOperator
	ExportStarAs
	ForAwait
	ForOf
	Generator
	Hashbang
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

var jsTable = map[JSFeature]map[Engine][]int{
	ArraySpread: {
		Chrome:  {46},
		Edge:    {13},
		ES:      {2015},
		Firefox: {36},
		IOS:     {10},
		Node:    {5},
		Safari:  {10},
	},
	Arrow: {
		Chrome:  {49},
		Edge:    {13},
		ES:      {2015},
		Firefox: {45},
		IOS:     {10},
		Node:    {6},
		Safari:  {10},
	},
	AsyncAwait: {
		Chrome:  {55},
		Edge:    {15},
		ES:      {2017},
		Firefox: {52},
		IOS:     {11},
		Node:    {7, 6},
		Safari:  {11},
	},
	AsyncGenerator: {
		Chrome:  {63},
		Edge:    {79},
		ES:      {2018},
		Firefox: {57},
		IOS:     {12},
		Node:    {10, 0},
		Safari:  {12},
	},
	BigInt: {
		Chrome:  {67},
		Edge:    {79},
		ES:      {2020},
		Firefox: {68},
		IOS:     {14},
		Node:    {10, 4},
		Safari:  {14},
	},
	Class: {
		Chrome:  {49},
		Edge:    {13},
		ES:      {2015},
		Firefox: {45},
		IOS:     {10},
		Node:    {6},
		Safari:  {10},
	},
	ClassField: {
		Chrome:  {73},
		Edge:    {79},
		Firefox: {69},
		IOS:     {14},
		Node:    {12, 0},
		Safari:  {14},
	},
	ClassPrivateAccessor: {
		Chrome: {84},
		Edge:   {84},
		Node:   {14, 6},
	},
	ClassPrivateField: {
		Chrome: {84},
		Edge:   {84},
		Node:   {14, 6},
		Safari: {14, 1},
	},
	ClassPrivateMethod: {
		Chrome: {84},
		Edge:   {84},
		Node:   {14, 6},
	},
	ClassPrivateStaticAccessor: {
		Chrome: {84},
		Edge:   {84},
		Node:   {14, 6},
	},
	ClassPrivateStaticField: {
		Chrome: {74},
		Edge:   {79},
		Node:   {12, 0},
		Safari: {14, 1},
	},
	ClassPrivateStaticMethod: {
		Chrome: {84},
		Edge:   {84},
		Node:   {14, 6},
	},
	ClassStaticField: {
		Chrome:  {73},
		Edge:    {79},
		Firefox: {75},
		Node:    {12, 0},
		Safari:  {14, 1},
	},
	Const: {
		Chrome:  {49},
		Edge:    {14},
		ES:      {2015},
		Firefox: {51},
		IOS:     {11},
		Node:    {6},
		Safari:  {11},
	},
	DefaultArgument: {
		Chrome:  {49},
		Edge:    {14},
		ES:      {2015},
		Firefox: {53},
		IOS:     {10},
		Node:    {6},
		Safari:  {10},
	},
	Destructuring: {
		Chrome:  {51},
		Edge:    {18},
		ES:      {2015},
		Firefox: {53},
		IOS:     {10},
		Node:    {6, 5},
		Safari:  {10},
	},
	ExponentOperator: {
		Chrome:  {52},
		Edge:    {14},
		ES:      {2016},
		Firefox: {52},
		IOS:     {10, 3},
		Node:    {7},
		Safari:  {10, 1},
	},
	ExportStarAs: {
		Chrome:  {72},
		Edge:    {79},
		ES:      {2020},
		Firefox: {80},
		Node:    {12},
	},
	ForAwait: {
		Chrome:  {63},
		Edge:    {79},
		ES:      {2018},
		Firefox: {57},
		IOS:     {12},
		Node:    {10, 0},
		Safari:  {12},
	},
	ForOf: {
		Chrome:  {51},
		Edge:    {15},
		ES:      {2015},
		Firefox: {53},
		IOS:     {10},
		Node:    {6, 5},
		Safari:  {10},
	},
	Generator: {
		Chrome:  {50},
		Edge:    {13},
		ES:      {2015},
		Firefox: {53},
		IOS:     {10},
		Node:    {6},
		Safari:  {10},
	},
	Hashbang: {
		Chrome:  {74},
		Edge:    {79},
		Firefox: {67},
		IOS:     {13, 4},
		Node:    {12, 0},
		Safari:  {13, 1},
	},
	ImportMeta: {
		Chrome:  {64},
		Edge:    {79},
		ES:      {2020},
		Firefox: {62},
		IOS:     {12},
		Node:    {10, 4},
		Safari:  {11, 1},
	},
	Let: {
		Chrome:  {49},
		Edge:    {14},
		ES:      {2015},
		Firefox: {51},
		IOS:     {11},
		Node:    {6},
		Safari:  {11},
	},
	LogicalAssignment: {
		Chrome:  {85},
		Edge:    {85},
		Firefox: {79},
		IOS:     {14},
		Node:    {15, 0},
		Safari:  {14},
	},
	NestedRestBinding: {
		Chrome:  {49},
		Edge:    {14},
		ES:      {2016},
		Firefox: {47},
		IOS:     {10, 3},
		Node:    {6},
		Safari:  {10, 1},
	},
	NewTarget: {
		Chrome:  {46},
		Edge:    {14},
		ES:      {2015},
		Firefox: {41},
		IOS:     {10},
		Node:    {5},
		Safari:  {10},
	},
	NullishCoalescing: {
		Chrome:  {80},
		Edge:    {80},
		ES:      {2020},
		Firefox: {72},
		IOS:     {13, 4},
		Node:    {14, 0},
		Safari:  {13, 1},
	},
	ObjectAccessors: {
		Chrome:  {5},
		Edge:    {12},
		ES:      {5},
		Firefox: {2},
		IOS:     {6},
		Node:    {0, 10},
		Safari:  {3, 1},
	},
	ObjectExtensions: {
		Chrome:  {44},
		Edge:    {12},
		ES:      {2015},
		Firefox: {34},
		IOS:     {10},
		Node:    {4},
		Safari:  {10},
	},
	ObjectRestSpread: {
		ES:      {2018},
		Firefox: {55},
		IOS:     {11, 3},
		Safari:  {11, 1},
	},
	OptionalCatchBinding: {
		Chrome:  {66},
		Edge:    {79},
		ES:      {2019},
		Firefox: {58},
		IOS:     {11, 3},
		Node:    {10, 0},
		Safari:  {11, 1},
	},
	OptionalChain: {
		Chrome:  {80},
		Edge:    {80},
		ES:      {2020},
		Firefox: {74},
		IOS:     {13, 4},
		Node:    {14, 0},
		Safari:  {13, 1},
	},
	RestArgument: {
		Chrome:  {47},
		Edge:    {12},
		ES:      {2015},
		Firefox: {43},
		IOS:     {10},
		Node:    {6},
		Safari:  {10},
	},
	TemplateLiteral: {
		Chrome:  {41},
		Edge:    {13},
		ES:      {2015},
		Firefox: {34},
		IOS:     {9},
		Node:    {4},
		Safari:  {9},
	},
	TopLevelAwait: {},
	UnicodeEscapes: {
		Chrome:  {44},
		Edge:    {12},
		ES:      {2015},
		Firefox: {53},
		IOS:     {9},
		Node:    {4},
		Safari:  {9},
	},
}

func isVersionLessThan(a []int, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] > b[i] {
			return false
		}
		if a[i] < b[i] {
			return true
		}
	}
	return len(a) < len(b)
}

// Return all features that are not available in at least one environment
func UnsupportedJSFeatures(constraints map[Engine][]int) (unsupported JSFeature) {
	for feature, engines := range jsTable {
		for engine, version := range constraints {
			if minVersion, ok := engines[engine]; !ok || isVersionLessThan(version, minVersion) {
				unsupported |= feature
			}
		}
	}
	return
}
