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

## エクスポートされる名前

- 大文字で始めればエクスポートされる
- 小文字はエクスポートされない
