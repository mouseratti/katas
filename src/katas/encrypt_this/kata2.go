package encrypt_this

import (
	"fmt"
	"strings"
	"sync"
)

func make_kata2(in string) string {
	tokens := strings.Split(in, " ")
	concurrent_encrypt_tokens(tokens)
	return strings.Join(tokens, " ")
}

func encrypt_tokens(tokens []string) {
	for pos := range tokens {
		encrypt_token(&tokens[pos])
	}
	return
}

func concurrent_encrypt_tokens(tokens []string) {
	anchor := new(sync.WaitGroup)
	anchor.Add(len(tokens))
	for pos := range tokens {
		val := pos
		go func() {
			encrypt_token(&tokens[val])
			anchor.Done()
		}()
	}
	anchor.Wait()
}

func encrypt_token(t *string) {
	clear_token(t)
	buffer := []rune(*t)
	maxElement := len(buffer) - 1
	buffer[1], buffer[maxElement] = buffer[maxElement], buffer[1]
	*t = fmt.Sprintf("%d%s", buffer[0], string(buffer[1:]))
	return
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
