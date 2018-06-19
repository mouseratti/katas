package main

import "katas/kata7"

func main() {
	inp := []interface{}{1, "two", 3.0, 11, "four", 6.0, 111, "eight", 9.0}
	ch := make(chan interface{})
	kata7.Make_kata_1(inp, &ch)
}
