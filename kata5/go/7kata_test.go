package kata5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	fix := []MyInt{2, 1, 4, 5, 6, 7, -5, -11, 16}
	assert.Equal(t,
		[]MyInt{-11, -5, 1, 2, 4, 5, 6, 7, 16},
		Sort(fix),
		"not equal!!",
	)
}
