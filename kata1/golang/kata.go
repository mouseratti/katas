package kata

import "strings"

func MakeKata(inputted string) string {
    lower := strings.ToLower(inputted)

    outputted := ""
    d := make(map[string]string)

    for _, symbol := range lower {
        symbol := string(symbol)
        if d[symbol] == "" {
            d[symbol] = "("
        } else {
            d[symbol] = ")"
        }
    }
    for _, symbol := range lower {
        symbol := string(symbol)
        outputted += d[symbol]
    }
    return outputted
}
