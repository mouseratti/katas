package fixtures

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFixture_Is_sorted(t *testing.T) {
	f := new(Fixture)
	f.Input = []int64{3,2,1}
	assert.False(t, f.Is_sorted())
	f.Input = []int64{5,6,7}
	assert.True(t, f.Is_sorted())
}
