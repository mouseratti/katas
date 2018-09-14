package house_numbers_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fixtures = []struct {
	in       []int64
	expected int64
}{
	{[]int64{1, 5, 7, 0}, 13},
	{[]int64{1, 7, 0, 1, 8}, 8},
	{[]int64{0, 1, 7, 0, 1, 8}, 0},
}

func Test_kata(t *testing.T) {
	for _, input := range fixtures {
		t.Run(fmt.Sprint(input.in), func(t *testing.T) {
			assert.Equal(t, input.expected, make_kata1(input.in), "kata1")
			assert.Equal(t, input.expected, make_kata2(input.in), "kata2")
			assert.Equal(t, input.expected, make_kata3(input.in...), "kata3")
		})
	}

}

func Benchmark_make_kata1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata1([]int64{1, 2, 3, 5, 7, 8, 9, 0, 4, 0, 5, 0})
	}
}
func Benchmark_make_kata2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata2([]int64{1, 2, 3, 5, 7, 8, 9, 0, 4, 0, 5, 0})
	}
}

func Benchmark_make_kata3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata3([]int64{1, 2, 3, 5, 7, 8, 9, 0, 4, 0, 5, 0}...)
	}
}
