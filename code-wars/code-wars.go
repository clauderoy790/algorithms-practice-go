package code_wars

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CodeWars struct {}

func (cw *CodeWars) Resolve() {
	s1 := DNAStrand("ATTGC")
	s2 := DNAStrand("GTAT")
	fmt.Printf("s1: %v, s2: %v",s1,s2)
}

func chocolateRes() {
//a := []int{121,144,19,161,19,144,19,11}
//b := []int{121,14641,20736,361,25921,361,20736,361}
//c := []int{121,144,19,161,19,144,19,11}
//d := []int{231,14641,20736,361,25921,361,20736,361}
//e := []int{641,20736,36100,25921,361,20736,361}
g := []int {2,2,3}
h := []int {4,9,9}
fmt.Println("comp: ",Comp(g,h))
}

func Comp(a []int, b []int) bool {
	if a == nil || b == nil || len(b) != len(a) {
		return false
	}

	squared := make([]int, len(a)) // make new slice and append all a numbers squared
	for i, nb := range a {
		squared[i] = nb * nb
	}

	for _, nb := range b { // loop b elements until we find one that is not contained in squared
		if ind := indexOf(squared, nb); ind != -1 { // check if they are contained in b
			squared[ind] = squared[len(squared)-1] // replace index with last element of slice
			squared = append(squared[:ind], squared[ind+1:]...) //shrink array
		} else {
			return false
		}
	}
	return true
}

func indexOf(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func BreakChocolate(n, m int) int {
	if n <= 0 || m <= 0 {
		return 0
	}
	return n*m-1
}

func ip_res() {
	ips := []string{"1.1.1.1", "255.244.255.222", "11.23.32", "salut","1.2.3.4","123.45.67.89","1.2.3","1.2.3.4.5","123.456.78.90","123.045.067.089","0.34.82.53"}

	for _, ip := range ips {
		fmt.Printf("ip %v is valid %v\n", ip, is_valid_ip(ip))
	}
}

func is_valid_ip(ip string) bool {
	regStr := `^(0|([1-9]\d{0,2})).(0|([1-9]\d{0,2})).(0|([1-9]\d{0,2})).(0|([1-9]\d{0,2}))$`
	var validIP = regexp.MustCompile(regStr)

	validFormat := validIP.MatchString(ip)

	if validFormat {
		splitDigits := strings.Split(ip,".")
		for _,str := range splitDigits {
			digit, err := strconv.Atoi(str)
			if err != nil || digit < 0 || digit > 255 {
				return false
			}
		}
		return true
	}
	return false
}

var complements = map[rune]rune {
	'A':'T',
	'T':'A',
	'C':'G',
	'G':'C',
}

func DNAStrand(dna string) string {
	runes := []rune(dna)

	for i, r := range runes { // loop through each characters
		if c, ok := findComplement(r);ok {
			runes[i] = c
		}
	}
	return string(runes)
}

func findComplement(r rune) (compl rune,ok bool) {
	for k, v := range complements {
		if k == r {
			ok = true
			compl = v
		}
	}
	return
}
