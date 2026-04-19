package generate

import "strings"

func goName(name string) string {
	return strings.TrimPrefix(name, "CX")
}
