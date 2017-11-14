package kata6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"katas/fixtures"
)

func Test_make_kata1(t *testing.T) {
	f := fixtures.GetFixture(100)
	fmt.Println(f.Input)
	make_kata1(&f)
	assert.True(t, f.Is_sorted(), fmt.Sprintf("%v is not sorted!", f.Input))
	fmt.Println(f.Input)
}

func Benchmark_make_kata1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		f:= fixtures.GetFixture(100)
		make_kata1(&f)
	}
}