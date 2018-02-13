package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

- mapの操作
	- 要素の挿入や更新
		- ```m[key] = elem```
	- 要素の取得
		- ```elem = m[key]```
	- 要素の削除
		- ```delete(m, key)```
	- キーに対する要素の存在確認
		- ```elem, ok = m[key]```
		- ```elem, ok := m[key]```
		- mにkeyがあればokがtrueとなり、elemに要素が入る
		- mにkeyがなければokがfalseとなり、elemはmapの要素の型のゼロ値になる
