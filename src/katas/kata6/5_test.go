package kata6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"katas/fixtures"
)

func Test_make_kata5(t *testing.T) {
	f := new(fixtures.Fixture)
	f.Input = []int64{256,124,18,17,6543,122,765,5,3,8,2333}
	fmt.Println(f.Input)
	make_kata5(f)
	assert.True(t, f.Is_sorted(), fmt.Sprintf("%v is not sorted!", f.Input))
	fmt.Println(f.Input)
}

func Benchmark_make_kata5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := new(fixtures.Fixture)
		f.Input = []int64{256,124,18,17,122,765,5,3,8,2333}
		make_kata5(f)
	}
}