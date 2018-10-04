package kata5

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func Test_is_sorted(t *testing.T) {
	var arrays [][]int = [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 3, 4, 5, 5, 6},
		{5, 1, 2, 3, 4, 5, 6},
		{1, 2, 3, 1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6, 5},
	}
	var sortedValues []bool = []bool{true, true, false, false, false}
	for pos, val := range arrays {
		assert.Equal(t, sortedValues[pos], is_sorted(val), "not sorted!")

	}

}

func Test_make_kata5(t *testing.T) {
	fixtures := [][]int{
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 7, 6, 1},
		{10, 9, 8, 7, 6, 5, 344, 3, 888992, 1, 0, 7, 6, 10000000000000},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 7, 6, 1},
		{1110, 9, 8, 1117, 22226, 5, 4, 3, 2, 1, 0, 7, 6, 1},
		{10, 9, 8, 7, 6, 57777777, 4, 3, 2, 1, 0, 7, 6, 1},
		{10, 9, -8, -7, 6, -57777777, 4, -3, 2, 1, 0, 7, 6, 1},
		{10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355},
		{10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355, -400},
	}
	for _, val := range fixtures {
		make_kata5(val)
		assert.True(t, is_sorted(val), fmt.Sprintf("%v is not sorted properly!", val))
	}
}

func Test_kata55(t *testing.T) {
	inp := []int{10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355}
	expected := []int{-355, -100, -2, -1, 0, 1, 4, 5, 6, 7, 8, 9, 10, 355, 2848}
	assert.Equal(t, expected, make_kata55(inp), "not equal!")
}
