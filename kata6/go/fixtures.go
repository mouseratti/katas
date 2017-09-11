package kata6


type changePostitionFixture struct {
	input     []int
	expected  []int
	positions [2]int
}

var changePositionFixtureList []changePostitionFixture = []changePostitionFixture{
	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{2, 6, 5, 4, 3, 1},
		[2]int{4, 0},
	},
	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{6, 2, 5, 4, 3, 1},
		[2]int{4, 1},
	},

	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{5, 4, 3, 6, 2, 1},
		[2]int{0, 3},
	},
	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{6, 5, 4, 3, 2, 1},
		[2]int{0, 0},
	},
	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{6, 5, 4, 3, 2, 1},
		[2]int{5, 5},
	},
	changePostitionFixture{
		[]int{6, 5, 4, 3, 2, 1},
		[]int{5, 6, 4, 3, 2, 1},
		[2]int{0, 1},
	},

}

var kataFixtureList = [][]int{
	[]int{10, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 6},
	[]int{8, 9, 9999, 10, 9, 8, 7, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 6, 126, 127, 55555},
	[]int{10, 9, 8, 7, 111, 113, 1234, 3212, 8, 6, 5, 4, 3, 4, 5, 6, 7, 8, 9, 6},
	[]int{10, 9, 8, 73, 3, 123, 6, 5, 18, 4, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9, 6},
}