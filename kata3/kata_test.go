package kata3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKata(t *testing.T) {
	for _, fix := range fixtures {
		subnet, broadcast := MakeKata(fix.input)
		assert.Equal(t, fix.subnet, subnet)
		assert.Equal(t, fix.broadcast, broadcast)
	}
}

func BenchmarkKata(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeKata("109.130.40.1/28")
	}
}
