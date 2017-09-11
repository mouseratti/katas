package kata6

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_kata(t *testing.T) {
	for _, fix := range kataFixtureList {
		make_kata(fix)
		assert.True(t, is_sorted(fix), fmt.Sprintf("%v is not sorted properly", fix))
	}
}

func Test_changePosition(t *testing.T) {
	for _, f := range changePositionFixtureList {
		changePosition(f.input, f.positions[0], f.positions[1])
		assert.Equal(t, f.expected, f.input, "not equal")

	}

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