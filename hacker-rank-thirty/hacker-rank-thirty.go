package hacker_rank_thirty

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HackerRankThirty struct {}

func (hr *HackerRankThirty) Resolve() {
	scanner := bufio.NewScanner(os.Stdin)

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	scanner.Scan()
	nb,_ := strconv.ParseUint(scanner.Text(),10,8)
	scanner.Scan()
	fl,_ := strconv.ParseFloat(scanner.Text(),8)
	scanner.Scan()
	str := scanner.Text()

	nb2 := nb + i
	nb3 := fl +d
	fmt.Println(nb2)
	fmt.Printf("%.1f\n",nb3)
	fmt.Printf("%v%v\n",s,str)

}