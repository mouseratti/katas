package binary_chop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_make_kata2(t *testing.T) {
	for _, f := range fixtures {
		t.Run(fmt.Sprint(f), func(t *testing.T) {
			assert.Equal(t, f.pos, make_kata2(f.num, f.rang))
		})
	}

}

func Benchmark_kata2(b *testing.B) {
	length := 10000
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = 5 * i
	}
	for i := 0; i < b.N; i++ {
		make_kata2(1, slice)
	}
}

func Benchmark_find_recursively(b *testing.B) {
	length := 10000
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = 5 * i
	}
	for i := 0; i < b.N; i++ {
		find_recursively(1, slice, 0, 9999)
	}
}
