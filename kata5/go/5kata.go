package kata5


func make_kata5(inp []int) []int {
	var storage int
	lastElement := len(inp) - 1
	for lastElement > 0 {
		for i := 0; i < lastElement; i++ {
			if inp[i] > inp[i+1] {
				storage = inp[i]
				inp[i], inp[i+1] = inp[i+1], storage
			}
		}
		lastElement -= 1
	}
	return inp
}
