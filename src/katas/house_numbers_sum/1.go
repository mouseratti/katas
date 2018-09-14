package house_numbers_sum

func make_kata1(input []int64) (out int64) {
	for _, val := range input {
		if val == 0 {
			return
		}
		out += val
	}
	return
}

func make_kata2(input []int64) int64 {
	var out int64
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			return out
		}
		out += input[i]

	}
	return out
}

func make_kata3(input ...int64) (result int64) {
	for pos, _ := range input {
		if input[pos] == 0 {
			return
		}
		result += input[pos]
	}
	return
}
