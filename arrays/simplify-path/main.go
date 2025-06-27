package main

import (
	"strings"

	stack "github.com/idsulik/go-collections/stack"
)

func SimplifyPath(str string) string {
	st := stack.New[string](len(str))

	sp := strings.Split(str, "/")
	for _, s := range sp {
		switch s {
		case "":
			// Do nothing, consecutive "/"
		case ".":
			// Do nothing, "." doesn't change the path
		case "..":
			if !st.IsEmpty() {
				st.Pop()
			}
		default:
			// Default case is non empty string that is not "." or ".."
			st.Push(s)
		}
	}

	out := ""
	for !st.IsEmpty() {
		e, _ := st.Pop()
		out = "/" + e + out
	}
	if out == "" {
		return "/"
	}

	return out
}
