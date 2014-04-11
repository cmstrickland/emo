package main

import (
	"fmt"
	"io"
	"sort"
)

type codemap map[string]int

type Emoji struct {
	codes codemap
}

func NewEmoji() *Emoji {
	e := &Emoji{codepoints}
	return e
}

func (e *Emoji) String() string {
	return fmt.Sprintf("%v", e.codes)
}

func (e *Emoji) StringByName(name string) (string, bool) {
	cp, ok := e.codes[name]
	return fmt.Sprintf("%c", cp), ok
}

func (e *Emoji) CpByName(name string) (string, bool) {
	cp, ok := e.codes[name]
	return fmt.Sprintf("%U", cp), ok
}

func (e *Emoji) OctStringByName(name string, pad_prefix bool) (string, bool) {
	cp, ok := e.StringByName(name)
	bytes := []byte(cp)
	template := "\\%o"
	if pad_prefix {
		template = "\\0%o"
	}
	var str string
	for _, b := range bytes {
		str = str + fmt.Sprintf(template, b)
	}
	return str, ok
}

func (e *Emoji) PrettyPrint(w io.Writer) {
	var keys []string
	for k := range e.codes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		io.WriteString(w, fmt.Sprintf("%c - %s\n", e.codes[k], k))
	}
}
