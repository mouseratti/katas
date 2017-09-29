package kata6

func changePosition(inp []int, position int, newPosition int) {
	switch newPosition - position {
	case 1, -1:
		s := inp[position]
		inp[position], inp[newPosition] = inp[newPosition], s
		return
	default:
		break
	}

	for inp[position] != inp[newPosition] {
		nextPosition := getNextPos(position, newPosition)
		changePosition(inp, position, nextPosition)
		position = nextPosition
	}
	return

}

func getNextPos(pos int, newPos int) (result int) {
	result = pos
	if pos > newPos {
		result = pos - 1
	} else {
		result = pos + 1
	}
	return result
}

func is_sorted(input []int) bool {
	length := len(input)
	for pos, val := range input {
		if pos == length - 1 {
			break
		}
		if val > input[pos + 1] {
			return false
		}
	}
	return true
}

func make_kata(input []int) {
	for pos, val := range input {
		for sortedPos, sortedVal := range input[:pos] {
			if val < sortedVal {
				changePosition(input, pos, sortedPos)
				break
			}
		}
	}
}
