package kata5

func make_kata6(input []int) []int {
    max := len(input) - 1
    var storage int
    for max > 0 {
        for i := 0; i < max; i++ {
            if input[i] > input[i+1] {
                storage = input[i]
                input[i], input[i+1] = input[i+1], storage
            }
        }
        max -= 1
    }
    return input
}
