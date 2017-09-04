package kata5

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestFixture(t *testing.T) {
	f:= fixture{[]int64{3,2,1},[]int64{1,2,3}}
	f.reverseInput()
	assert.Equal(t, f.expected, f.input, "not equal" )
}