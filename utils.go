package lisper

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func tokenizer(i string) []string {
	p1 := strings.Replace(i, "\n", " ", -1)   // Remove \n
	p2 := strings.Replace(p1, "(", " ( ", -1) // Detach (
	p3 := strings.Replace(p2, ")", " ) ", -1) // Detach )
	p4 := strings.TrimSpace(p3)
	return regexp.MustCompile("\\s+").Split(p4, -1)
}

func parser(t []string) Atom {
	fmt.Printf("%#v\n", t)
	if len(t) == 0 {
		panic(errors.New("unexpected EOF"))
	}

	token, t := t[0], t[1:]

	if token == "(" {
		var l List
		for token != ")" {
			l = L(parser(t))
			t = t[1:]
		}
		return A(l)
	} else if token == ")" {
		panic("Unexpected )")
	}

	return A(V(token))
}

func floatEquals(a, b float64) bool {
	var eps = 0.00000001
	if (a-b) < eps && (b-a) < eps {
		return true
	}
	return false
}
