package buble_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fixtures = []struct{
	input []int
	expected []int
	reversed bool
}{
	{[]int{500,100,676,1}, []int{1,100,500,676}, false},
	{[]int{500,100,676,1}, []int{676,500,100,1}, true},
	{[]int{7,9,500,1,2,3,4,5}, []int{500,9,7,5,4,3,2,1}, true},
	{[]int{7,9,500,1,2,3,5,4}, []int{1,2,3,4,5,7,9,500}, false},
	{[]int{3,1,4,2,1}, []int{1,1,2,3,4}, false},
	{[]int{3,1,4,2,1}, []int{4,3,2,1,1}, true},
}




func TestKata1(t *testing.T) {
	for _, val := range fixtures {
		t.Run(fmt.Sprint(val.input), func(t *testing.T) {
			Kata1(val.input, val.reversed)
			assert.Equal(t, val.expected, val.input)
		})
	}
}

func TestKata2(t *testing.T) {
	for _, val := range fixtures {
		t.Run(fmt.Sprint(val.input), func(t *testing.T) {
			Kata2(val.input, val.reversed)
			assert.Equal(t, val.expected, val.input)
		})
	}
}


func BenchmarkKata1(b *testing.B) {
	fixture:= fixtures[0]
	for i:=0; i< b.N; i++ {
		Kata1(fixture.input, fixture.reversed)
	}
}


func BenchmarkKata2(b *testing.B) {
	fixture:= fixtures[0]
	for i:=0; i< b.N; i++ {
		Kata2(fixture.input, fixture.reversed)
	}
}