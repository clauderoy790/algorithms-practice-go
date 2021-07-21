package hacker_rank_thirty

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HackerRankThirty struct{}

func (hr *HackerRankThirty) Resolve() {
	dayNine()
}

func scannerExample() {
	scanner := bufio.NewScanner(os.Stdin)

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	scanner.Scan()
	nb, _ := strconv.ParseUint(scanner.Text(), 10, 8)
	scanner.Scan()
	fl, _ := strconv.ParseFloat(scanner.Text(), 8)
	scanner.Scan()
	str := scanner.Text()

	nb2 := nb + i
	nb3 := fl + d
	fmt.Println(nb2)
	fmt.Printf("%.1f\n", nb3)
	fmt.Printf("%v%v\n", s, str)
}

func moduloString() {
	fmt.Println("enter nb")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < nb; i++ {
		beg, end := "", ""
		scanner.Scan()
		text := scanner.Text()
		for i, r := range text {
			if i%2 == 0 {
				beg += string(r)
			} else {
				end += string(r)
			}
		}
		fmt.Printf("%v %v\n", beg, end)
	}
}

func daySeven() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	//nb, err := strconv.Atoi(scanner.Text())
	//if err != nil {
	//	panic(err)
	//}
	//sl := make([]int,nb)

	scanner.Scan()
	strs := strings.Split(scanner.Text()," ")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	fmt.Println(strings.Join(strs, " "))
}

func dayEight() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nb,err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("error scanning number: ",err)
	}

	m := make(map[string]int)
	for i := 0;i < nb; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text()," ")
		if len(sl) != 2 {
			panic("need to only have one space in the string")
		}
		v, err := strconv.Atoi(sl[1])
		if err != nil {
			panic(err)
		}
		m[sl[0]] = v
	}

	for scanner.Scan(){
		key := scanner.Text()
		if len(key) == 0 {
			break
		}
		if v, ok := m[key]; ok {
			fmt.Printf("%v=%v\n",key,v)
		} else {
			fmt.Println("Not found")
		}

	}
}

func dayNine() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	var fac func(n int) int
	fac = func(n int) int {
		if n <= 0 {
			return 0
		} else if n == 1 {
			return 1
		} else {
			return n* fac(n-1)
		}
	}

	res := fac(nb)
	fmt.Println(res)
}
