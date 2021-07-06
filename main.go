package main

import (
	code_wars "algorithms-practice/code-wars"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	//var iface resolver = &remove_islands.IslandsRemover{}
	var iface resolver = &code_wars.CodeWars{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

