package main

import (
	"fmt"
	"io"
	"sort"
)

type codemap map[string]string

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

func (e *Emoji) ByName(name string) (string, bool) {
	str, ok := e.codes[name]
	return str, ok
}

func (e *Emoji) PrettyPrint(w io.Writer) {
	var keys []string
	for k := range e.codes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		io.WriteString(w, fmt.Sprintf("%s  - %s\n", e.codes[k], k))
	}
}
