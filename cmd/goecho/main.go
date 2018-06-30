package main

import (
	"flag"
	"fmt"

	"github.com/umaumax/goecho"
)

var (
	noNewLineFlag bool
	unescapeFlag  bool
)

func init() {
	flag.BoolVar(&noNewLineFlag, "n", false, "do not output the trailing newline")
	flag.BoolVar(&unescapeFlag, "E", false, "disable interpretation of backslash escapes")
}

func main() {
	flag.Parse()
	out := goecho.Echo(noNewLineFlag, unescapeFlag, flag.Args())
	fmt.Print(out)
}
