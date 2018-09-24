package binary_chop

func make_kata2(num int, rang []int) int {
	return find_recursively(num, rang, 0, len(rang)-1)
}

func find_recursively(num int, rang []int, start, finish int) int {
	if num < rang[start] || num > rang[finish] {
		return -1
	}

	if num == rang[start] {
		return start
	}
	if num == rang[finish] {
		return finish
	}
	var avgIndex int = (start + finish) / 2
	if num == rang[avgIndex] {
		return avgIndex
	}
	if num < rang[avgIndex] {
		return find_recursively(num, rang, start, avgIndex-1)
	}
	if num > rang[avgIndex] {
		return find_recursively(num, rang, avgIndex+1, finish)
	}
	return -2
}
