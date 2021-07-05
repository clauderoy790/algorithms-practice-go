package ip_validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Resolve() {
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
