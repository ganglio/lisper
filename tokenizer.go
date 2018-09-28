package lisper

import (
	"regexp"
	"strings"
)

// Tokens is a list of tokens and the counter for the current one
type Tokens struct {
	t []string
	c int
}

// Next returns the next token
func (t *Tokens) Next() (string, bool) {
	if t.c >= len(t.t) {
		return "", true
	}
	toc := t.t[t.c]
	t.c++
	return toc, false
}

// Tokenize converts a string to a sequence of tokens. TODO: Improve the code as it's not massively elegant
func Tokenize(s string) Tokens {

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		s = strings.TrimPrefix(s, "(")
		s = strings.TrimSuffix(s, ")")
	}

	s = strings.Replace(s, "\n", " ", -1)  // Remove \n
	s = strings.Replace(s, "(", " ( ", -1) // Detach (
	s = strings.Replace(s, ")", " ) ", -1) // Detach )
	s = strings.TrimSpace(s)
	return Tokens{regexp.MustCompile("\\s+").Split(s, -1), 0}
}
