// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 86, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 5, 2, 35, 10,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 54, 10, 11, 13, 11, 14, 11,
	55, 3, 11, 3, 11, 6, 11, 60, 10, 11, 13, 11, 14, 11, 61, 5, 11, 64, 10,
	11, 3, 12, 3, 12, 7, 12, 68, 10, 12, 12, 12, 14, 12, 71, 11, 12, 3, 13,
	5, 13, 74, 10, 13, 3, 14, 3, 14, 5, 14, 78, 10, 14, 3, 15, 6, 15, 81, 10,
	15, 13, 15, 14, 15, 82, 3, 15, 3, 15, 2, 2, 16, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 2, 23, 12, 25, 2, 27, 2, 29, 13,
	3, 2, 4, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	89, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 3, 34, 3,
	2, 2, 2, 5, 36, 3, 2, 2, 2, 7, 38, 3, 2, 2, 2, 9, 40, 3, 2, 2, 2, 11, 42,
	3, 2, 2, 2, 13, 44, 3, 2, 2, 2, 15, 46, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2,
	19, 50, 3, 2, 2, 2, 21, 53, 3, 2, 2, 2, 23, 65, 3, 2, 2, 2, 25, 73, 3,
	2, 2, 2, 27, 77, 3, 2, 2, 2, 29, 80, 3, 2, 2, 2, 31, 35, 7, 96, 2, 2, 32,
	33, 7, 44, 2, 2, 33, 35, 7, 44, 2, 2, 34, 31, 3, 2, 2, 2, 34, 32, 3, 2,
	2, 2, 35, 4, 3, 2, 2, 2, 36, 37, 7, 44, 2, 2, 37, 6, 3, 2, 2, 2, 38, 39,
	7, 49, 2, 2, 39, 8, 3, 2, 2, 2, 40, 41, 7, 45, 2, 2, 41, 10, 3, 2, 2, 2,
	42, 43, 7, 47, 2, 2, 43, 12, 3, 2, 2, 2, 44, 45, 7, 48, 2, 2, 45, 14, 3,
	2, 2, 2, 46, 47, 7, 42, 2, 2, 47, 16, 3, 2, 2, 2, 48, 49, 7, 43, 2, 2,
	49, 18, 3, 2, 2, 2, 50, 51, 5, 21, 11, 2, 51, 20, 3, 2, 2, 2, 52, 54, 4,
	50, 59, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	55, 56, 3, 2, 2, 2, 56, 63, 3, 2, 2, 2, 57, 59, 5, 13, 7, 2, 58, 60, 4,
	50, 59, 2, 59, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 57, 3, 2, 2, 2, 63, 64, 3,
	2, 2, 2, 64, 22, 3, 2, 2, 2, 65, 69, 5, 25, 13, 2, 66, 68, 5, 27, 14, 2,
	67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3,
	2, 2, 2, 70, 24, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 74, 9, 2, 2, 2, 73,
	72, 3, 2, 2, 2, 74, 26, 3, 2, 2, 2, 75, 78, 5, 25, 13, 2, 76, 78, 4, 50,
	59, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 28, 3, 2, 2, 2, 79,
	81, 9, 3, 2, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 8, 15, 2, 2, 85, 30,
	3, 2, 2, 2, 11, 2, 34, 55, 61, 63, 69, 73, 77, 82, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "SCIENTIFIC_NUMBER",
	"VARIABLE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "SCIENTIFIC_NUMBER",
	"NUMBER", "VARIABLE", "VALID_ID_START", "VALID_ID_CHAR", "WHITESPACE",
}

type AbacusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAbacusLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AbacusLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAbacusLexer(input antlr.CharStream) *AbacusLexer {
	l := new(AbacusLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Abacus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbacusLexer tokens.
const (
	AbacusLexerPOW               = 1
	AbacusLexerMUL               = 2
	AbacusLexerDIV               = 3
	AbacusLexerADD               = 4
	AbacusLexerSUB               = 5
	AbacusLexerPOINT             = 6
	AbacusLexerLPAREN            = 7
	AbacusLexerRPAREN            = 8
	AbacusLexerSCIENTIFIC_NUMBER = 9
	AbacusLexerVARIABLE          = 10
	AbacusLexerWHITESPACE        = 11
)
