package kata2

import _ "fmt"
import "strings"

/*
func rot13(i byte) byte {
    switch {
    case 65 <= i && i <= 77, 97 <= i && i <= 109:
        return i + 13
    case 78 <= i && i <= 100, 110 <= i && i <= 122:
        return i - 13
    }
    return i
}

func MakeKata(input string) string {
    length := len(input)
    counter := 0
    out := ""
    for {
        if counter == length {
            break
        }
        out += fmt.Sprintf("%c", rot13(input[counter]))
        counter += 1
    }
    return out
}

func rot13(i byte) byte {
    switch {
    case 65 <= i && i <= 77, 97 <= i && i <= 109:
        return i + 13
    case 78 <= i && i <= 100, 110 <= i && i <= 122:
        return i - 13
    }
    return i
}

func MakeKata(input string) string {
    length := len(input)
    var counter int
    for {
        if counter == length {
            break
        }
        input += string(rot13(input[counter]))
        counter += 1
    }
    return input[length:]
}

func rot13(s byte) byte {
    var s13 byte
    switch {
    case 65 <= s && s <= 77, 97 <= s && s <= 109:
        s13 = s + 13
    case 78 <= s && s <= 100, 110 <= s && s <= 122:
        s13 = s - 13
    default:
        s13 = s
    }
    return s13
}

func MakeKata(input string) string {
    length := len(input)
    for i := 0; i < length; i++ {
        input += string(rot13(input[i]))
    }
    return input[length:]

}
*/

func rot13(i byte) byte {
	switch {
	case 65 <= i && i <= 77, 97 <= i && i <= 109:
		i += 13
	case 78 <= i && i <= 100, 110 <= i && i <= 122:
		i -= 13
	}

	return i
}

func MakeKata(input string) string {
	length := len(input)
	out := make([]byte, length)
	br := strings.NewReader(input)
	br.Read(out)
	for pos, val := range out {
		out[pos] = rot13(val)
	}
	return string(out)
}
