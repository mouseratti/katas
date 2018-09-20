package encrypt_this

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encrypt_token(t *testing.T) {
	for _, f := range fixtures_encrypt_word {
		t.Run(fmt.Sprintf("%v to %v", f.input, f.expected),
			func(t *testing.T) {
				encrypt_token(&f.input)
				assert.Equal(t, f.expected, f.input, f.expected)
			})
	}

}

func Test_encrypt_tokens(t *testing.T) {
	var fixtures_encrypt_tokens_copy []struct {
		input    []string
		expected []string
	}
	copy(fixtures_encrypt_tokens_copy, fixtures_encrypt_tokens)
	for _, f := range fixtures_encrypt_tokens_copy {
		t.Run(fmt.Sprintf("%v to %v", f.input, f.expected),
			func(t *testing.T) {
				encrypt_tokens(f.input)
				assert.Equal(t, f.expected, f.input)
			})
	}

}

func Test_concurrent_encrypt_tokens(t *testing.T) {
	var fixtures_encrypt_tokens_copy []struct {
		input    []string
		expected []string
	}
	copy(fixtures_encrypt_tokens_copy, fixtures_encrypt_tokens)
	for _, f := range fixtures_encrypt_tokens_copy {
		t.Run(fmt.Sprintf("%v to %v", f.input, f.expected),
			func(t *testing.T) {
				concurrent_encrypt_tokens(f.input)
				assert.Equal(t, f.expected, f.input)
			})
	}

}

func Test_clear_token(t *testing.T) {
	for _, f := range fixtures_unused_symbols {
		t.Run(fmt.Sprintf("%v to %v", f.input, f.expected),
			func(t *testing.T) {
				clear_token(&f.input)
				assert.Equal(t, f.expected, f.input, f.expected)
			})
	}

}

func Test_make_kata2(t *testing.T) {
	for _, f := range fixtures {
		t.Run(f.input, func(t *testing.T) {
			assert.Equal(t, f.expected, make_kata2(f.input), f.input)
		})
	}

}

func Benchmark_kata2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata2(fixtures[0].input)
	}
}

func Benchmark_encrypt_tokens(b *testing.B) {
	str := strings.Split(longString, " ")
	for i := 0; i < b.N; i++ {
		encrypt_tokens(str)
	}
}

func Benchmark_concurrent_encrypt_tokens(b *testing.B) {
	str := strings.Split(longString, " ")
	for i := 0; i < b.N; i++ {
		concurrent_encrypt_tokens(str)
	}
}
