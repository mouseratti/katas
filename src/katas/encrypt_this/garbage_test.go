package encrypt_this

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var garbage_fixtures = []struct {
	elem     interface{}
	set      garbage
	expected bool
}{
	{1, garbage{3, 2, 1}, true},
	{'.', _GARBAGE_SYMBOLS, true},
	{'!', _GARBAGE_SYMBOLS, true},
	{"ora", garbage{"que", "ora", "Es"}, true},
	{'1', _GARBAGE_SYMBOLS, false},
}

func Test_garbage_contains(t *testing.T) {
	for _, f := range garbage_fixtures {
		t.Run(fmt.Sprintf("%v", f),
			func(t *testing.T) {
				assert.Equal(t, f.expected, f.set.contains(f.elem))
			})
	}
}

func Test_makeGarbageMap(t *testing.T) {
	expected := _GarbageMap{1: struct{}{}, 2: struct{}{}, "3": struct{}{}}
	input := []interface{}{1, 2, "3"}
	assert.Equal(t, expected, makeGarbageMap(input))
}

func Benchmark_garbage_contains(b *testing.B) {
	elem := "1!!"
	for i := 0; i < b.N; i++ {
		largeSet.contains(elem)
	}
}

func Benchmark_garbageMap_contains(b *testing.B) {
	elem := "1!!"
	gm := makeGarbageMap(largeSet)
	for i := 0; i < b.N; i++ {
		gm.contains(elem)
	}
}
