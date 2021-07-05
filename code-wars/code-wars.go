package code_wars

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CodeWars struct {}

func (cw *CodeWars) Resolve() {
	chocolate_res()
}

func chocolate_res() {
a := []int {8, -89, -1, 12, 48, -55, -89, 28, -59, -11, 50}
b := []int {1, 64, 121, 144, 784, 2304, 2500, 3025, 3481, 7921, 7921}
fmt.Println("comp: ",Comp(a,b))
}

func Comp(a []int, b []int) bool {
	if a == nil || b == nil || len(b) == 0 {
		return false
	}

	squared := make([]int,len(a))
	for i, nb := range a {
		squared[i] = nb*nb
	}
	fmt.Printf("a: %v, b: %v",a,b)

	for _, nb := range b {
		if !contains(squared,nb) {
			return false
		} else {

		}
	}
	return true
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
