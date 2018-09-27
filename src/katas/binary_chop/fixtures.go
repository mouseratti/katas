package binary_chop

var fixtures = []struct {
	num  int
	rang []int
	pos  int
}{
	{1, []int{0, 5, 12, 183}, -1},
	{1, []int{5, 12, 183}, -1},
	{5, []int{1, 2, 3}, -1},
	{5, []int{1, 2, 3, 4, 5}, 4},
	{5, []int{5, 15, 25, 35, 135, 250, 300}, 0},
	{5, []int{1, 2, 4, 5, 13, 24, 245, 333}, 3},
	{99, []int{1, 2, 4, 5, 13, 15, 18, 20, 24, 55, 66, 77, 88, 99, 111, 123, 245, 333}, 13},
}