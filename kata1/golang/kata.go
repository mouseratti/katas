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

func MakeKata(input string) string {
    input = strings.ToLower(input)
    var out []rune = []rune{}
    symbols := make(map[rune]rune)
    for _, sym := range input {
        switch symbols[sym] {
        case ')':
            continue
        case '(':
            symbols[sym] = ')'
        default:
            symbols[sym] = '('
        }
    }
    for _, sym := range input {
        out = append(out, symbols[sym])
    }
    return string(out)
}
