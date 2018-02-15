package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		fmt.Println("DEBUG:", v)
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}