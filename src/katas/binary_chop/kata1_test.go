package binary_chop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kata1(t *testing.T) {
	for _, v := range fixtures {
		t.Run(fmt.Sprint(v), func(t *testing.T) {
			assert.Equal(t, v.pos, make_kata1(v.rang, v.num))
		})
	}
}
func Benchmark_kata1(b *testing.B) {
	length := 10000
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = 5 * i
	}
	for i := 0; i < b.N; i++ {
		make_kata1(slice, 1)
	}
}
