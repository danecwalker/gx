package test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/danecwalker/gx/pkg/gx"
)

func TestCharacters(t *testing.T) {
	input := `=(){}-+/*`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.ASSIGN, "="},
		{gx.LPAREN, "("},
		{gx.RPAREN, ")"},
		{gx.LBRACE, "{"},
		{gx.RBRACE, "}"},
		{gx.MINUS, "-"},
		{gx.PLUS, "+"},
		{gx.SLASH, "/"},
		{gx.ASTERISK, "*"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFuncAssignment(t *testing.T) {
	input := `let a = fn(){}`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.LET, "let"},
		{gx.IDENT, "a"},
		{gx.ASSIGN, "="},
		{gx.FUNC, "fn"},
		{gx.LPAREN, "("},
		{gx.RPAREN, ")"},
		{gx.LBRACE, "{"},
		{gx.RBRACE, "}"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestAddProgram(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.LET, "let"},
		{gx.IDENT, "five"},
		{gx.ASSIGN, "="},
		{gx.NUMBER, "5"},
		{gx.SEMI, ";"},
		{gx.LET, "let"},
		{gx.IDENT, "ten"},
		{gx.ASSIGN, "="},
		{gx.NUMBER, "10"},
		{gx.SEMI, ";"},
		{gx.LET, "let"},
		{gx.IDENT, "add"},
		{gx.ASSIGN, "="},
		{gx.FUNC, "fn"},
		{gx.LPAREN, "("},
		{gx.IDENT, "x"},
		{gx.COMMA, ","},
		{gx.IDENT, "y"},
		{gx.RPAREN, ")"},
		{gx.LBRACE, "{"},
		{gx.IDENT, "x"},
		{gx.PLUS, "+"},
		{gx.IDENT, "y"},
		{gx.SEMI, ";"},
		{gx.RBRACE, "}"},
		{gx.SEMI, ";"},
		{gx.LET, "let"},
		{gx.IDENT, "result"},
		{gx.ASSIGN, "="},
		{gx.IDENT, "add"},
		{gx.LPAREN, "("},
		{gx.IDENT, "five"},
		{gx.COMMA, ","},
		{gx.IDENT, "ten"},
		{gx.RPAREN, ")"},
		{gx.SEMI, ";"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestLessThan(t *testing.T) {
	input := `let a = 1 < 2;`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.LET, "let"},
		{gx.IDENT, "a"},
		{gx.ASSIGN, "="},
		{gx.NUMBER, "1"},
		{gx.LT, "<"},
		{gx.NUMBER, "2"},
		{gx.SEMI, ";"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestConstant(t *testing.T) {
	input := `const a = 1;`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.CONST, "const"},
		{gx.IDENT, "a"},
		{gx.ASSIGN, "="},
		{gx.NUMBER, "1"},
		{gx.SEMI, ";"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestModifier(t *testing.T) {
	input := `export const Home = fn(props) {
	props.a
}`

	tests := []struct {
		expectedType    gx.TokenType
		expectedLiteral string
	}{
		{gx.EXPORT, "export"},
		{gx.CONST, "const"},
		{gx.IDENT, "Home"},
		{gx.ASSIGN, "="},
		{gx.FUNC, "fn"},
		{gx.LPAREN, "("},
		{gx.IDENT, "props"},
		{gx.RPAREN, ")"},
		{gx.LBRACE, "{"},
		{gx.IDENT, "props"},
		{gx.MODIFIER, "."},
		{gx.IDENT, "a"},
		{gx.RBRACE, "}"},
		{gx.EOF, "\x00"},
	}

	in := bufio.NewReader(strings.NewReader(input))

	l := gx.NewLexer(in)

	for i, test := range tests {
		tok := l.Next()
		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, gx.TokenMap[test.expectedType], gx.TokenMap[tok.Type])
		}
		if test.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
