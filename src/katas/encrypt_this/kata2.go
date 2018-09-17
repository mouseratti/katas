package encrypt_this

import (
	"fmt"
	"strings"
)

func make_kata2(in string) string {
	tokens := strings.Split(in, " ")
	encrypt_tokens(tokens)
	return strings.Join(tokens, " ")
}

func encrypt_tokens(tokens []string) {
	//FIXME
	for pos := range tokens {
		go encrypt_token(&tokens[pos])
	}
}

func encrypt_token(t *string) {
	clear_token(t)
	buffer := []rune(*t)
	maxElement := len(buffer) - 1
	buffer[1], buffer[maxElement] = buffer[maxElement], buffer[1]
	*t = fmt.Sprintf("%d%s", buffer[0], string(buffer[1:]))
}

func clear_token(t *string) {
	buffer := []rune(*t)
	maxElem := len(buffer) - 1
	for i := maxElem; i >= 0; i-- {
		if !_GARBAGE_SYMBOLS.contains(buffer[i]) {
			break
		}
		buffer = buffer[:i]
	}
	*t = string(buffer)
}
