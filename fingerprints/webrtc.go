package fingerprint

import (
	"fmt"
	"math/rand"
)

func GenerateRandomIPv6() string {
	ip := make([]byte, 16)
	for i := 0; i < len(ip); i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", ip[0:2], ip[2:4], ip[4:6], ip[6:8], ip[8:]) + ".local"
}
