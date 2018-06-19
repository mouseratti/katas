package kata8

import (
	"fmt"
	"katas/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kata8(t *testing.T) {
	fixture := fixtures.GetFixture(99)
	make_kata(fixture.Input)
	assert.True(t, fixture.Is_sorted(), fmt.Sprintf("%v is not sorted!", fixture))
}

func TestMakeKataNum(t *testing.T) {
	f := []interface{}{"1", 2, 3}
	f2, f3 := []interface{}{3, 2.0, 100}, []interface{}{2.0, 3, 100}
	assert.Panics(t, func() { makeKataNum(f) })
	makeKataNum(f2)
	assert.Equal(t, f2, f3)
}

func Test_isSorted(t *testing.T) {
	fixture := []struct {
		s []int64
		b bool
	}{
		{[]int64{1, 2, 3, 5, 7}, true},
		{[]int64{9, 1, 2, 3, 5, 7}, false},
		{[]int64{1, 2, 3, 5, 8}, true},
	}

	for _, val := range fixture {
		assert.Equal(
			t,
			isSorted(val.s),
			val.b,
			fmt.Sprintf("slice %v is not sorted!", val.s),
		)
	}
}

func Test_make_kata2(t *testing.T) {
	fixtures := [3][]int64{
		{5, 11, 3, 12, 8, 16},
		{3, 2, 1, 8, 4, 0},
		{115, 1, 113, 2, 82, 5},
	}
	for _, v := range fixtures {
		make_kata2(v)
		assert.True(t, isSorted(v), fmt.Sprintf("%v is not sorted", v))
	}
}
