package kata5

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func Test_make_kata6(t *testing.T) {
    fixtures := [][]int{
        {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 7, 6, 1},
        {10, 9, 8, 7, 6, 5, 344, 3, 888992, 1, 0, 7, 6, 10000000000000},
        {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 7, 6, 1},
        {1110, 9, 8, 1117, 22226, 5, 4, 3, 2, 1, 0, 7, 6, 1},
        {10, 9, 8, 7, 6, 57777777, 4, 3, 2, 1, 0, 7, 6, 1},
        {10, 9, -8, -7, 6, -57777777, 4, -3, 2, 1, 0, 7, 6, 1},
        {10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355},
        {10, 9, 8, 7, 6, 5, 4, 355, 2848, 1, 0, -1, -2, -100, -355, -400},
    }
    for _, val := range fixtures {
        make_kata6(val)
        assert.True(t, is_sorted(val), fmt.Sprintf("%v is not sorted properly!", val))
    }
}
