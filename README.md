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

## 配列

- [n]T で T型のn個の配列になる
  - 配列の長さも方の一部分なのでサイズを変えることはできない

```
var a[10]int
```

- 配列へのアクセスは一般的。```a[1]```とか。
- まとめてアクセスもできる

```
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## スライス

- 配列は固定長
- スライスは可変長
- 柔軟な配列、というか、スライスのほうが一般的に使われる

```
var a []int
```

- スライスへのアクセス

```
a[1:2]
```

- スライスは配列への参照のようなもの
  - スライス自体はデータを格納しているわけではない
  - 配列の部分列を示しているだけ
  - スライスを変更するともとの配列の要素が変更される

- スライスリテラル
  - 長さのない配列リテラルのようなもの
  - 配列リテラルの場合だとこう

```
[3]bool{true,false,true}
```

- スライスリテラル

```
[]bool{true,false,true}
```

- スライスのデフォルト値
  - スライスの下限・上限を省略すると初期値が設定される

```
var a [10]int

//いかはすべておなじ
a[0:10]
a[:10]
a[0:]
a[:]
```

- スライスには長さと容量がある
  - 長さは含まれる要素数
    - len(s)
  - 容量はもととなる配列の要素数
    - cap(s)
  - 容量があるスライスを再スライスするとスライスの長さを伸ばせる
  - 容量を超えてスライスの長さを伸ばそうとするとエラーになる

- スライスのゼロ値はnil
  - 0の長さと容量

- makeを使ってスライスを作成することもできる
  - 動的サイズの配列を作成する
    - ゼロ化された配列を割り当ててそれを指すスライス
      - ```a := make([]int, 5)```
        - lenもcapも5、中身は全部ゼロ
    - makeでスライス容量を指定できる
      - ```a := make([]int, 0, 5)```
        - lenは0、capは5、よって中身はない

- スライスのスライス
  - スライスに、スライスを含む、他の型を含めることもできる

- スライスへの追加はappendを使う
  - ```s = append(s, 0)```
    - 容量が小さければ拡張される
  - ```s = append(s, 1, 2, 3, 4)```
    - 複数まとめてもできる

# レンジ

- スライスやマップでイテレーターのために使う
- スライスをrangeで繰り返すと、indexと要素のコピーをそれぞれ返す

```
var pow = []int{1,2,4,8,16,32,64,128}

func main() {
  for i, v := range pow {
    // iがindexで、vがスライスの各要素
  }
}
```

- _を使うことでインデックスや値を捨てることができる
- インデックスだけが必要ならば値は省略できる

```
for i := range pow {
  // indexはiでとって、あとは捨てる
}
for _, value := range pow {
  // indexはすてて、値だけvalueでとる
}
```

## エクササイズ: スライス

```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
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
```

## マップ

- キーと値を関連付けする
- ゼロ値はnil
  - キーを持っていない
  - キーを追加できない
- makeで初期化できる

```
var m map[string]Vertex
m = make(map[string]Vertex)
m["hogehoge"] = Vertex{
  40.88888, -74.39967,
}
```

- mapリテラル
  - キーが必ず必要

```
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

- mapに渡すトップレベルの型が単純な型名の場合は要素から推測できる場合省略できる

```
var m = map[string]Vertex{
	"Bell Labs": { 40.68433, -74.39967 },
	"Google": { 37.42202, -122.08408 },
}
```

## エクササイズ: マップ

```
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
```

## 関数値

- 関数を変数として使う
- 関数に関数を渡す、とか、戻り値として関数を返すとか

```
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

## クロージャ

- 外部から変数を参照する関数


## エクササイズ：フィボナッチクロージャ

```
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
```

一旦ここで終了、イマイチ理解できていないところが結構ある

- 配列、スライス、マップがはっきり理解できてない、PerlとかJSとかだと結構自由にかけちゃうイメージなので。
- クロージャ自体は理解しているが、ちょっと書き方がピンときてない

このあたりを理解しないと、ちょっとこのまま進めても意味が無い気がするので別の書籍等を参照してみることとしたい。
