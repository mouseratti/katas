package kata5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKata2(t *testing.T) {
	for _, f := range fixtures {
		f.reverseInput()
		assert.Equal(t, f.expected, make_kata2(f.input), "not equal")
	}

}
