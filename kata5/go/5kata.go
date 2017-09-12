package kata5

func is_sorted(input []int) bool {
	length := len(input)
	for pos, val := range input {
		if val > input[pos + 1] {
			return false
		}
		if pos == length - 2 {
			break
		}

	}
	return true
}

func make_kata5(input []int) {
	maxPosition := len(input) - 1
	var s int
	for maxPosition > 1 {
		for i := 0; i < maxPosition; i++ {
			if input[i] > input[i + 1] {
				s = input[i]
				input[i], input[i + 1] = input[i + 1], s
			}
		}
		maxPosition -= 1

	}
}