package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fi_2,fi_1,f,i := 0,1,0,0
	return func() int {
		if i == 0 {
			i++
			return fi_2
		} else if i == 1 {
			i++
			return fi_1
		} else {
			f = fi_1+fi_2
			fi_2,fi_1 = fi_1,f
			i++
			return f
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
