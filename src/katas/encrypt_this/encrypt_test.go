package encrypt_this

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_make_kata1(t *testing.T) {
	for _, f := range fixtures {
		t.Run(f.input, func(t *testing.T) {
			assert.Equal(t, f.expected, make_kata1(f.input), f.input)
		})
	}

}
func Test_encrypt_word(t *testing.T) {
	for _, f := range fixtures_encrypt_word {
		t.Run(f.input, func(t *testing.T) {
			assert.Equal(t, f.expected, encrypt_word(f.input), f.input)
		})
	}

}

func Test__trim_unused_symbols(t *testing.T) {
	for _, f := range fixtures_unused_symbols {
		t.Run(f.input, func(t *testing.T) {
			trim_unused_symbols(&f.input)
			assert.Equal(t, f.expected, f.input, f.input)
		})
	}

}

func Test_isInSet(t *testing.T) {
	for pos, f := range fixtures_isInSet {
		t.Run(fmt.Sprintf("Test %v", pos), func(t *testing.T) {
			assert.Equal(
				t,
				f.expected,
				isInSet(f.element, f.set...),
				fmt.Sprintf("Test %v", pos))
		})
	}

}

func Benchmark_isInSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isInSet('!', 1, 2, 5, 6, 7, 8, '!', '2', "22")
	}

}

func Benchmark_make_kata1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata1(fixtures[0].input)
	}

}
