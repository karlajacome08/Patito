// Code generated from grammar/Patito.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Patito

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PatitoParser struct {
	*antlr.BaseParser
}

var PatitoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func patitoParserInit() {
	staticData := &PatitoParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "':'", "','", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='",
		"'=='", "'!='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'!'", "'programa'",
		"'vars'", "'inicio'", "'fin'", "'si'", "'sino'", "'mientras'", "'haz'",
		"'escribe'", "'nula'", "'entero'", "'flotante'", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "PROGRAMA", "VARS", "INICIO", "FIN", "SI", "SINO", "MIENTRAS",
		"HAZ", "ESCRIBE", "NULA", "ENTERO", "FLOTANTE", "RETURN", "ID", "INT",
		"FLOAT", "STRING", "WS", "LINE_COMMENT", "BODY_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "vars", "varDecl", "idList", "type_", "funcDef", "params",
		"param", "body", "stmt", "returnStmt", "assign", "print", "printItemList",
		"printItem", "ifStmt", "ifMarkThen", "elseMark", "whileStmt", "whileStart",
		"whileCond", "call", "expr", "relop", "sum", "prod", "unary", "primary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 267, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 61, 8, 0, 1, 0, 5, 0,
		64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 5, 1, 76, 8, 1, 10, 1, 12, 1, 79, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 5, 3, 89, 8, 3, 10, 3, 12, 3, 92, 9, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 3, 5, 98, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 103, 8, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 108, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6,
		117, 8, 6, 10, 6, 12, 6, 120, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		5, 8, 128, 8, 8, 10, 8, 12, 8, 131, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 141, 8, 9, 10, 9, 12, 9, 144, 9, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 151, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 165, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 173, 8, 13, 10, 13, 12, 13, 176,
		9, 13, 1, 14, 1, 14, 3, 14, 180, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 192, 8, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 5, 21, 219, 8, 21, 10, 21, 12, 21, 222, 9, 21, 3, 21, 224,
		8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 232, 8, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 5, 24, 239, 8, 24, 10, 24, 12, 24, 242,
		9, 24, 1, 25, 1, 25, 1, 25, 5, 25, 247, 8, 25, 10, 25, 12, 25, 250, 9,
		25, 1, 26, 1, 26, 1, 26, 3, 26, 255, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 265, 8, 27, 1, 27, 0, 0, 28, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 0, 5, 1, 0, 30, 31, 1, 0, 11, 14, 1, 0, 15, 16,
		1, 0, 17, 18, 2, 0, 15, 16, 19, 19, 268, 0, 56, 1, 0, 0, 0, 2, 73, 1, 0,
		0, 0, 4, 80, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10, 97,
		1, 0, 0, 0, 12, 113, 1, 0, 0, 0, 14, 121, 1, 0, 0, 0, 16, 125, 1, 0, 0,
		0, 18, 150, 1, 0, 0, 0, 20, 152, 1, 0, 0, 0, 22, 156, 1, 0, 0, 0, 24, 161,
		1, 0, 0, 0, 26, 169, 1, 0, 0, 0, 28, 179, 1, 0, 0, 0, 30, 181, 1, 0, 0,
		0, 32, 195, 1, 0, 0, 0, 34, 197, 1, 0, 0, 0, 36, 199, 1, 0, 0, 0, 38, 209,
		1, 0, 0, 0, 40, 211, 1, 0, 0, 0, 42, 213, 1, 0, 0, 0, 44, 227, 1, 0, 0,
		0, 46, 233, 1, 0, 0, 0, 48, 235, 1, 0, 0, 0, 50, 243, 1, 0, 0, 0, 52, 254,
		1, 0, 0, 0, 54, 264, 1, 0, 0, 0, 56, 57, 5, 20, 0, 0, 57, 58, 5, 33, 0,
		0, 58, 60, 5, 1, 0, 0, 59, 61, 3, 2, 1, 0, 60, 59, 1, 0, 0, 0, 60, 61,
		1, 0, 0, 0, 61, 65, 1, 0, 0, 0, 62, 64, 3, 10, 5, 0, 63, 62, 1, 0, 0, 0,
		64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1,
		0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 22, 0, 0, 69, 70, 3, 16, 8, 0,
		70, 71, 5, 23, 0, 0, 71, 72, 5, 0, 0, 1, 72, 1, 1, 0, 0, 0, 73, 77, 5,
		21, 0, 0, 74, 76, 3, 4, 2, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 3, 1, 0, 0, 0, 79, 77, 1, 0, 0,
		0, 80, 81, 3, 6, 3, 0, 81, 82, 5, 2, 0, 0, 82, 83, 3, 8, 4, 0, 83, 84,
		5, 1, 0, 0, 84, 5, 1, 0, 0, 0, 85, 90, 5, 33, 0, 0, 86, 87, 5, 3, 0, 0,
		87, 89, 5, 33, 0, 0, 88, 86, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 7, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93,
		94, 7, 0, 0, 0, 94, 9, 1, 0, 0, 0, 95, 98, 5, 29, 0, 0, 96, 98, 3, 8, 4,
		0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100,
		5, 33, 0, 0, 100, 102, 5, 4, 0, 0, 101, 103, 3, 12, 6, 0, 102, 101, 1,
		0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 5, 5, 0,
		0, 105, 107, 5, 6, 0, 0, 106, 108, 3, 2, 1, 0, 107, 106, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 3, 16, 8, 0, 110, 111,
		5, 7, 0, 0, 111, 112, 5, 1, 0, 0, 112, 11, 1, 0, 0, 0, 113, 118, 3, 14,
		7, 0, 114, 115, 5, 3, 0, 0, 115, 117, 3, 14, 7, 0, 116, 114, 1, 0, 0, 0,
		117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		13, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122, 5, 33, 0, 0, 122, 123,
		5, 2, 0, 0, 123, 124, 3, 8, 4, 0, 124, 15, 1, 0, 0, 0, 125, 129, 5, 6,
		0, 0, 126, 128, 3, 18, 9, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0,
		129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 132, 133, 5, 7, 0, 0, 133, 17, 1, 0, 0, 0, 134, 151, 3,
		22, 11, 0, 135, 151, 3, 24, 12, 0, 136, 151, 3, 30, 15, 0, 137, 151, 3,
		36, 18, 0, 138, 142, 5, 8, 0, 0, 139, 141, 3, 18, 9, 0, 140, 139, 1, 0,
		0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0,
		143, 145, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 151, 5, 9, 0, 0, 146,
		147, 3, 42, 21, 0, 147, 148, 5, 1, 0, 0, 148, 151, 1, 0, 0, 0, 149, 151,
		3, 20, 10, 0, 150, 134, 1, 0, 0, 0, 150, 135, 1, 0, 0, 0, 150, 136, 1,
		0, 0, 0, 150, 137, 1, 0, 0, 0, 150, 138, 1, 0, 0, 0, 150, 146, 1, 0, 0,
		0, 150, 149, 1, 0, 0, 0, 151, 19, 1, 0, 0, 0, 152, 153, 5, 32, 0, 0, 153,
		154, 3, 44, 22, 0, 154, 155, 5, 1, 0, 0, 155, 21, 1, 0, 0, 0, 156, 157,
		5, 33, 0, 0, 157, 158, 5, 10, 0, 0, 158, 159, 3, 44, 22, 0, 159, 160, 5,
		1, 0, 0, 160, 23, 1, 0, 0, 0, 161, 162, 5, 28, 0, 0, 162, 164, 5, 4, 0,
		0, 163, 165, 3, 26, 13, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 167, 5, 5, 0, 0, 167, 168, 5, 1, 0, 0, 168,
		25, 1, 0, 0, 0, 169, 174, 3, 28, 14, 0, 170, 171, 5, 3, 0, 0, 171, 173,
		3, 28, 14, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1,
		0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 27, 1, 0, 0, 0, 176, 174, 1, 0, 0,
		0, 177, 180, 5, 36, 0, 0, 178, 180, 3, 44, 22, 0, 179, 177, 1, 0, 0, 0,
		179, 178, 1, 0, 0, 0, 180, 29, 1, 0, 0, 0, 181, 182, 5, 24, 0, 0, 182,
		183, 5, 4, 0, 0, 183, 184, 3, 44, 22, 0, 184, 185, 5, 5, 0, 0, 185, 186,
		3, 32, 16, 0, 186, 191, 3, 16, 8, 0, 187, 188, 5, 25, 0, 0, 188, 189, 3,
		34, 17, 0, 189, 190, 3, 16, 8, 0, 190, 192, 1, 0, 0, 0, 191, 187, 1, 0,
		0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 5, 1, 0, 0,
		194, 31, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 33, 1, 0, 0, 0, 197, 198,
		1, 0, 0, 0, 198, 35, 1, 0, 0, 0, 199, 200, 5, 26, 0, 0, 200, 201, 3, 38,
		19, 0, 201, 202, 5, 4, 0, 0, 202, 203, 3, 44, 22, 0, 203, 204, 5, 5, 0,
		0, 204, 205, 3, 40, 20, 0, 205, 206, 5, 27, 0, 0, 206, 207, 3, 16, 8, 0,
		207, 208, 5, 1, 0, 0, 208, 37, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 39,
		1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 41, 1, 0, 0, 0, 213, 214, 5, 33,
		0, 0, 214, 223, 5, 4, 0, 0, 215, 220, 3, 44, 22, 0, 216, 217, 5, 3, 0,
		0, 217, 219, 3, 44, 22, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0,
		220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 223, 215, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 226, 5, 5, 0, 0, 226, 43, 1, 0, 0, 0, 227, 231, 3, 48,
		24, 0, 228, 229, 3, 46, 23, 0, 229, 230, 3, 48, 24, 0, 230, 232, 1, 0,
		0, 0, 231, 228, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 45, 1, 0, 0, 0,
		233, 234, 7, 1, 0, 0, 234, 47, 1, 0, 0, 0, 235, 240, 3, 50, 25, 0, 236,
		237, 7, 2, 0, 0, 237, 239, 3, 50, 25, 0, 238, 236, 1, 0, 0, 0, 239, 242,
		1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 49, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 243, 248, 3, 52, 26, 0, 244, 245, 7, 3, 0,
		0, 245, 247, 3, 52, 26, 0, 246, 244, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0,
		248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 51, 1, 0, 0, 0, 250, 248,
		1, 0, 0, 0, 251, 252, 7, 4, 0, 0, 252, 255, 3, 52, 26, 0, 253, 255, 3,
		54, 27, 0, 254, 251, 1, 0, 0, 0, 254, 253, 1, 0, 0, 0, 255, 53, 1, 0, 0,
		0, 256, 257, 5, 4, 0, 0, 257, 258, 3, 44, 22, 0, 258, 259, 5, 5, 0, 0,
		259, 265, 1, 0, 0, 0, 260, 265, 3, 42, 21, 0, 261, 265, 5, 33, 0, 0, 262,
		265, 5, 34, 0, 0, 263, 265, 5, 35, 0, 0, 264, 256, 1, 0, 0, 0, 264, 260,
		1, 0, 0, 0, 264, 261, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0,
		0, 0, 265, 55, 1, 0, 0, 0, 22, 60, 65, 77, 90, 97, 102, 107, 118, 129,
		142, 150, 164, 174, 179, 191, 220, 223, 231, 240, 248, 254, 264,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PatitoParserInit initializes any static state used to implement PatitoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPatitoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PatitoParserInit() {
	staticData := &PatitoParserStaticData
	staticData.once.Do(patitoParserInit)
}

// NewPatitoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPatitoParser(input antlr.TokenStream) *PatitoParser {
	PatitoParserInit()
	this := new(PatitoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PatitoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Patito.g4"

	return this
}

// PatitoParser tokens.
const (
	PatitoParserEOF          = antlr.TokenEOF
	PatitoParserT__0         = 1
	PatitoParserT__1         = 2
	PatitoParserT__2         = 3
	PatitoParserT__3         = 4
	PatitoParserT__4         = 5
	PatitoParserT__5         = 6
	PatitoParserT__6         = 7
	PatitoParserT__7         = 8
	PatitoParserT__8         = 9
	PatitoParserT__9         = 10
	PatitoParserT__10        = 11
	PatitoParserT__11        = 12
	PatitoParserT__12        = 13
	PatitoParserT__13        = 14
	PatitoParserT__14        = 15
	PatitoParserT__15        = 16
	PatitoParserT__16        = 17
	PatitoParserT__17        = 18
	PatitoParserT__18        = 19
	PatitoParserPROGRAMA     = 20
	PatitoParserVARS         = 21
	PatitoParserINICIO       = 22
	PatitoParserFIN          = 23
	PatitoParserSI           = 24
	PatitoParserSINO         = 25
	PatitoParserMIENTRAS     = 26
	PatitoParserHAZ          = 27
	PatitoParserESCRIBE      = 28
	PatitoParserNULA         = 29
	PatitoParserENTERO       = 30
	PatitoParserFLOTANTE     = 31
	PatitoParserRETURN       = 32
	PatitoParserID           = 33
	PatitoParserINT          = 34
	PatitoParserFLOAT        = 35
	PatitoParserSTRING       = 36
	PatitoParserWS           = 37
	PatitoParserLINE_COMMENT = 38
	PatitoParserBODY_COMMENT = 39
)

// PatitoParser rules.
const (
	PatitoParserRULE_program       = 0
	PatitoParserRULE_vars          = 1
	PatitoParserRULE_varDecl       = 2
	PatitoParserRULE_idList        = 3
	PatitoParserRULE_type_         = 4
	PatitoParserRULE_funcDef       = 5
	PatitoParserRULE_params        = 6
	PatitoParserRULE_param         = 7
	PatitoParserRULE_body          = 8
	PatitoParserRULE_stmt          = 9
	PatitoParserRULE_returnStmt    = 10
	PatitoParserRULE_assign        = 11
	PatitoParserRULE_print         = 12
	PatitoParserRULE_printItemList = 13
	PatitoParserRULE_printItem     = 14
	PatitoParserRULE_ifStmt        = 15
	PatitoParserRULE_ifMarkThen    = 16
	PatitoParserRULE_elseMark      = 17
	PatitoParserRULE_whileStmt     = 18
	PatitoParserRULE_whileStart    = 19
	PatitoParserRULE_whileCond     = 20
	PatitoParserRULE_call          = 21
	PatitoParserRULE_expr          = 22
	PatitoParserRULE_relop         = 23
	PatitoParserRULE_sum           = 24
	PatitoParserRULE_prod          = 25
	PatitoParserRULE_unary         = 26
	PatitoParserRULE_primary       = 27
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROGRAMA() antlr.TerminalNode
	ID() antlr.TerminalNode
	INICIO() antlr.TerminalNode
	Body() IBodyContext
	FIN() antlr.TerminalNode
	EOF() antlr.TerminalNode
	Vars() IVarsContext
	AllFuncDef() []IFuncDefContext
	FuncDef(i int) IFuncDefContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PROGRAMA() antlr.TerminalNode {
	return s.GetToken(PatitoParserPROGRAMA, 0)
}

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *ProgramContext) INICIO() antlr.TerminalNode {
	return s.GetToken(PatitoParserINICIO, 0)
}

func (s *ProgramContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgramContext) FIN() antlr.TerminalNode {
	return s.GetToken(PatitoParserFIN, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PatitoParserEOF, 0)
}

func (s *ProgramContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *ProgramContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PatitoParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(PatitoParserPROGRAMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PatitoParserVARS {
		{
			p.SetState(59)
			p.Vars()
		}

	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0 {
		{
			p.SetState(62)
			p.FuncDef()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(PatitoParserINICIO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Body()
	}
	{
		p.SetState(70)
		p.Match(PatitoParserFIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(PatitoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARS() antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_vars
	return p
}

func InitEmptyVarsContext(p *VarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_vars
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) VARS() antlr.TerminalNode {
	return s.GetToken(PatitoParserVARS, 0)
}

func (s *VarsContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarsContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitVars(s)
	}
}

func (s *VarsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitVars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PatitoParserRULE_vars)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(PatitoParserVARS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserID {
		{
			p.SetState(74)
			p.VarDecl()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdList() IIdListContext
	Type_() IType_Context

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *VarDeclContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PatitoParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.IdList()
	}
	{
		p.SetState(81)
		p.Match(PatitoParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Type_()
	}
	{
		p.SetState(83)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_idList
	return p
}

func InitEmptyIdListContext(p *IdListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_idList
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserID)
}

func (s *IdListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserID, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterIdList(s)
	}
}

func (s *IdListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitIdList(s)
	}
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PatitoParserRULE_idList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserT__2 {
		{
			p.SetState(86)
			p.Match(PatitoParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(PatitoParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTERO() antlr.TerminalNode
	FLOTANTE() antlr.TerminalNode

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) ENTERO() antlr.TerminalNode {
	return s.GetToken(PatitoParserENTERO, 0)
}

func (s *Type_Context) FLOTANTE() antlr.TerminalNode {
	return s.GetToken(PatitoParserFLOTANTE, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PatitoParserRULE_type_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PatitoParserENTERO || _la == PatitoParserFLOTANTE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Body() IBodyContext
	NULA() antlr.TerminalNode
	Type_() IType_Context
	Params() IParamsContext
	Vars() IVarsContext

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *FuncDefContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FuncDefContext) NULA() antlr.TerminalNode {
	return s.GetToken(PatitoParserNULA, 0)
}

func (s *FuncDefContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FuncDefContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncDefContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (s *FuncDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitFuncDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PatitoParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PatitoParserNULA:
		{
			p.SetState(95)
			p.Match(PatitoParserNULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PatitoParserENTERO, PatitoParserFLOTANTE:
		{
			p.SetState(96)
			p.Type_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(99)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(PatitoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PatitoParserID {
		{
			p.SetState(101)
			p.Params()
		}

	}
	{
		p.SetState(104)
		p.Match(PatitoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(PatitoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PatitoParserVARS {
		{
			p.SetState(106)
			p.Vars()
		}

	}
	{
		p.SetState(109)
		p.Body()
	}
	{
		p.SetState(110)
		p.Match(PatitoParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PatitoParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Param()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserT__2 {
		{
			p.SetState(114)
			p.Match(PatitoParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Param()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_() IType_Context

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *ParamContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PatitoParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(PatitoParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PatitoParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(PatitoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13237223680) != 0 {
		{
			p.SetState(126)
			p.Stmt()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(PatitoParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext
	Print_() IPrintContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Call() ICallContext
	ReturnStmt() IReturnStmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StmtContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *StmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PatitoParserRULE_stmt)
	var _la int

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Print_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.IfStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.WhileStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Match(PatitoParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13237223680) != 0 {
			{
				p.SetState(139)
				p.Stmt()
			}

			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(145)
			p.Match(PatitoParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
			p.Call()
		}
		{
			p.SetState(147)
			p.Match(PatitoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(149)
			p.ReturnStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PatitoParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(PatitoParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Expr()
	}
	{
		p.SetState(154)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PatitoParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(PatitoParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Expr()
	}
	{
		p.SetState(159)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCRIBE() antlr.TerminalNode
	PrintItemList() IPrintItemListContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) ESCRIBE() antlr.TerminalNode {
	return s.GetToken(PatitoParserESCRIBE, 0)
}

func (s *PrintContext) PrintItemList() IPrintItemListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintItemListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintItemListContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PatitoParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(PatitoParserESCRIBE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(PatitoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&128849641488) != 0 {
		{
			p.SetState(163)
			p.PrintItemList()
		}

	}
	{
		p.SetState(166)
		p.Match(PatitoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintItemListContext is an interface to support dynamic dispatch.
type IPrintItemListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrintItem() []IPrintItemContext
	PrintItem(i int) IPrintItemContext

	// IsPrintItemListContext differentiates from other interfaces.
	IsPrintItemListContext()
}

type PrintItemListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintItemListContext() *PrintItemListContext {
	var p = new(PrintItemListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_printItemList
	return p
}

func InitEmptyPrintItemListContext(p *PrintItemListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_printItemList
}

func (*PrintItemListContext) IsPrintItemListContext() {}

func NewPrintItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintItemListContext {
	var p = new(PrintItemListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_printItemList

	return p
}

func (s *PrintItemListContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintItemListContext) AllPrintItem() []IPrintItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintItemContext); ok {
			len++
		}
	}

	tst := make([]IPrintItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintItemContext); ok {
			tst[i] = t.(IPrintItemContext)
			i++
		}
	}

	return tst
}

func (s *PrintItemListContext) PrintItem(i int) IPrintItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintItemContext)
}

func (s *PrintItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrintItemList(s)
	}
}

func (s *PrintItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrintItemList(s)
	}
}

func (s *PrintItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrintItemList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) PrintItemList() (localctx IPrintItemListContext) {
	localctx = NewPrintItemListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PatitoParserRULE_printItemList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.PrintItem()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserT__2 {
		{
			p.SetState(170)
			p.Match(PatitoParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.PrintItem()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintItemContext is an interface to support dynamic dispatch.
type IPrintItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Expr() IExprContext

	// IsPrintItemContext differentiates from other interfaces.
	IsPrintItemContext()
}

type PrintItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintItemContext() *PrintItemContext {
	var p = new(PrintItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_printItem
	return p
}

func InitEmptyPrintItemContext(p *PrintItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_printItem
}

func (*PrintItemContext) IsPrintItemContext() {}

func NewPrintItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintItemContext {
	var p = new(PrintItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_printItem

	return p
}

func (s *PrintItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintItemContext) STRING() antlr.TerminalNode {
	return s.GetToken(PatitoParserSTRING, 0)
}

func (s *PrintItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrintItem(s)
	}
}

func (s *PrintItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrintItem(s)
	}
}

func (s *PrintItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrintItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) PrintItem() (localctx IPrintItemContext) {
	localctx = NewPrintItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PatitoParserRULE_printItem)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PatitoParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(PatitoParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PatitoParserT__3, PatitoParserT__14, PatitoParserT__15, PatitoParserT__18, PatitoParserID, PatitoParserINT, PatitoParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SI() antlr.TerminalNode
	Expr() IExprContext
	IfMarkThen() IIfMarkThenContext
	AllBody() []IBodyContext
	Body(i int) IBodyContext
	SINO() antlr.TerminalNode
	ElseMark() IElseMarkContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) SI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSI, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) IfMarkThen() IIfMarkThenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfMarkThenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfMarkThenContext)
}

func (s *IfStmtContext) AllBody() []IBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyContext); ok {
			len++
		}
	}

	tst := make([]IBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyContext); ok {
			tst[i] = t.(IBodyContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Body(i int) IBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *IfStmtContext) SINO() antlr.TerminalNode {
	return s.GetToken(PatitoParserSINO, 0)
}

func (s *IfStmtContext) ElseMark() IElseMarkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseMarkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseMarkContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PatitoParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(PatitoParserSI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(PatitoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Expr()
	}
	{
		p.SetState(184)
		p.Match(PatitoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.IfMarkThen()
	}
	{
		p.SetState(186)
		p.Body()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PatitoParserSINO {
		{
			p.SetState(187)
			p.Match(PatitoParserSINO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.ElseMark()
		}
		{
			p.SetState(189)
			p.Body()
		}

	}
	{
		p.SetState(193)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfMarkThenContext is an interface to support dynamic dispatch.
type IIfMarkThenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfMarkThenContext differentiates from other interfaces.
	IsIfMarkThenContext()
}

type IfMarkThenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfMarkThenContext() *IfMarkThenContext {
	var p = new(IfMarkThenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_ifMarkThen
	return p
}

func InitEmptyIfMarkThenContext(p *IfMarkThenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_ifMarkThen
}

func (*IfMarkThenContext) IsIfMarkThenContext() {}

func NewIfMarkThenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfMarkThenContext {
	var p = new(IfMarkThenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_ifMarkThen

	return p
}

func (s *IfMarkThenContext) GetParser() antlr.Parser { return s.parser }
func (s *IfMarkThenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfMarkThenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfMarkThenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterIfMarkThen(s)
	}
}

func (s *IfMarkThenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitIfMarkThen(s)
	}
}

func (s *IfMarkThenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitIfMarkThen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) IfMarkThen() (localctx IIfMarkThenContext) {
	localctx = NewIfMarkThenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PatitoParserRULE_ifMarkThen)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseMarkContext is an interface to support dynamic dispatch.
type IElseMarkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElseMarkContext differentiates from other interfaces.
	IsElseMarkContext()
}

type ElseMarkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseMarkContext() *ElseMarkContext {
	var p = new(ElseMarkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_elseMark
	return p
}

func InitEmptyElseMarkContext(p *ElseMarkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_elseMark
}

func (*ElseMarkContext) IsElseMarkContext() {}

func NewElseMarkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseMarkContext {
	var p = new(ElseMarkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_elseMark

	return p
}

func (s *ElseMarkContext) GetParser() antlr.Parser { return s.parser }
func (s *ElseMarkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseMarkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseMarkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterElseMark(s)
	}
}

func (s *ElseMarkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitElseMark(s)
	}
}

func (s *ElseMarkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitElseMark(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ElseMark() (localctx IElseMarkContext) {
	localctx = NewElseMarkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PatitoParserRULE_elseMark)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MIENTRAS() antlr.TerminalNode
	WhileStart() IWhileStartContext
	Expr() IExprContext
	WhileCond() IWhileCondContext
	HAZ() antlr.TerminalNode
	Body() IBodyContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) MIENTRAS() antlr.TerminalNode {
	return s.GetToken(PatitoParserMIENTRAS, 0)
}

func (s *WhileStmtContext) WhileStart() IWhileStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStartContext)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) WhileCond() IWhileCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileCondContext)
}

func (s *WhileStmtContext) HAZ() antlr.TerminalNode {
	return s.GetToken(PatitoParserHAZ, 0)
}

func (s *WhileStmtContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PatitoParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(PatitoParserMIENTRAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.WhileStart()
	}
	{
		p.SetState(201)
		p.Match(PatitoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Expr()
	}
	{
		p.SetState(203)
		p.Match(PatitoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.WhileCond()
	}
	{
		p.SetState(205)
		p.Match(PatitoParserHAZ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Body()
	}
	{
		p.SetState(207)
		p.Match(PatitoParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStartContext is an interface to support dynamic dispatch.
type IWhileStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhileStartContext differentiates from other interfaces.
	IsWhileStartContext()
}

type WhileStartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStartContext() *WhileStartContext {
	var p = new(WhileStartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileStart
	return p
}

func InitEmptyWhileStartContext(p *WhileStartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileStart
}

func (*WhileStartContext) IsWhileStartContext() {}

func NewWhileStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStartContext {
	var p = new(WhileStartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_whileStart

	return p
}

func (s *WhileStartContext) GetParser() antlr.Parser { return s.parser }
func (s *WhileStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterWhileStart(s)
	}
}

func (s *WhileStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitWhileStart(s)
	}
}

func (s *WhileStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitWhileStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) WhileStart() (localctx IWhileStartContext) {
	localctx = NewWhileStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PatitoParserRULE_whileStart)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileCondContext is an interface to support dynamic dispatch.
type IWhileCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhileCondContext differentiates from other interfaces.
	IsWhileCondContext()
}

type WhileCondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileCondContext() *WhileCondContext {
	var p = new(WhileCondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileCond
	return p
}

func InitEmptyWhileCondContext(p *WhileCondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_whileCond
}

func (*WhileCondContext) IsWhileCondContext() {}

func NewWhileCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileCondContext {
	var p = new(WhileCondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_whileCond

	return p
}

func (s *WhileCondContext) GetParser() antlr.Parser { return s.parser }
func (s *WhileCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterWhileCond(s)
	}
}

func (s *WhileCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitWhileCond(s)
	}
}

func (s *WhileCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitWhileCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) WhileCond() (localctx IWhileCondContext) {
	localctx = NewWhileCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PatitoParserRULE_whileCond)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *CallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PatitoParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(PatitoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(PatitoParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60130164752) != 0 {
		{
			p.SetState(215)
			p.Expr()
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PatitoParserT__2 {
			{
				p.SetState(216)
				p.Match(PatitoParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(217)
				p.Expr()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(225)
		p.Match(PatitoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RelExprContext struct {
	ExprContext
}

func NewRelExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelExprContext {
	var p = new(RelExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExprContext) AllSum() []ISumContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISumContext); ok {
			len++
		}
	}

	tst := make([]ISumContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISumContext); ok {
			tst[i] = t.(ISumContext)
			i++
		}
	}

	return tst
}

func (s *RelExprContext) Sum(i int) ISumContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *RelExprContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *RelExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterRelExpr(s)
	}
}

func (s *RelExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitRelExpr(s)
	}
}

func (s *RelExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitRelExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PatitoParserRULE_expr)
	var _la int

	localctx = NewRelExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Sum()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0 {
		{
			p.SetState(228)
			p.Relop()
		}
		{
			p.SetState(229)
			p.Sum()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_relop
	return p
}

func InitEmptyRelopContext(p *RelopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_relop
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (s *RelopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitRelop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PatitoParserRULE_relop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_sum
	return p
}

func InitEmptySumContext(p *SumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_sum
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_sum

	return p
}

func (s *SumContext) GetParser() antlr.Parser { return s.parser }

func (s *SumContext) CopyAll(ctx *SumContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AddSubContext struct {
	SumContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptySumContext(&p.SumContext)
	p.parser = parser
	p.CopyAll(ctx.(*SumContext))

	return p
}

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllProd() []IProdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProdContext); ok {
			len++
		}
	}

	tst := make([]IProdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProdContext); ok {
			tst[i] = t.(IProdContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Prod(i int) IProdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProdContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Sum() (localctx ISumContext) {
	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PatitoParserRULE_sum)
	var _la int

	localctx = NewAddSubContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Prod()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserT__14 || _la == PatitoParserT__15 {
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PatitoParserT__14 || _la == PatitoParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)
			p.Prod()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProdContext is an interface to support dynamic dispatch.
type IProdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsProdContext differentiates from other interfaces.
	IsProdContext()
}

type ProdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProdContext() *ProdContext {
	var p = new(ProdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_prod
	return p
}

func InitEmptyProdContext(p *ProdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_prod
}

func (*ProdContext) IsProdContext() {}

func NewProdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProdContext {
	var p = new(ProdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_prod

	return p
}

func (s *ProdContext) GetParser() antlr.Parser { return s.parser }

func (s *ProdContext) CopyAll(ctx *ProdContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ProdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	ProdContext
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyProdContext(&p.ProdContext)
	p.parser = parser
	p.CopyAll(ctx.(*ProdContext))

	return p
}

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllUnary() []IUnaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryContext); ok {
			len++
		}
	}

	tst := make([]IUnaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryContext); ok {
			tst[i] = t.(IUnaryContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Unary(i int) IUnaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Prod() (localctx IProdContext) {
	localctx = NewProdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PatitoParserRULE_prod)
	var _la int

	localctx = NewMulDivContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Unary()
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserT__16 || _la == PatitoParserT__17 {
		{
			p.SetState(244)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PatitoParserT__16 || _la == PatitoParserT__17) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(245)
			p.Unary()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) CopyAll(ctx *UnaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrefixContext struct {
	UnaryContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (s *PrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToPrimaryContext struct {
	UnaryContext
}

func NewToPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToPrimaryContext {
	var p = new(ToPrimaryContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *ToPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToPrimaryContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ToPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterToPrimary(s)
	}
}

func (s *ToPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitToPrimary(s)
	}
}

func (s *ToPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitToPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PatitoParserRULE_unary)
	var _la int

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PatitoParserT__14, PatitoParserT__15, PatitoParserT__18:
		localctx = NewPrefixContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&622592) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(252)
			p.Unary()
		}

	case PatitoParserT__3, PatitoParserID, PatitoParserINT, PatitoParserFLOAT:
		localctx = NewToPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Primary()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PatitoParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatContext struct {
	PrimaryContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PatitoParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentContext struct {
	PrimaryContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	PrimaryContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallPrimContext struct {
	PrimaryContext
}

func NewCallPrimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallPrimContext {
	var p = new(CallPrimContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *CallPrimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallPrimContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *CallPrimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterCallPrim(s)
	}
}

func (s *CallPrimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitCallPrim(s)
	}
}

func (s *CallPrimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitCallPrim(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	PrimaryContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(PatitoParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PatitoParserRULE_primary)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParensContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(PatitoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Expr()
		}
		{
			p.SetState(258)
			p.Match(PatitoParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewCallPrimContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Call()
		}

	case 3:
		localctx = NewIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Match(PatitoParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Match(PatitoParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(263)
			p.Match(PatitoParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
