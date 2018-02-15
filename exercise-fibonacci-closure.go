package main

import "fmt"

func fibonacci() func() int {
	f1, f2 := 0, 1
	return func() int {
		f := f1
		f1, f2 = f2, f1+f2
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
