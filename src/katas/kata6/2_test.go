package kata6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"katas/fixtures"
)

func Test_make_kata2(t *testing.T) {
	f := fixtures.GetFixturePointer(100)
	fmt.Println(f.Input)
	make_kata2(f)
	assert.True(t, f.Is_sorted(), fmt.Sprintf("%v is not sorted!", f.Input))
	fmt.Println(f.Input)
}

func Benchmark_make_kata2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f:= fixtures.GetFixturePointer(100)
		make_kata2(f)
	}
}