package kata

// import "strings"

import f "fmt"

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

func RemoveFromString(str string, symbol rune, count int) string {
    index := 0
    max_length := len(str)
    concat := func(inp *string, ind int) {
        inp_copy := *inp
        if ind == 0 {
            inp_copy = inp_copy[ind+1:]
        } else {
            inp_copy = inp_copy[:ind-1] + inp_copy[:ind+1]
        }
        *inp = inp_copy
    }

    for {
        if symbol == rune(str[index]) {
            concat(&str, index)
        }
        index += 1
        if index >= len(str) || index >= max_length {
            break
        }
        f.Println(str)
    }
    return str
}

// func MakeKata(input string) string {
//     uniq := ""
//     not_uniq := ""
//     out := ""
//     lower = strings.ToLower(input)

//     for _, symbol := range lower {
//         if in_string(symbol, uniq) {
//             remove_from_string(symbol, *uniq)
//             if !in_string(symbol, not_uniq) {
//                 not_uniq += string(symbol)
//             }
//         } else {
//             uniq += string(symbol)
//         }
//     }

//     for _, symbol := range lower {
//         if in_string(symbol, uniq) {
//             out += "("
//         } else {
//             out += ")"
//         }
//     }

//     return out
// }
