package encrypt_this

import (
	"fmt"
	"strings"
)

func make_kata1(in string) string {
	words := strings.Split(in, " ")
	for pos, w := range words {
		words[pos] = encrypt_word(w)
	}
	return fmt.Sprintf("%v", strings.Join(words, " "))
}

func encrypt_word(s string) string {
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
