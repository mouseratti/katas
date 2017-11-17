package kata5

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"katas/fixtures"
	"fmt"
)

func TestMake_kata4(t *testing.T) {

	f:= fixtures.GetFixturePointer(8)
	fmt.Println(f.Input)
	make_kata4(f)
	fmt.Println(f.Input)
	assert.True(t, f.Is_sorted())
}


func Benchmark_make_kata4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := fixtures.Fixture{[]int64{100,59,888,67,16,892,11,34, 1111,12,7654,818}}
		make_kata4(&f)
	}
}