package main

import (
	remove_islands "algorithms-practice/remove-islands"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	var iface resolver = &remove_islands.IslandsRemover{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

