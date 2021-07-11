package main

import (
	string_split "algorithms-practice/string-split"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	//var iface resolver = &remove_islands.IslandsRemover{}
	//var iface resolver = &hacker_rank.HackerRank{}
	//var iface resolver = &codility.Codility{}
	//var iface resolver = &hacker_rank_thirty.HackerRankThirty{}
	var iface resolver = &string_split.StringSplit{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

