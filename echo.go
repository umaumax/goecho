package goecho

import (
	"strings"
)

func Echo(nFlag bool, EFlag bool, args []string) (out string) {
	out = strings.Join(args, " ")
	if !nFlag {
		out += "\n"
	}
	if !EFlag {
		out = EscapeBackslash(out)
	}
	return out
}
