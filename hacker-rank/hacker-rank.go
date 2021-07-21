package hacker_rank

import (
	"bufio"
	"fmt"
	"os"
)

type HackerRank struct {}

func (h *HackerRank) Resolve() {
	//socks := []int {1,2,1,2,1,3,2}
	//pairs := getPairCount(socks)
	//fmt.Printf("there are %v pair(s) of socks in %v",pairs,socks)

	moduloString()
}

func getPairCount(socks []int) (out int) {
	m := map[int]int{}

	for _, val := range socks { // Count pairs
		if _, ok := m[val]; !ok {
			m[val] = 0
		}
		m[val]++

		if m[val] % 2 == 0 {
			out++
		}
	}
	return
}

func moduloString() {
	//beg, end := "",""
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	for i, r := range text {
		fmt.Printf("I: %v, v: %v\n",i,r)
	}
}