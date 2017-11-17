package kata5

import (
	"testing"
	"katas/fixtures"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestMake_kata3(t *testing.T) {
	f := fixtures.GetFixture(30)
	make_kata3(&f)
	assert.True(t, f.Is_sorted(), fmt.Sprintf("%v is not sorted", f))
}


func Benchmark_make_kata3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := fixtures.Fixture{[]int64{100,59,888,67,16,892,11,34, 1111,12,7654,818}}
		make_kata3(&f)
	}
}