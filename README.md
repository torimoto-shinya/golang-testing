# Golang Testing

## URL

https://golang.org/pkg/testing/

## 概要

* Goでは `testing` というデフォルトで入っているライブラリにて、パッケージの自動テストをサポートしている
* テスト用の関数は `func TestXxx(*testing.T)` のように記述する。 `Xxx` の部分は小文字スタートにしてはいけない
* テストファイルの名前は `_test.go` で終わる必要がある

### Benchmarks
* `func BenchmarkXxx(*testing.B)` と記述すると、 `go test` コマンド時に `-bench` フラグを付けることで、ベンチマークを測定できる
* サンプルコードは [こちら](hello/)
* 実行の方法は、以下
```
# ベンチマークのテストコードがたくさん含まれている場合、すべて実行
$ go test -bench .  
BenchmarkHello-4        20000000                83.6 ns/op

# Hello とつくものだけ実行
$ go test -bench Hello

# 同じテストを10回実行
$ go test -bench Hello -count 10
```

* `-count` を指定すると、同じテストを指定回数実行できる
  * 例： `go test -bench hello -count 10`
* `-benchtime` のデフォルトは1秒。変化させることで、 `b.N` の値が変化する。
  * 例： ` go test -bench Hello -benchtime 0.1s`
* 他のオプションは https://golang.org/cmd/go/#hdr-Testing_flags に記載

#### 並列処理でのベンチマーク
* サンプルコードは [こちら](bench-parallel/)
* `-cpu` にて、CPUのコア数を指定することを前提として考えられている

### Examples
* 出力の例を書いて、テストをすることができる。
* コメント欄にサンプルコード書くような気持ちでテストが書ける（という理解）
