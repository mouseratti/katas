package kata5

type fixture struct {
	input    []int64
	expected []int64
}

type push_buble_fixture struct {
	input_slice []int64
	position int
	expected_slice []int64
}


var fixtures []fixture = []fixture{
	{
		[]int64{1008, 5, 3, 2, 6, 11, 291, 144, 516},
		[]int64{2, 3, 5, 6, 11, 144, 291, 516, 1008},
	},
	{
		[]int64{10,9,8, 7, 6,5,4,3,2,1,0},
		[]int64{0,1,2,3,4,5,6,7,8,9,10},
	},


}

var push_buble_fixtures = []push_buble_fixture{
	{
		[]int64{1,3,2},
		1,
		[]int64{1,2,3},
	},
	{
		[]int64{1,3,2,6,8},
		0,
		[]int64{3,1,2,6,8},
	},

}