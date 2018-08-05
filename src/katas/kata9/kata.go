package kata9

func make_kata(inp ...int64) []int64 {
	result := make([]int64, len(inp))
	for pos, val := range inp {
		result[pos] = val * 2
	}
	return result
}

func make_kata2(i ...int64) []int64 {
	c := 0
	last := len(i) - 1
	for {
		i[c] *= 2
		if c == last {
			break
		}
		c++
	}
	return i
}
