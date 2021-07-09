package main

import (
	"algorithms-practice/codility"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	//var iface resolver = &remove_islands.IslandsRemover{}
	//var iface resolver = &hacker_rank.HackerRank{}
	var iface resolver = &codility.Codility{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

