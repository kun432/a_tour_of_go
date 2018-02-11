package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // スライスの長さを０にする
	printSlice(s)

	s = s[:4] // スライスの長さを伸ばす
	printSlice(s)

	s = s[2:] // スライスから値を減らして容量も減らす
	printSlice(s)

	//s = s[:6] // 容量を超えては伸ばせない
	//printSlice(s) // slice out of rangeになる
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
