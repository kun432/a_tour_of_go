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
