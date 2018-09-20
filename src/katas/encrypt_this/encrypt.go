package encrypt_this

import (
	"fmt"
	"strings"
)

var unused_symbols = []interface{}{
	'!',
	',',
	'.',
}

func make_kata1(in string) string {
	words := strings.Split(in, " ")
	for pos, w := range words {
		words[pos] = encrypt_word(w)
	}
	return strings.Join(words, " ")
}

func encrypt_word(s string) string {
	trim_unused_symbols(&s)
	inputLength := len(s)
	bufferLength := inputLength - 1
	buffer := make([]byte, bufferLength)
	buffer[0] = s[inputLength-1]
	buffer[bufferLength-1] = s[1]
	for i := 1; i < bufferLength-1; i++ {
		buffer[i] = s[i+1]
	}
	return fmt.Sprintf("%d%s", s[0], buffer)
}

func trim_unused_symbols(s *string) {
	buffer := []rune(*s)
	max := len(*s) - 1
	for i := max; i >= 0; i-- {
		if isInSet(buffer[i], unused_symbols...) {
			*s = string(buffer[:i])
		} else {
			break
		}
	}
	return
}

func isInSet(elem interface{}, set ...interface{}) bool {
	for _, val := range set {
		if elem == val {
			return true
		}
	}
	return false
}
