package main

import (
	"flag"
	"os"
)

func main() {
	e := NewEmoji()

	var list = flag.Bool("list", false, "List all known emoji")
	var n = flag.Bool("n", false, "Supress newline")
	var U = flag.Bool("U", false, "Print code point")

	flag.Parse()

	if *list {
		e.PrettyPrint(os.Stdout)
		os.Exit(0)
	}

	var args = flag.Args()

	if len(args) != 1 {
		print("Usage: emo <descriptive-name>. Try emo --list for a list of known emoji. --help for help\n")
		os.Exit(1)
	}

	var str string
	var ok bool
	if *U {
		str, ok = e.CpByName(args[0])
	} else {
		str, ok = e.StringByName(args[0])
	}
	if ok {
		if *n == false {
			str += "\n"
		}
		os.Stdout.Write([]byte(str))
		os.Exit(0)
	} else {
		os.Exit(2)
	}
}
