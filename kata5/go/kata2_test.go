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

func BenchmarkKata2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		make_kata2([]int64{10,9,8,7,6,5,4,3,2,1})
	}
}
