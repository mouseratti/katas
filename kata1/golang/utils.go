package kata

import "sort"

func in_string(symbol rune, str string) bool {
    for _, s := range str {
        if symbol == s {
            return true
        }
    }
    return false
}

func cut(inp *string, ind int) {
    var mapa map[int]rune = map[int]rune{}
    for pos, s := range *inp {
        if pos == ind {
            continue
        }
        mapa[pos] = s
    }
    *inp = ""
    var keys []int
    for key := range mapa {
        keys = append(keys, key)
    }
    sort.Ints(keys)
    for _, key := range keys {
        *inp += string(mapa[key])
    }
}

func RemoveFromString(str string, symbol rune, count int) string {

    return ""
}
