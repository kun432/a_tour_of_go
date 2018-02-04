# a_tour_of_go

A Tour of Goを試してみる

## パッケージ(packages.go)

- プログラムはパッケージで構成される
- mainパッケージでプログラムは起動する
- importで別のパッケージを読み込む
- パッケージ名はインポートパスの最後で読み込む
  - ```math/rand```ならば```rand.xxx```)

## インポート(imports.go)

```
import "fmt"
import "math"
```

or 

```
import (
  "fmt"
  "math"
)
```

## エクスポートされる名前(exports.go)

- 大文字で始めればエクスポートされる
- 小文字はエクスポートされない

## 関数(functions.go/functions-continued.go)

- 引数を取る場合は型を指定
  - 型は変数名の後になる

```
func hoge(x int, y int) int {
  ・・・
}
```

  - 引数が複数で肩が同じ場合は型を一つだけ書けば良い

```
func hoge(x, y int) int {
  ・・・
}
```

## 戻り値(multiple-results.go)

- 戻り値を複数取れる
- その分だけ型を指定する

```
func swap(x, y string) (string, string) {
	return y, x
}
```

## 戻り値に名前をつける(named-results.go)

- return で指定する必要はない(naked return)
- 関数ドキュメント等でわかりやすくなる

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

## 変数の初期化(variables-with-initializers.go)

- 変数の初期化はvarを使う
  - var 変数名 型
- 初期値を入れる場合は=で初期値を指定

```
var i, j int = 1, 2
```

## 変数宣言の省略化(short-variables-declairations.go)

- varを使わず、```:=```でも宣言できる
- ただし関数内のみ
  - 関数の外ではvarを使うこと

```
	k := 3
	c, python, java := true, false, "no!"
```

## 基本的な型(basic-types.go)

- 真偽値
  - bool
  - true / false
- 文字列
  - string
  - "xxxx"
- 数値
  - int / int8 / int16 / int32 / int64
  - uint / uint8 / uint16 / uint32 / uint64 / uintptr
  - float32 / float64
  - byte
  - rune
  - complex64 / complex128

## ゼロ(zero.go)

- 初期値が与えられない変数は型によってゼロな値がそれぞれ入る
  - 数値なら0
  - 文字列なら""
  - 真偽値ならfalse
