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
* `-cpu` にて、CPUのコア数を指定することを想定している

### Examples
* 出力の例を書いて、テストをすることができる。
* コメント欄にサンプルコード書くような気持ちでテストが書ける（という理解）
* godocにもサンプルコードが反映されるが、命名規則をきちんとしなければならない
* サンプルは[こちら](examples/)
```
func Example（）{}}
func ExampleF（）{...}
func ExampleT（）{...}
func ExampleT_M（）{...}
```

### Subtests and Sub-benchmarks
* テストを階層にできるらしいが、使いどころがよくわからない

### Main
* Mainのゴルーチンで、テストの実行タイミングを制御できる。
* テストケースの前や後に何かしらの処理を実行させたい場合に便利
* `TestMain()` が存在する場合、 `go test` は `TestMain()` のみの実行となる
* サンプルは[こちら](main/)

```
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
```

### Assert
* Goの `testing` では `assert` が提供されていない
* サードパーティ製のライブラリは存在する
  * https://godoc.org/github.com/stretchr/testify/assert
* サンプルは[こちら](assert/)
