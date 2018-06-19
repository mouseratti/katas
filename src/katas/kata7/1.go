package kata7

import "fmt"

func ReadList(input []interface{}, c *chan interface{}, pos int) {
	period := 3
	for i := pos; i < len(input); i += period {
		fmt.Printf("Readlist%v: reading %v\n", pos, input[i])
		*c <- input[i]
		input[i] = nil
		fmt.Printf("Radlist%v: input is %v\n", pos, input)
	}
}

func locker(list []interface{}, ch *chan interface{}) {
	var lock bool
	for {
		lock = true
		for _, val := range list {
			if val != nil {
				lock = false
			}
		}
		if lock {
			close(*ch)
			return
		}
	}
}

func Make_kata_1(inp []interface{}, channel *chan interface{}) {
	fmt.Println(inp)
	for pos := 0; pos < 3; pos++ {
		go ReadList(inp, channel, pos)
	}
	go locker(inp, channel)
	for {
		val, ok := <-*channel
		if !ok {
			fmt.Printf("%v is closed! stopping...", *channel)
			break
		}
		fmt.Println(val)
	}
	fmt.Println(inp)
}
