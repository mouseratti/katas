package kata6

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func Test_kata(t *testing.T) {
	kataFixture.Sort()
	assert.Equal(t, kataFixture.expected, kataFixture.input, "not equal")
}

func Test_changePosition(t *testing.T) {
	f := getFixture()
	f.changePosition(1, 5)
	assert.Equal(t, f.input, []int{6, 4, 3, 2, 1, 5})
	f.changePosition(4, 2)
	assert.Equal(t, f.input, []int{6, 4, 1, 3, 2, 5})
	f.changePosition(0, 5)
	assert.Equal(t, f.input, []int{4, 1, 3, 2, 5, 6})
	f.changePosition(0, 5)
	assert.Equal(t, f.input, []int{1, 3, 2, 5, 6, 4})
	f.changePosition(5, 0)
	assert.Equal(t, f.input, []int{4, 1, 3, 2, 5, 6})
	f.changePosition(0, 0)
	assert.Equal(t, f.input, []int{4, 1, 3, 2, 5, 6})
	f.changePosition(5, 5)
	assert.Equal(t, f.input, []int{4, 1, 3, 2, 5, 6})

}

func Test_getNextPos(t *testing.T) {
	fs := [3][3]int{
		{1, 5, 2},
		{8, 4, 7},
		{10, 1, 9},
	}
	for _, f := range fs {
		assert.Equal(t, f[2], getNextPos(f[0], f[1]), "got wrong position!")
	}
}