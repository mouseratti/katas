package kata5

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKata5(t *testing.T) {

	for _, val := range fixtures {

		assert.Equal(t, val.expected, make_kata(val.input), "They should be equal")
	}

}


func Test_push_buble(t *testing.T) {
	for _, f := range push_buble_fixtures {
		push_buble(f.input_slice, f.position)
		assert.Equal(t, f.expected_slice, f.input_slice, "not equal")
	}
}