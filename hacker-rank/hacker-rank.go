package hacker_rank

import "fmt"

type HackerRank struct {}

func (h *HackerRank) Resolve() {
	socks := []int {1,2,1,2,1,3,2}
	pairs := getPairCount(socks)
	fmt.Printf("there are %v pair(s) of socks in %v",pairs,socks)


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