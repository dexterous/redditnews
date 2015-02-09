package main

import (
	"fmt"
	"github.com/dexterous/redditnews"
	"log"
)

func main() {
	items, err := redditnews.Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
