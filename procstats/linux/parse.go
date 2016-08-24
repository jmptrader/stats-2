package linux

import (
	"strings"
	"unicode"
)

func forEachLine(text string, call func(string)) {
	for len(text) != 0 {
		var line string

		if i := strings.IndexByte(text, '\n'); i >= 0 {
			line, text = text[:i], text[i+1:]
		} else {
			line, text = text, ""
		}

		if line = strings.TrimSpace(line); len(line) != 0 {
			call(line)
		}
	}
}

func forEachLineExceptFirst(text string, call func(string)) {
	first := true
	forEachLine(text, func(line string) {
		if first {
			first = false
		} else {
			call(line)
		}
	})
}

func forEachColumn(line string, call func(string)) {
	for line = skipSpaces(line); len(line) != 0; line = skipSpaces(line) {
		var column string

		if i := strings.Index(line, "  "); i >= 0 {
			column, line = line[:i], line[i+2:]
		} else {
			column, line = line, ""
		}

		call(column)
	}
}

func forEachProperty(text string, call func(string, string)) {
	forEachLine(text, func(line string) { call(splitProperty(line)) })
}

func splitProperty(text string) (key string, val string) {
	if i := strings.IndexByte(text, ':'); i >= 0 {
		key, val = text[:i], text[i+1:]
	} else {
		val = text
	}
	key = strings.TrimSpace(key)
	val = strings.TrimSpace(val)
	return
}

func skipSpaces(text string) string {
	for i, c := range text {
		if !unicode.IsSpace(c) {
			return text[i:]
		}
	}
	return ""
}

func skipLine(text string) string {
	if i := strings.IndexByte(text, '\n'); i >= 0 {
		return text[i+1:]
	}
	return ""
}