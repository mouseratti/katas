package kata5

func Make_kata2(input ...int) []int {
	last := len(input) - 1
	store := new(int)
	for last > 0 {
		for i := 0; i < last; i++ {
			if input[i] > input[i+1] {
				*store = input[i]
				input[i] = input[i+1]
				input[i+1] = *store
			}
		}
		last -= 1
	}
	return input
}

func is_sorted(input []int) bool {
	length := len(input)
	for pos, val := range input {
		if val > input[pos+1] {
			return false
		}
		if pos == length-2 {
			break
		}

	}
	return true
}
