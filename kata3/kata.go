package kata3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type IpAddress struct {
	ip   [4]int
	mask int
}

func (self IpAddress) Stringer() string {
	return fmt.Sprintf("%v.%v.%v.%v/%v", self.ip[0], self.ip[1], self.ip[2], self.ip[3], self.mask)
}

func (self IpAddress) Binary() [32]bool {
	var result [32]bool
	ipString := fmt.Sprintf("%08b%08b%08b%08b", self.ip[0], self.ip[1], self.ip[2], self.ip[3])
	for i := 0; i < 32; i++ {
		if ipString[i] == '1' {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func (self IpAddress) Subnet() IpAddress {
	binaryIp := self.Binary()
	for i := self.mask; i < 32; i++ {
		binaryIp[i] = false
	}
	return fromBinary(binaryIp, self.mask)
}

func (self IpAddress) Broadcast() IpAddress {
	binaryIp := self.Binary()
	for i := self.mask; i < 32; i++ {
		binaryIp[i] = true
	}
	return fromBinary(binaryIp, self.mask)
}

func FromString(in string) IpAddress {
	var ip [4]int
	ipSlice := strings.Split(in, "/")
	mask, _ := strconv.Atoi(ipSlice[1])
	for pos, val := range strings.Split(ipSlice[0], ".") {
		ip[pos], _ = strconv.Atoi(val)
	}
	return IpAddress{ip, mask}
}

func fromBinary(ip [32]bool, mask int) IpAddress {
	var integerIp [4]int
	mult := 7.0
	index := 0
	for pos, val := range ip {
		switch pos {
		case 8, 16, 24:
			mult = 7.0
			index += 1
		}
		if val {
			integerIp[index] += int(math.Pow(2.0, mult))
		}
		mult -= 1
	}
	return IpAddress{integerIp, mask}
}

func MakeKata(ipaddr string) (string, string) {
	ip := FromString(ipaddr)
	subnet := ip.Subnet()
	broadcast := ip.Broadcast()
	return subnet.Stringer(), broadcast.Stringer()
}
