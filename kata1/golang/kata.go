package kata

import "strings"

// func MakeKata(inputted string) string {
//     lower := strings.ToLower(inputted)

//     outputted := ""
//     d := make(map[string]string)

//     for _, symbol := range lower {
//         symbol := string(symbol)
//         if d[symbol] == "" {
//             d[symbol] = "("
//         } else {
//             d[symbol] = ")"
//         }
//     }
//     for _, symbol := range lower {
//         symbol := string(symbol)
//         outputted += d[symbol]
//     }
//     return outputted
// }

// func MakeKata(inputted string) string {
//     lower := strings.ToLower(inputted)
//     outputted := ""
//     d := make(map[rune]rune)

//     for _, symbol := range lower {
//         if d[symbol] == 0 {
//             d[symbol] = '('
//         } else {
//             d[symbol] = ')'
//         }
//     }
//     for _, symbol := range lower {
//         outputted += string(d[symbol])
//     }
//     return outputted
// }

func in_string(symbol rune, str string) bool {
    for _, s := range str {
        if symbol == s {
            return true
        }
    }
    return false
}

func remove_from_string(symbol rune, str *string) bool {
    for pos, sym := range &str {
        if symbol == sym {
            a = a[:pos-1] + a[pos]
            return true
        }
    }
    return false
}

func MakeKata(input string) string {
    uniq := ""
    not_uniq := ""
    out := ""
    lower = strings.ToLower(input)

    for _, symbol := range lower {
        if in_string(symbol, uniq) {
            remove_from_string(symbol, *uniq)
            if !in_string(symbol, not_uniq) {
                not_uniq += string(symbol)
            }
        } else {
            uniq += string(symbol)
        }
    }

    for _, symbol := range lower {
        if in_string(symbol, uniq) {
            out += "("
        } else {
            out += ")"
        }
    }

    return out
}
