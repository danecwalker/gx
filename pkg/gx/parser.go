package gx

import (
	"io"
	"strings"
)

type TokenStruct struct {
	T Token
	V string
}
type Token int

const (
	EOF Token = iota
	ILLEGAL
	OPEN
	CLOSE
	TEXT
)

var Tokens = map[Token]string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	OPEN:    "OPEN",
	CLOSE:   "CLOSE",
	TEXT:    "TEXT",
}

func lex(s string) *NodeP {
	r := strings.NewReader(s)

	tokens := []TokenStruct{}

	for {
		t, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}

		switch t {
		case ' ':
		case '\t':
		case '\n':
			continue
		case '<':
			lexme := ""
			for {
				t, _, err := r.ReadRune()
				if err == io.EOF {
					break
				}
				if t == '>' {
					break
				}
				lexme += string(t)
			}
			if strings.HasPrefix(lexme, "/") {
				tokens = append(tokens, TokenStruct{T: CLOSE, V: lexme[1:]})
			} else {
				tokens = append(tokens, TokenStruct{T: OPEN, V: lexme})
			}
		default:
			if t != '<' && t != '>' && t != '/' {
				lexme := string(t)
				for {
					t, _, err := r.ReadRune()
					if err == io.EOF {
						break
					}
					if t != '<' && t != '>' && t != '/' {
						lexme += string(t)
					} else {
						tokens = append(tokens, TokenStruct{T: TEXT, V: lexme})
						r.UnreadRune()
						break
					}
				}
			}
		}
	}

	return parse(tokens)
}

type NodeP struct {
	Tag     string
	Attr    map[string]string
	Body    []NodeP
	Content string
}

func parse(ts []TokenStruct) *NodeP {
	var root *NodeP
	var stack []*NodeP

	for i := 0; i < len(ts); i++ {
		currentToken := ts[i]
		switch currentToken.T {
		case OPEN:
			node := NodeP{Tag: currentToken.V, Attr: make(map[string]string)}
			stack = append(stack, &node)
		case CLOSE:
			if len(stack) > 1 {
				parent := stack[len(stack)-2]
				parent.Body = append(parent.Body, *stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			} else {
				root = stack[0]
				stack = []*NodeP{}
			}
		case TEXT:
			stack[len(stack)-1].Content = currentToken.V
		}
	}

	return root
}
