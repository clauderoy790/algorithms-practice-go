package main

import (
	hacker_rank "algorithms-practice/hacker-rank"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	//var iface resolver = &remove_islands.IslandsRemover{}
	var iface resolver = &hacker_rank.HackerRank{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

