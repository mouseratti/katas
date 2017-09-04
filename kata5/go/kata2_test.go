package kata5

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKata2(t *testing.T)  {
	for _, f := range fixtures {
		f.reverseInput()
		assert.Equal(t, f.expected, make_kata2(f.input), "not equal")
	}

}
