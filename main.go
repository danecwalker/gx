package main

import (
	"io"
	"os"
	"os/exec"

	gx "github.com/danecwalker/gx/pkg/gx"
)

func main() {
	f, err := os.Open("./examples/h1/home.gx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	res := gx.GetRender(string(s))
	_f, err := os.Create("./examples/h1/home.go")
	if err != nil {
		panic(err)
	}
	defer _f.Close()
	_f.Write([]byte(res))

	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", "main.wasm", "./examples/h1")
	cmd.Run()
}
