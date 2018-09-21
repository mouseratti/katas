package binary_chop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fixtures = []struct {
	num  int
	rang []int
	pos  int
}{
	{5, []int{1, 2, 3}, -1},
	{5, []int{1, 2, 3, 4, 5}, 4},
	{5, []int{5, 15, 25, 35, 135, 250, 300}, 0},
	{5, []int{1, 2, 4, 5, 13, 24, 245, 333}, 3},
	{99, []int{1, 2, 4, 5, 13, 15, 18, 20, 24, 55, 66, 77, 88, 99, 111, 123, 245, 333}, 13},
}

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
