package gx

import "fmt"

type TokenType int
type Token struct {
	Type    TokenType
	Literal string
	Column  int
	Line    int
}

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literals
	IDENT
	NUMBER

	// Operators
	ASSIGN
	PLUS
	MINUS
	SLASH
	ASTERISK
	MODIFIER
	EQ
	NOT_EQ
	GT_EQ
	LT_EQ

	// Delimiters
	LT
	GT
	COLON
	COMMA
	SEMI
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	BANG

	// Keywords
	FUNC
	LET
	CONST
	NUMBER_TYPE
	STRING_TYPE
	BOOL_TYPE
	BOOLEAN
	IMPORT
	EXPORT
	IF
	ELSE
	RETURN
)

var TokenMap = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	// Identifiers + literals
	IDENT:  "IDENT",
	NUMBER: "NUMBER",

	// Operators
	ASSIGN:   "ASSIGN",
	PLUS:     "PLUS",
	MINUS:    "MINUS",
	SLASH:    "SLASH",
	ASTERISK: "ASTERISK",
	MODIFIER: "MODIFIER",
	EQ:       "EQ",
	NOT_EQ:   "NOT_EQ",
	GT_EQ:    "GT_EQ",
	LT_EQ:    "LT_EQ",

	// Delimiters
	LT:     "LT",
	GT:     "GT",
	COMMA:  "COMMA",
	SEMI:   "SEMI",
	COLON:  "COLON",
	LPAREN: "LPAREN",
	RPAREN: "RPAREN",
	LBRACE: "LBRACE",
	RBRACE: "RBRACE",
	BANG:   "BANG",

	// Keywords
	FUNC:        "FUNC",
	LET:         "LET",
	CONST:       "CONST",
	NUMBER_TYPE: "NUMBER_TYPE",
	STRING_TYPE: "STRING_TYPE",
	BOOL_TYPE:   "BOOL_TYPE",
	BOOLEAN:     "BOOLEAN",
	IMPORT:      "IMPORT",
	EXPORT:      "EXPORT",
	IF:          "IF",
	ELSE:        "ELSE",
	RETURN:      "RETURN",
}

func (tt Token) String() string {
	return fmt.Sprintf("%d:%-6d%-20s%s", tt.Line, tt.Column, TokenMap[tt.Type], tt.Literal)
}
