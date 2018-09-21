package binary_chop

func make_kata1(r []int, n int) int {
	startPosition := 0
	endPosition := len(r) - 1
	var currentPosition int

	for {
		currentPosition = (startPosition + endPosition) / 2
		if n == r[currentPosition] {
			return currentPosition
		}
		if startPosition == endPosition {
			break
		}
		switch {
		case n > r[currentPosition]:
			startPosition = currentPosition + 1
		case n < r[currentPosition]:
			endPosition = currentPosition - 1
		}

	}
	return -1
}
