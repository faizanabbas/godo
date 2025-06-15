package godo

import (
	"errors"
	"fmt"
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

func (l *List) Complete(index int) error {
	if index < 0 || index >= len(l.Godos) {
		return errors.New("godo does not exist")
	}
	l.Godos[index].Done = true
	return nil
}

func (l *List) String() string {
	if len(l.Godos) == 0 {
		return "No godos in list"
	}
	result := "Godo list:\n"
	for i, godo := range l.Godos {
		status := " "
		if godo.Done {
			status = "✓"
		}
		result += fmt.Sprintf("%d. [%s] %s\n", i+1, status, godo.Text)
	}
	return result
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
