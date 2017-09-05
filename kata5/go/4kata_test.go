package kata5

import "testing"
import "github.com/stretchr/testify/assert"

var fixtures4 []map[string][]int = []map[string][]int{
	map[string][]int{"input":[]int{4, 8, 5, 9, 1, 11, 2, 12}, "expected":[]int{1, 2, 4, 5, 8, 9, 11, 12}},
	map[string][]int{"input":[]int{12,12,11,10,8,9,10,11,5,4,4,5}, "expected":[]int{4,4,5,5,8,9,10,10,11,11,12,12}},

}

func Test_kata4(t *testing.T) {
	for _, f := range fixtures4 {
		assert.Equal(t, f["expected"], make_kata4(f["input"]), "not equal")
	}

}