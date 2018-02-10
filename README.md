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

## 型キャスト

- 型名がそのままキャストする関数になる
- 明示的な型変換が必要になる
  - しないとエラーになったりする

## 型推論

- 型を指定せずに変数宣言すると、右辺の値から推測される
  - 数値だと精度によって変わる

## 定数

- constを使って宣言
- := は使えない
- 文字、文字列、真偽値、数値のみ
  - 型の指定は不要

## 数値の定数

- 型の指定は不要
  - 数値と言いつつ、高精度な値
  - 必要な型を推論する
- intでは足りない場合はconstを使う？？？？
  - 型指定がないため、無限制度になる（らしい

参考) https://qiita.com/hkurokawa/items/a4d402d3182dff387674

## for

- Cライクだが()不要

```
for i := 0; i < 10; i++ {
	・・・
}
```

- 初期化と後処理は除外できる

```
for ; i < 10; {
	・・・
}
```

- さらにセミコロンを除く
  - whileっぽくなる
  - すなわちgoではforしか使わない

```
for i < 10 {
	・・・
}
```

- 条件も省くと無限ループになる

```
for {
	・・・
}
```

## if

- ()は不要

```
if x < 0 {
  ・・・
}
```

- ifのスコープ内だけで有効な変数を条件の前に指定できる
  - スコープの外では使えない

```
if v := 0; x < 0 {
  ・・・
}
```

- elseも使える
  - ifで宣言された変数はelseでも使える

```
if v := 0; x < 0 {
  ・・・
  return v
} else {
  ・・・
  return v
}
```

## エクササイズ：ループと関数

ニュートン法というのがあるらしい、がさっぱりわからない。
いろいろ調べてみた

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("debug    : ", z)
	}
	return z
}

func main() {
	fmt.Println("sqrt     : ", sqrt(2))
	fmt.Println("math.Sqrt: ", math.Sqrt(2))
}
```

実行結果

```
go run excecise-loops-and-functions.go
debug    :  1.5
debug    :  1.4166666666666667
debug    :  1.4142156862745099
debug    :  1.4142135623746899
debug    :  1.4142135623730951
debug    :  1.414213562373095
debug    :  1.4142135623730951
debug    :  1.414213562373095
debug    :  1.4142135623730951
debug    :  1.414213562373095
sqrt     :  1.414213562373095
math.Sqrt:  1.4142135623730951
```

要は、公式に従って繰り返し計算していけば平方根に近づく、というものらしい
イケてないのは、10回という上限を設定してしまっているところなのでここを
「次に値が変化しなくなった (もしくはごくわずかな変化しかしなくなった) 場合にループを停止させます。」という条件をつける

いろいろ調べてみたが

- 前回のzと今回のzを比べて、ごくごく小さい変化になった場合にループを抜ける
- ごくごく小さい変化の基準を1.0e-6としている
- 小さい変化の絶対値と上記を比較

ということの様子

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for diff := 1.0; math.Abs(diff) > 1e-6; {
		diff = (z*z - x) / (2 * z)
		z -= diff
		fmt.Println("debug    : ", z)
	}
	return z
}

func main() {
	fmt.Println("sqrt     : ", sqrt(2))
	fmt.Println("math.Sqrt: ", math.Sqrt(2))
}
```

実行結果

```
debug    :  1.5
debug    :  1.4166666666666667
debug    :  1.4142156862745099
debug    :  1.4142135623746899
debug    :  1.4142135623730951
sqrt     :  1.4142135623730951
math.Sqrt:  1.4142135623730951
```

## switch

- breakを書く必要はない
  - 上から下に評価される
  - 選択されたcaseだけが実行される
  - 自動的にbreakされる

```
switch x {
  case "a":
    ・・・
  case "b":
    ・・・
  default:
    ・・・
}
```

- switchステートメントに変数を宣言することもできる
  - もちろんスコープはその中になる

```
switch x := 0; x {
  ・・・
}
```

- 条件無しで書くこともできる
  - switch trueと同じこと
  - if else をシンプルに表現できる

```
switch {
  case 0:
    ・・・
  case 1:
    ・・・
  default:
    ・・・
}
```
## defer

- 遅延実行
  - deferに関数を渡す
  - 呼び出し元の処理の終わり（return）まで遅延
  - 評価自体はdeferで渡された時点でお行われる

```
func main() {
  defer fmt.Println("defer")

  fmt.Println("End")
}
```

実行結果

```
End
defer
```

- deferに複数の関数を渡した場合はスタックされる
  - last-in-first-outで実行される

```
func main() {
	fmt.Println("count start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
```

実行結果

```
count start
done
9
8
7
6
5
4
3
2
1
0
```

## ポインタ

- int型の変数のポインタ変数を宣言
  - ゼロ値はnilになる

```
var p *int
```

- &で変数のアドレスを取得する

```
i := 42
p = &i
```

- *で変数のアドレスから中身にアクセスする
  - デリファレンスのこと

```
fmt.Println(*p)
*p = 21
```

## 構造体

```
type Vertex struct {
  X int
  Y int
}

fmt.Println(Vertex{1,2})
```

- .を使ってフィールドにアクセスする

```
v := Vertex{1,2}
fmt.Println(v.X) // 1
fmt.Println(v.Y) // 2
```

- ポインタを使って構造体へアクセスする
  - structのフィールドへのアクセスもポインタから行える
  - ```(*p).X```でもよいし```p.X```でもよい

```
type Vertex struct {
  X int
  Y int
}

function main() {
  v := Vertex{1,2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
```

実行結果

```
{1000000000 2}
```

- structリテラル
  - フィールドの値を```Name: ```で列挙して初期値を与える

```
type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1,2}
  v2 = Vertex{X:1} // Yは0で初期化される
  v3 = Vertex{} // XもYも0
  p = &Vertex{1,2} // *Vertex型のポインタ
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
```

実行結果

```
{1 2} &{1 2} {1 0} {0 0}
```