package gx

import (
	"bufio"
	"io"
	"unicode"
)

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	input    *bufio.Reader
	position Position
}

type ILexer interface {
	Next() Token
	Peek() rune
	Consume() rune
}

func NewLexer(input *bufio.Reader) ILexer {
	return &Lexer{input: input}
}

func (l *Lexer) NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.position.Line,
		Column:  l.position.Column,
	}
}

func (l *Lexer) Next() Token {
	l.skipWhiteSpace()

	ch := l.Peek()

	if ch == 0 {
		l.Consume()
		return l.NewToken(EOF, "\x00")
	}

	switch ch {
	case '.':
		l.Consume()
		return l.NewToken(MODIFIER, string(ch))
	case '=':
		lit := string(l.Consume())
		if l.Peek() == '=' {
			lit += string(l.Consume())
			return l.NewToken(EQ, lit)
		}
		return l.NewToken(ASSIGN, lit)
	case '*':
		l.Consume()
		return l.NewToken(ASTERISK, string(ch))
	case '/':
		l.Consume()
		return l.NewToken(SLASH, string(ch))
	case '-':
		l.Consume()
		return l.NewToken(MINUS, string(ch))
	case '+':
		l.Consume()
		return l.NewToken(PLUS, string(ch))
	case ';':
		l.Consume()
		return l.NewToken(SEMI, string(ch))
	case ':':
		l.Consume()
		return l.NewToken(COLON, string(ch))
	case ',':
		l.Consume()
		return l.NewToken(COMMA, string(ch))
	case '(':
		l.Consume()
		return l.NewToken(LPAREN, string(ch))
	case ')':
		l.Consume()
		return l.NewToken(RPAREN, string(ch))
	case '{':
		l.Consume()
		return l.NewToken(LBRACE, string(ch))
	case '}':
		l.Consume()
		return l.NewToken(RBRACE, string(ch))
	case '<':
		lit := string(l.Consume())
		if l.Peek() == '=' {
			lit += string(l.Consume())
			return l.NewToken(LT_EQ, lit)
		}
		return l.NewToken(LT, lit)
	case '>':
		lit := string(l.Consume())
		if l.Peek() == '=' {
			lit += string(l.Consume())
			return l.NewToken(GT_EQ, lit)
		}
		return l.NewToken(GT, lit)
	case '!':
		lit := string(l.Consume())
		if l.Peek() == '=' {
			lit += string(l.Consume())
			return l.NewToken(NOT_EQ, lit)
		}
		return l.NewToken(BANG, lit)
	default:
		if unicode.IsLetter(ch) {
			lit := l.readIdentifier()
			return l.NewToken(l.lookupKeyword(lit), lit)
		} else if unicode.IsDigit(ch) {
			lit := l.readNumber()
			return l.NewToken(NUMBER, lit)
		} else {
			l.Consume()
			return l.NewToken(ILLEGAL, string(ch))
		}
	}
}

func (l *Lexer) Peek() rune {
	ch, err := l.input.Peek(1)
	if err == io.EOF {
		return 0
	}
	return rune(ch[0])
}

func (l *Lexer) Consume() rune {
	l.position.Column++
	ch, _, err := l.input.ReadRune()
	if err == io.EOF {
		return 0
	}
	return ch
}

func (l *Lexer) skipWhiteSpace() {
	for l.Peek() == ' ' || l.Peek() == '\n' || l.Peek() == '\t' || l.Peek() == '\r' {
		if l.Peek() == '\n' {
			l.position.Line++
			l.position.Column = 0
		}
		l.Consume()
	}
}

var keywords = map[string]TokenType{
	"fn":     FUNC,
	"let":    LET,
	"const":  CONST,
	"export": EXPORT,
	"import": IMPORT,
	"string": STRING_TYPE,
	"number": NUMBER_TYPE,
	"bool":   BOOL_TYPE,
	"true":   BOOLEAN,
	"false":  BOOLEAN,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func (l *Lexer) lookupKeyword(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}

func (l *Lexer) readIdentifier() string {
	lit := string(l.Consume())
	for {
		ch := l.Peek()
		if ch == 0 {
			l.Consume()
			return lit
		}

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			lit += string(l.Consume())
		} else {
			break
		}
	}

	return lit
}

func (l *Lexer) readNumber() string {
	lit := string(l.Consume())
	hasDot := false
	for {
		ch := l.Peek()
		if ch == 0 {
			l.Consume()
			return lit
		}

		if unicode.IsDigit(ch) || ch == '.' {
			if ch == '.' {
				if hasDot {
					break
				}
				hasDot = true
			}
			lit += string(l.Consume())
		} else {
			break
		}
	}

	return lit
}
