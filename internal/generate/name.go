package generate

import (
	"slices"
	"strings"
	"unicode"
)

func exportedGoName(name string) string {
	goName := cleanName(name)
	if len(goName) > 0 {
		upperFirstChar := unicode.ToUpper(rune(goName[0]))
		goName = string(upperFirstChar) + goName[1:]
	}

	return goName
}

func goName(name string) string {
	goName := cleanName(name)

	if slices.Contains(goKeywords, goName) {
		goName += "_"
	}

	return goName
}

func cleanName(name string) string {
	goName := strings.TrimPrefix(name, "CX")
	goName = strings.TrimPrefix(goName, "_")
	if len(goName) > 0 {
		lowerFirstChar := unicode.ToLower(rune(goName[0]))
		goName = string(lowerFirstChar) + goName[1:]
	}

	return goName
}

var goKeywords = []string{
	"args", // not a Go keyword, but is used as a var name in function bodies
	"func",
	"interface",
	"len",
	"map",
	"range",
	"select",
	"string",
	"type",
	"var",
}
