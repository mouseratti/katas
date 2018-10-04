package kata5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_3kata(t *testing.T) {
	for _, f := range fixtures {
		f.reverseInput()
		assert.Equal(t, f.expected, make_kata3(f.input), "not equal")
		assert.Equal(t,
			[]int64{0, 1, 4, 5, 5, 7, 8, 8, 8, 10, 11},
			make_kata3([]int64{7, 5, 8, 4, 10, 8, 5, 8, 1, 11, 0}),
			"not equal",
		)
	}
}
