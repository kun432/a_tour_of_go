package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 戻りのスライスを作る
	image := make([][]uint8, dy)
	// スライスを二次元にする
	for y := range image {
		image[y] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			image[y][x] = uint8((x + y) / 2)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
