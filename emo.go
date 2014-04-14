package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
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

func (e *Emoji) bytesByName(name string) ([]byte, bool) {
	cp, ok := e.StringByName(name)
	return []byte(cp), ok
}

func (e *Emoji) prettyBytes(template string, bytes []byte) (str string) {
	for _, b := range bytes {
		str = str + fmt.Sprintf(template, b)
	}
	return str
}

func (e *Emoji) OctStringByName(name string, pad_prefix bool) (string, bool) {
	bytes, ok := e.bytesByName(name)
	template := "\\%o"
	if pad_prefix {
		template = "\\0%o"
	}
	return e.prettyBytes(template, bytes), ok
}

func (e *Emoji) HexStringByName(name string) (string, bool) {
	bytes, ok := e.bytesByName(name)
	template := "\\x%x"
	return e.prettyBytes(template, bytes), ok
}

func (e *Emoji) InterpolateString(s string) (string, bool) {
	rx := regexp.MustCompile(`\\e[-a-z]+`)
	matches := rx.FindAllString(s, -1)
	for _, m := range matches {
		name := m[2:]
		em, ok := e.StringByName(name)
		if ok {
			s = strings.Replace(s, m, em, 1)
		}
	}
	return s, true
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
