package kata5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_make_kata(t *testing.T) {
	a := []bool{true, false, true, false, false}
	Make_kata(a)
	assert.Len(t, a, 5)
	assert.Equal(t, []bool{false, false, false, true, true}, a)
}

func BenchmarkMake_kata(b *testing.B) {
	a := make([]bool, 500)
	for i := 0; i < 500; i++ {
		a[i] = rand.Float64() > 0.5
	}
	Make_kata(a)
	fmt.Println(a)
}
