package kata5

import "testing"
import "github.com/stretchr/testify/assert"

func Test_kata5(t *testing.T) {
	inp := []int{10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355}
	expected := []int{-355, -100, -2, -1, 0, 1, 4, 5, 6, 7, 8, 9, 10, 355, 2848}
	assert.Equal(t, expected, make_kata5(inp), "not equal!")
}
