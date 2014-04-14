package main

import (
	"flag"
	"os"
)

func main() {
	emj := NewEmoji()

	var list = flag.Bool("list", false, "List all known emoji")
	var n = flag.Bool("n", false, "Supress newline")
	var U = flag.Bool("U", false, "Print code point")
	var o = flag.Bool("o", false, "Print octal escapes suitable for echo -e")
	var O = flag.Bool("O", false, "Print octal escapes suitable for $PS1")
	var x = flag.Bool("x", false, "Print hex escapes suitable for echo -e")
	var e = flag.Bool("e", false, "work like echo, but interpolate \\e as escapes for emoji")
	flag.Parse()

	if *list {
		emj.PrettyPrint(os.Stdout)
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
		str, ok = emj.CpByName(args[0])
	} else if *o {
		str, ok = emj.OctStringByName(args[0], true)
	} else if *O {
		str, ok = emj.OctStringByName(args[0], false)
	} else if *x {
		str, ok = emj.HexStringByName(args[0])
	} else if *e {
		str, ok = emj.InterpolateString(args[0])
	} else {
		str, ok = emj.StringByName(args[0])
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
