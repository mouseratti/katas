package kata5


type fixture struct {
	input    []int64
	expected []int64
}

func (f *fixture) reverseInput() {
	var s int64
	last := len(f.input) - 1
	for i:=0;i < last / 2; i++ {
		s = f.input[i]
		f.input[i] = f.input[last-i]
		f.input[last-i] = s
	}
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
	{
		[]int64{8,1,8,2,8,15,19,18},
		[]int64{1,2,8,8,8, 15,18,19},
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
