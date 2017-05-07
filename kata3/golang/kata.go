package kata3

import "strings"
import "strconv"
import "fmt"
import "math"

func MakeKata(ipaddress string) (subnet, broadcast string) {
    ipslice := strings.Split(ipaddress, "/")
    ip, strmask := ipslice[0], ipslice[1]
    mask, _ := strconv.Atoi(strmask)
    binary_ip := make([]bool, 32)
    binary_ip_index := 0
    for _, val := range strings.Split(ip, ".") {
        intval, _ := strconv.Atoi(val)
        binary_val := fmt.Sprintf("%8b", intval)
        for _, val := range binary_val {
            switch val {
            case '1':
                binary_ip[binary_ip_index] = true
            default:
                binary_ip[binary_ip_index] = false
            }
            binary_ip_index += 1
        }
    }
    // counting subnet
    octet_slice := make([]byte, 4)
    index := 0
    mult := 7.0
    for i := mask; i < 32; i++ {
        binary_ip[i] = false
    }
    for i := 0; i < 32; i++ {
        if binary_ip[i] {
            octet_slice[index] += byte(math.Pow(2.0, mult))
        }
        mult -= 1
        if mult < 0 {
            index += 1
            mult = 7.0
        }
    }
    subnet = fmt.Sprintf("%v.%v.%v.%v", octet_slice[0], octet_slice[1], octet_slice[2], octet_slice[3])

    // counting broadcast
    octet_slice = make([]byte, 4)
    index = 0
    mult = 7.0
    for i := mask; i < 32; i++ {
        binary_ip[i] = true
    }
    for i := 0; i < 32; i++ {
        if binary_ip[i] {
            octet_slice[index] += byte(math.Pow(2.0, mult))
        }
        mult -= 1
        if mult < 0 {
            index += 1
            mult = 7.0
        }
    }
    broadcast = fmt.Sprintf("%v.%v.%v.%v", octet_slice[0], octet_slice[1], octet_slice[2], octet_slice[3])

    return
}
