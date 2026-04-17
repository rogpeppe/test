package fmt

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
)

// Fmt does something...
func Fmt(x string) string {
	return "foo2: " + x
}











func Guillemet(strs []string, args []cue.Value) (string, error) {
	var buf strings.Builder
	for i, s := range strs {
		buf.WriteString(s)
		if i < len(args) {
			fmt.Fprintf(&buf, "«%v»", args[i])
		}
	}
	return buf.String(), nil
}
