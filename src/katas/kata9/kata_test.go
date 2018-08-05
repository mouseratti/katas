package kata9

import "testing"
import "github.com/stretchr/testify/assert"

var bigint64 int64 = int64(999999999999999)
var fixtures [][2][]int64 = [][2][]int64{
	{[]int64{1, 2, 3}, []int64{2, 4, 6}},
	{[]int64{2, 3, 4}, []int64{4, 6, 8}},
	{[]int64{3, 2, 1, 0}, []int64{6, 4, 2, 0}},
	{[]int64{0, bigint64, 2, 1, 0}, []int64{0, bigint64 * 2, 4, 2, 0}},
}

func Test_make_kata(t *testing.T) {

	for _, expected := range fixtures {
		result := make_kata(expected[0]...)
		assert.Equal(t, expected[1], result)
	}
}

func Test_make_kata2(t *testing.T) {

	for _, expected := range fixtures {
		result := make_kata(expected[0]...)
		assert.Equal(t, expected[1], result)
	}
}
