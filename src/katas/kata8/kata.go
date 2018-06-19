package kata8

import (
	"fmt"
	"reflect"
)

func isSorted(s []int64) bool {
	c := 0
	maxPos := len(s) - 1
	for {
		if c == maxPos {
			return true
		}
		if s[c] > s[c+1] {
			return false
		}
		c++
	}
}

func make_kata(s []int64) []int64 {
	c, length := 0, len(s)
	for {
		if c == length {
			return s
		}
		for i := c; i > 0; i-- {
			if s[i] > s[i-1] {
				break
			}
			s[i], s[i-1] = s[i-1], s[i]
		}
		c++
	}
}

func makeKataNum(s []interface{}) {
	for i := 0; i < len(s); i++ {
		typeI := reflect.TypeOf(s[i])
		switch typeI.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		default:
			panic(fmt.Sprintf("%v is not a number value!", s[i]))
		}
		for j := i; j > 0; j-- {
			var current, previous interface{}
			switch reflect.TypeOf(s[j]).Kind() {
			case reflect.Float32, reflect.Float64:
				current = reflect.ValueOf(s[j]).Float()
			default:
				current = float64(reflect.ValueOf(s[j]).Int())
			}
			switch reflect.TypeOf(s[j-1]).Kind() {
			case reflect.Float32, reflect.Float64:
				previous = reflect.ValueOf(s[j-1]).Float()
			default:
				previous = float64(reflect.ValueOf(s[j-1]).Int())
			}
			if current.(float64) < previous.(float64) {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}

}

func make_kata2(i []int64) {
	var length int = len(i)
	for i1 := length - 1; i1 > 0; i1-- {
		for i2 := i1; i2 < length; i2++ {
			if i[i2] < i[i2-1] {
				i[i2], i[i2-1] = i[i2-1], i[i2]
			}
		}
	}

}
