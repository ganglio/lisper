package lisper

import (
	"encoding/json"
)

type Lisper struct{}

func NewLisper() *Lisper {
	return &Lisper{}
}

func (l *Lisper) Parse(r string) error {
	out := tokenizer(r)
	exp := parser(out)

	_, _ = json.Marshal(exp)

	return nil
}
