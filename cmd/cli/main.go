package main

import (
	"fmt"
	"log"

	"github.com/faizanabbas/godo/internal/list"
)

func main() {
	list, err := list.New("godo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()
	godo1, err := list.Add("Water the plants")
	if err != nil {
		log.Fatal(err)
	}
	_, err = list.Add("Buy milk")
	if err != nil {
		log.Fatal(err)
	}
	err = list.Complete(godo1.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
}
