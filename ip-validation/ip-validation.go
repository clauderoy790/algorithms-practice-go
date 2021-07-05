package ip_validation

import "fmt"

func Resolve() {
	ips := []string{"1.1.1.1", "255.244.255.222", "11.23.32", "salut"}

	for _, ip := range ips {
		fmt.Printf("ip %v is valid %v\n", ip, is_valid_ip(ip))
	}
}

func is_valid_ip(ip string) bool {
	return false
}
