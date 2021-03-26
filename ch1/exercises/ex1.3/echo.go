package echo

import (
	"strings"
)

func EchoAddIter(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
	}
	return s
}

func EchoAddRange(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
	}
	return s
}

func EchoJoin(args []string) string {
	return strings.Join(args, " ")
}
