package spec

import (
	"fmt"
	"strings"
)

const stringIndent = "\t"

func specChildrenToString(r Reader, indents ...string) string {
	indents = append(indents, stringIndent)
	if r.ChildCount() > 0 {
		indendation := strings.Join(indents, "")
		entries := []string{}
		for _, child := range r.Children() {
			entries = append(entries, indendation+specAttrsToString(child, indents...))
		}
		return "\n" + strings.Join(entries, "\n")
	}
	return ""
}

func specAttrsToString(r Reader, indents ...string) string {
	result := fmt.Sprintf("%s(Width: %.2f, Height: %.2f", r.SpecName(), r.Width(), r.Height())
	if r.Text() != "" {
		result += ", Text: " + r.Text()
	}

	if r.ChildCount() > 0 {
		kidString := specChildrenToString(r, indents...)
		result += kidString
		return fmt.Sprintf("%v\n%v)", result, strings.Join(indents, ""))
	}

	return fmt.Sprintf("%v)", result)
}

func String(r Reader) string {
	if r == nil {
		return ""
	}
	return specAttrsToString(r)
}
