package encrypt_this

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var garbage_fixtures = []struct {
	elem     interface{}
	set      garbage
	expected bool
}{
	{1, garbage{3, 2, 1}, true},
	{'.', _GARBAGE_SYMBOLS, true},
	{'!', _GARBAGE_SYMBOLS, true},
	{"ora", garbage{"que", "ora", "Es"}, true},
	{'1', _GARBAGE_SYMBOLS, false},
}

func Test_garbage_contains(t *testing.T) {
	for _, f := range garbage_fixtures {
		t.Run(fmt.Sprintf("%v", f),
			func(t *testing.T) {
				assert.Equal(t, f.expected, f.set.contains(f.elem))
			})
	}
}
