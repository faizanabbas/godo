package godo

import (
	"errors"
	"strings"
)

type List struct {
	Godos []Godo
}

func NewList() *List {
	return &List{
		Godos: []Godo{},
	}
}

func (l *List) Add(text string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("text is empty")
	}
	l.Godos = append(l.Godos, NewGodo(text))
	return nil
}

type Godo struct {
	Text string
	Done bool
}

func NewGodo(text string) Godo {
	return Godo{
		Text: text,
		Done: false,
	}
}
