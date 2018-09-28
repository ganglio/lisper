package lisper

import (
	"regexp"
	"strings"
)

type Tokens struct {
	t []string
	c int
}

func (t *Tokens) Next() (string, bool) {
	if t.c >= len(t.t) {
		return "", true
	}
	toc := t.t[t.c]
	t.c++
	return toc, false
}

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
