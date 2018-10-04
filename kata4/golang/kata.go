package kata

import "fmt"

func MakeKata() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}

}

func fizzbuzz(x int) string {
	switch {
	default:
		return fmt.Sprintf("%v", x)
	case x%3 == 0 && x%5 == 0:
		return "fizzbuzz"
	case x%3 == 0:
		return "fizz"
	case x%5 == 0:
		return "buzz"
	}
}
