package main

import (
	"fmt"

	"github.com/faizanabbas/godo/internal/godo"
)

func main() {
	list := godo.NewList()
	list.Add("Water the plants")
	list.Add("Buy milk")
	list.Complete(0)
	fmt.Println(list)
}
