package kata3

import "strings"
import "strconv"
import "fmt"
import "math"

// func MakeKata(ipaddress string) (subnet, broadcast string) {
//     ipslice := strings.Split(ipaddress, "/")
//     ip, strmask := ipslice[0], ipslice[1]
//     mask, _ := strconv.Atoi(strmask)
//     binary_ip := make([]bool, 32)
//     binary_ip_index := 0
//     for _, val := range strings.Split(ip, ".") {
//         intval, _ := strconv.Atoi(val)
//         binary_val := fmt.Sprintf("%8b", intval)
//         for _, val := range binary_val {
//             switch val {
//             case '1':
//                 binary_ip[binary_ip_index] = true
//             default:
//                 binary_ip[binary_ip_index] = false
//             }
//             binary_ip_index += 1
//         }
//     }
//     // counting subnet
//     octet_slice := make([]byte, 4)
//     index := 0
//     mult := 7.0
//     for i := mask; i < 32; i++ {
//         binary_ip[i] = false
//     }
//     for i := 0; i < 32; i++ {
//         if binary_ip[i] {
//             octet_slice[index] += byte(math.Pow(2.0, mult))
//         }
//         mult -= 1
//         if mult < 0 {
//             index += 1
//             mult = 7.0
//         }
//     }
//     subnet = fmt.Sprintf("%v.%v.%v.%v", octet_slice[0], octet_slice[1], octet_slice[2], octet_slice[3])

//     // counting broadcast
//     octet_slice = make([]byte, 4)
//     index = 0
//     mult = 7.0
//     for i := mask; i < 32; i++ {
//         binary_ip[i] = true
//     }
//     for i := 0; i < 32; i++ {
//         if binary_ip[i] {
//             octet_slice[index] += byte(math.Pow(2.0, mult))
//         }
//         mult -= 1
//         if mult < 0 {
//             index += 1
//             mult = 7.0
//         }
//     }
//     broadcast = fmt.Sprintf("%v.%v.%v.%v", octet_slice[0], octet_slice[1], octet_slice[2], octet_slice[3])

//     return
// }

type Ipv4address struct {
    ip   [4]int
    mask int
}

func (self Ipv4address) to_binary() (out [32]bool) {
    binary_ip_string := fmt.Sprintf("%8b%8b%8b%8b", self.ip[0], self.ip[1], self.ip[2], self.ip[3])
    for pos, bit := range binary_ip_string {
        if bit == '1' {
            out[pos] = true
        }
    }
    return
}

func (self Ipv4address) count(flag bool) (out string) {
    binary_ip := self.to_binary()
    for i := self.mask; i < 32; i++ {
        binary_ip[i] = flag
    }
    octets := from_binary(binary_ip)
    out = fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
    return
}

func from_binary(binary_ip [32]bool) (out [4]int) {
    index := 0
    mult := 7.0
    for pos, bit := range binary_ip {
        switch pos {
        case 8, 16, 24:
            index += 1
            mult = 7
        }
        if bit {
            out[index] += int(math.Pow(2.0, mult))
        }
        mult -= 1

    }
    return
}

func to_octets(ip string) (out [4]int) {
    octets := strings.Split(ip, ".")
    for pos, octet := range octets {
        out[pos], _ = strconv.Atoi(octet)
    }
    return
}

func MakeKata(ipaddress string) (string, string) {
    ipslice := strings.Split(ipaddress, "/")
    mask, _ := strconv.Atoi(ipslice[1])
    ip := Ipv4address{to_octets(ipslice[0]), mask}
    return ip.count(false), ip.count(true)
}
