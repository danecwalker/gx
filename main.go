package main

import (
	"bufio"
	"io"
	"os"

	"github.com/danecwalker/gx/pkg/gx"
)

func main() {
	f, err := os.Open("./examples/h1/h1.gx")
	if err != nil {
		panic(err)
	}

	lex := gx.NewLexer(bufio.NewReader(f))
	p := gx.NewParser(lex)
	program := p.ParseProgram()
	evaluated := gx.Eval(program, gx.NewEnvironment())
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
	// repl := gx.NewREPL(">> ")
	// repl.Start()
}
