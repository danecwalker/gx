package gx

import (
	"fmt"
	"regexp"
	"strings"
)

func GetRender(doc string) string {
	exp := regexp.MustCompile(`(?P<return>return\s(?P<gx><([\w]*)>[\s\S]*<\/([\w]*)>))`)
	match := exp.FindStringSubmatch(doc)
	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	t := transpile(lex(result["gx"]))
	n := strings.Replace(doc, result["gx"], "NewComponent(instance, create_component)", -1)

	return t + "\n\n" + n
}

type vr struct {
	t      int
	node   NodeP
	name   string
	parent string
}

var vs = make(map[string]int)

func getVars(nodes []NodeP, parent string) []vr {
	a := []vr{}
	for _, node := range nodes {
		nn := fmt.Sprintf("%s%d", node.Tag, vs[node.Tag])
		vs[node.Tag] += 1
		a = append(a, vr{t: 0, name: nn, parent: parent, node: node})
		if node.Body != nil {
			a = append(a, getVars(node.Body, nn)...)
		}
		if node.Content != "" {
			a = append(a, vr{t: 1, name: fmt.Sprintf("t%d", vs["t"]), parent: nn, node: node})
			vs["t"] += 1
		}
	}

	return a
}

func transpile(node *NodeP) string {
	content := node.Body

	variables := getVars(content, "")
	vars := ""
	build := ""
	_d := ""
	for _, v := range variables {
		_t := ""
		if v.t == 0 {
			_t = "*El"
			build += fmt.Sprintf(`%s%s = Element("%s")%s`, strings.Repeat(" ", 6), v.name, v.node.Tag, "\n")
		} else if v.t == 1 {
			_t = "*TextEl"
			build += fmt.Sprintf(`%s%s = Text("%s")%s`, strings.Repeat(" ", 6), v.name, v.node.Content, "\n")
		}

		vars += fmt.Sprintf("%svar %s %s\n", strings.Repeat(" ", 2), v.name, _t)
		_d += fmt.Sprintf(`%sDetach(%s)%s`, strings.Repeat(" ", 6), v.name, "\n")
	}
	vars += "  var mounted bool\n"
	_d += fmt.Sprintf("%smounted = false\n", strings.Repeat(" ", 6))

	create := fmt.Sprintf(
		"%sC: func() {\n%s%s}",
		strings.Repeat(" ", 4),
		build,
		strings.Repeat(" ", 4),
	)

	_mi := ""
	_ma := ""

	for _, v := range variables {
		if v.parent == "" {
			_mi += fmt.Sprintf(`%sInsert(target, %s, anchor)%s`, strings.Repeat(" ", 6), v.name, "\n")
		} else {
			_ma += fmt.Sprintf(`%sAppend(%s, %s)%s`, strings.Repeat(" ", 6), v.parent, v.name, "\n")
		}
	}

	_m := _mi + _ma
	_m += fmt.Sprintf("%sif !mounted { mounted = true }\n", strings.Repeat(" ", 6))

	_l := ""

	mount := fmt.Sprintf(
		"%sM: func(target js.Value, anchor js.Value) {\n%s%s%s}",
		strings.Repeat(" ", 4),
		_m,
		_l,
		strings.Repeat(" ", 4),
	)

	destroy := fmt.Sprintf(
		"%sD: func() {\n%s%s}",
		strings.Repeat(" ", 4),
		_d,
		strings.Repeat(" ", 4),
	)

	component := "package main\n\nimport (\n  \"syscall/js\"\n  . \"github.com/danecwalker/gx/pkg/gxruntime\"\n)\n\n"
	component += "var (\n  v = []uint64{ 0b0000000000000000000000000000000000000000000000000000000000000000 }\n)\n"
	component += "func instance(self *ComponentInstance, _props *map[string]any, _invalidate InvalidateFn) []any {\n  return []any{}\n}\n"
	component += fmt.Sprintf("func create_component(ctx []any) Component {\n%s  return Component{\n%s,\n%s,\n%s,\n%s}\n}\n", vars, create, mount, destroy, strings.Repeat(" ", 2))
	// component += "func Home() *ComponentInstance {\n  return NewComponent(instance, create_component)\n}"
	return component
}
