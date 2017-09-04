package kata5

func make_kata3(input []int64) []int64 {
	//var s int
	last := len(input) - 1
	for {

		for i := 0; i < last; i++ {
			try_push3(i, last, input)
		}
		last -= 1
		if last < 1 {
			break
		}
	}
	return input
}

func try_push3(position int, lastElement int, input[]int64) {
	if position == lastElement {
		return
	}
	if input[position] > input[position + 1] {
		input[position + 1], input[position] = input[position], input[position + 1]
	}
}