package kata5

func make_kata4(input []int) []int {
	last := len(input) - 1
	for {
		needToSort := false
		for i := last; i > 0; i-- {
			needToSort = try_push4(i, input) || needToSort
		}
		if !needToSort {
			break
		}
	}
	return input
}

func try_push4(pos int, input []int) bool {
	if pos == 0 {
		return false
	}
	if input[pos] < input[pos-1] {
		storage := input[pos]
		input[pos] = input[pos-1]
		input[pos-1] = storage
		return true
	}
	return false
}
