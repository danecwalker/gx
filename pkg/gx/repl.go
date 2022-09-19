package gx

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type REPL struct {
	prompt string
}

func NewREPL(prompt string) *REPL {
	return &REPL{prompt: prompt}
}

func (r *REPL) Start() {
	env := NewEnvironment()
	for {
		fmt.Print(r.prompt)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		if input == "exit\n" {
			break
		}

		lex := NewLexer(bufio.NewReader(strings.NewReader(input)))
		parser := NewParser(lex)
		program := parser.ParseProgram()
		evaluated := Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
