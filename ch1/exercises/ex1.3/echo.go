package echo

import (
	"strings"
)

func EchoIter(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func EchoRange(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func EchoJoin(args []string) string {
	return strings.Join(args, " ")
}
