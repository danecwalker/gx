package test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/danecwalker/gx/pkg/gx"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}
	for _, tt := range tests {
		l := gx.NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		p := gx.NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestLet(t *testing.T) {
	input := `let a = 2;
let b = 4;
let c = 6;`

	l := gx.NewLexer(bufio.NewReader(strings.NewReader(input)))
	p := gx.NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturn(t *testing.T) {
	input := `return a;
return 2;`

	l := gx.NewLexer(bufio.NewReader(strings.NewReader(input)))
	p := gx.NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*gx.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *gx.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s gx.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*gx.LetStatement)
	if !ok {
		t.Errorf("s not *gx.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *gx.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestString(t *testing.T) {
	program := &gx.Program{
		Statements: []gx.Statement{
			&gx.LetStatement{
				Token: gx.Token{Type: gx.LET, Literal: "let"},
				Name: &gx.Identifier{
					Token: gx.Token{Type: gx.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &gx.Identifier{
					Token: gx.Token{Type: gx.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestIdentifiers(t *testing.T) {
	input := "foobar;"

	l := gx.NewLexer(bufio.NewReader(strings.NewReader(input)))
	p := gx.NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*gx.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*gx.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestNumbers(t *testing.T) {
	input := "5;"

	l := gx.NewLexer(bufio.NewReader(strings.NewReader(input)))
	p := gx.NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*gx.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*gx.NumberLiteral)
	if !ok {
		t.Fatalf("exp not *ast.NumberLiteral. got=%T", stmt.Expression)
	}
	if ident.Value != 5 {
		t.Errorf("ident.Value not %s. got=%f", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}
