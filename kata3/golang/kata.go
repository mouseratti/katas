package kata

import "strings"
import "strconv"
import "math"
import "fmt"

func to_binary(ip string) (binary_ip [32]bool) {
    binIpSlice := strings.Split(ip, ".")
    var ipArray [4]int
    for pos, val := range binIpSlice {
        ipArray[pos], _ = strconv.Atoi(val)
    }
    binaryIpString := fmt.Sprintf("%08b%08b%08b%08b", ipArray[0], ipArray[1], ipArray[2], ipArray[3])
    for pos, val := range binaryIpString {
        if val == '1' {
            binary_ip[pos] = true
        }
    }
    return
}

func to_integer(binary_ip [32]bool) (ip string) {
    var ipArray [4]int
    mult := 0.0
    index := 0
    for pos, val := range binary_ip {
        switch pos {
        case 8, 16, 24:
            index += 1
            mult = 0
        }

        if val {
            ipArray[index] += int(math.Pow(2.0, 7-mult))
        }
        mult += 1
    }
    ip = fmt.Sprintf("%v.%v.%v.%v", ipArray[0], ipArray[1], ipArray[2], ipArray[3])
    return
}

func MakeKata(input string) (subnet, broadcast string) {
    ipSlice := strings.Split(input, "/")
    mask, _ := strconv.Atoi(ipSlice[1])
    ip := to_binary(ipSlice[0])
    binSubnet, binBroadcast := ip, ip
    for i := mask; i < 32; i++ {
        binSubnet[i] = false
        binBroadcast[i] = true
    }
    subnet = to_integer(binSubnet)
    broadcast = to_integer(binBroadcast)
    return
}
