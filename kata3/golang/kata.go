package kata

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func MakeKata(ipaddr string) (s, b string) {
    toBinary := func(ip string) (binaryIp [32]bool) {
        ipslice := strings.Split(ip, ".")
        var ipInteger [4]int
        for p, v := range ipslice {
            ipInteger[p], _ = strconv.Atoi(v)
        }
        for p, v := range fmt.Sprintf(
            "%08b%08b%08b%08b",
            ipInteger[0],
            ipInteger[1],
            ipInteger[2],
            ipInteger[3]) {
            if v == '1' {
                binaryIp[p] = true
            }
        }
        return
    }

    toInt := func(binaryIp [32]bool) (ip string) {
        var ipInteger [4]float64
        index := 0
        mult := 7.0
        for p, v := range binaryIp {
            switch p {
            case 8, 16, 24:
                index += 1
                mult = 7
            }
            if v {
                ipInteger[index] += math.Pow(2.0, mult)
            }
            mult -= 1
        }
        ip = fmt.Sprintf("%v.%v.%v.%v", ipInteger[0], ipInteger[1], ipInteger[2], ipInteger[3])
        return
    }
    ipSlice := strings.Split(ipaddr, "/")
    binaryIp := toBinary(ipSlice[0])
    mask, _ := strconv.Atoi(ipSlice[1])
    binarySubnet, binaryBroadcast := binaryIp, binaryIp
    for i := mask; i < 32; i++ {
        binarySubnet[i], binaryBroadcast[i] = false, true
    }
    s, b = toInt(binarySubnet), toInt(binaryBroadcast)
    return
}
