package main

import (
	hacker_rank_thirty "algorithms-practice/hacker-rank-thirty"
	"fmt"
)

func main() {
	fmt.Println("Algorithms practice!")

	//var iface resolver = &remove_islands.IslandsRemover{}
	//var iface resolver = &hacker_rank.HackerRank{}
	//var iface resolver = &codility.Codility{}
	//var iface resolver = &home_test.HomeTest{}
	var iface resolver = &hacker_rank_thirty.HackerRankThirty{}
	iface.Resolve()
}

type resolver interface {
	Resolve()
}

