# Gest

English documentation located is after Japanese documentation.

## 日本語ドキュメント

### Introduction

Goのテストを[Jest](https://jestjs.io/)のスタイルで書くことができるテストフレームワークです。
Goのテストと100%互換性があるため、2つの種類のテスト混在させて書くことができます(実態はGoのテストのラッパーです)。
そのため、Goテストを残しつつ、徐々にGestに切り替えていくことも可能です。
さらに、100%互換性を保つため、CIの際にGestをCIサーバーにインストールしなくても、`go test`だけでテストを実行することができます。

JavaScript/TypeScriptで[Jest](https://jestjs.io/)を使ったことがある方は、Goに移行した際に、簡単にGestでユニットテストを書くことができます。

以下は、Gestで書いた簡単なサンプルです。

```go
func TestSuiteGest(testingT *testing.T) {

  t := gt.CreateTest(testingT)

  t.Describe("testing Jest style in Go!", func() {

    t.It("should return true if input true", func() {
      r := someFunc(true)
      gt.Expect(v, &r).ToBe(true)
    })

    t.It("should return 2 if 1 given", func() {
      r := addOne(1)
      gt.Expect(v, &r).ToBe(2)
    })

    // 非同期テスト
    t.Async().It("should execute test asynchronously", func() {
      time.Sleep(3 * time.Second)
    })

    // このテストをスキップ
    t.Skip().It("should be skipped", func(){
      // this test will be skipped!
    })

    // 後で書くテストはTodoへ
    t.Todo("some test to write afterward")

  })
}

```

### Gestコマンド

Gestでは、テストを実行する際に、`gest`コマンドを使います。ただし、`go test`を使っても、Gestで書いたテストをそのまま実行することも可能ですが、`gest`コマンドでテストを実行することで、よりわかりやすいテスト結果を得ることができます。

#### 出力サンプル: 

![sample output](/docs/images/sample_output.png)

### インストール方法

```
Go >= 1.21.5
```

#### gestコマンドとライブラリのインストール

```sh
# gestコマンドのインストール
go install github.com/yrichika/gest/cmd/gest@latest
# テスト用のライブラリのインストール
go get github.com/yrichika/gest/pkg/gt@latest
```

環境変数の`GOBIN`が設定されていない場合は、`GOBIN`の設定をしてください。そして、`PATH`に`GOBIN`のパスも含めてください。
`GOBIN`は`~/go/bin`にデフォルトで設定されていることが多いので、わからない方は`~/go/bin`があるか確認してみてください。


#### gestのライブラリインポート

gestをプロジェクト内で使う場合は、Gestの`gt`を`import`して使います


```go
import (
  "testing"

  "github.com/yrichika/gest/pkg/gt"
)

func TestSuiteGest(testingT *testing.T) {
  // ...
}
```


### 使い方

#### テストの作成方法

Gestは、Goの標準のテスト関数の中にテストを作成します。

まずは、Gestのテストを実行するために、`gt.CreateTest()`で、`*Test`を取得してください。

```go
// Goテストで使われる`t`と混同しないために、
// Test関数の引数は、`t *testing.T`とせずに、`testingT *testing.T`と
// 書くことをおすすめします。ただし、`t`と紛らわしくなければ、ここは完全に自由です。
func TestSuiteGest(testingT *testing.T) {
  // CreateTestに *testing.T を渡し、戻り値の *Test である`t`を取得します。
  t := gt.CreateTest(testingT)
}

```

#### テスト実行メソッド

Jestと少し違うところですが、Jestの場合は、`describe`と`it`を自由にネストすることができますが、Gestでは基本的に、`Describe`の中のコールバック関数の中に、1階層の`It`メソッドを書く構造になります。`Describe`の中に`Describe`、`It`の中に`It`もしくは、`It`の中に`Describe`といったような、自由なネストをすることはできません。
ただし、1つのGoテスト関数の中に、複数の`Describe`を書くことができ、`Describe`の中に、複数の`It`テストを書くことができます。
複数の`Describe`を書く際は、その都度`gt.CreateTest`で、`Test`オブジェクトを作成してください。

その際に、必ず`Describe`メソッドが`It`メソッドの外側になるように書いてください。逆に書いた場合はテストが実行されません。

```go
func TestSuiteGest(testingT *testing.T) {

  t1 := gt.CreateTest(testingT)

  // Describe()がテストの「外側」に
  t1.Describe("関数・テスト1の説明", func() {
    // It()はDescribe()の中に書いてください。
    t1.It("テスト1のテスト1", func() {
      // ...
    })

    t1.It("テスト1のテスト2", func() {
      // ...
    })

  })

  // 新しいテスト(t2 *Test)を作成し、
  t2 := gt.CreateTest(testingT)
  // Describe()を複数実行することが可能です。
  t2.Describe("関数・テスト2の説明", func() {
    t2.It("テスト2のテスト内容", func() {
      // ...
    })
  })
}
```

#### `BeforeAll`, `AfterAll`, `BeforeEach`, `AfterEach`

テストの前処理、後処理のために、 `BeforeAll`, `AfterAll`, `BeforeEach`, `AfterEach` が用意されています。

- `BeforeAll`: `Describe`メソッドの前に実行されます。
- `BeforeEach`: 各`It`メソッドの前に実行されます。
- `AfterEach`: 各`It`メソッドの後に実行されます。
- `AfterAll`: `Describe`メソッドの後に実行されます。

以下は、`BeforeEach`の例です。他のメソッドも、実行されるタイミングが違うだけで、書き方は同じです。

```go

func TestSuiteGest(testingT *testing.T) {

  t := gt.CreateTest(testingT)

  t.BeforeEach(func () {
    // テストに必要な前処理を書きます
  })

  t.Describe("", func() {
    // 各Itが実行される前に、BeforeEachの処理が実行されます。
    t.It("", func() {
      // some test
    })

  })
}

```

#### アサーション(Expect)

テスト結果をアサーとする場合は、`Expect`関数と、アサート用のメソッドを使います。`ToBe`や、`ToDeepEqual`などのメソッドを使って、結果の値をアサーとします。
`Expect`は`Expect[A any](test *Test, actual *A)`というシグニチャになっています。そのため、第一引数には、`t`を渡し、第二引数には確認したい値の**ポインタ**を渡してください。

プリミティブ型であれば、`ToBe`だけで、その等価性をアサートできます。

```go

func TestAssertions(testingT *testing.T) {
  t := gt.CreateTest(testingT)

  t.Describe("assertion sample", func() {
    t.It("should return true", func() {
      r1 := someFunc()
      expected1 = true
      // プリミティブ型 int, bool, string, ...etc はすべて`ToBe`だけでアサート可能です
      gt.Expect(t, &r1).ToBe(expected1)

      r2 := otherFunc()
      expected2 = 13
      gt.Expect(t, &r2).ToBe(expected2)
    })
  })
}
```


#### `gest`コマンドでテスト実行

`go.mod`のあるプロジェクトルートで実行することをおすすめします。ただし、必須ではありません。
`gest`コマンドは、コマンドの実行ディレクトリから再帰的に`_test.go`で終わるファイル名を探し、それらすべてのテストを実行します。
プロジェクトルート以外で実行した場合は、そのディレクトリ以下のファイルを再帰的に検索し、`_test.go`ファイルのテストをすべて実行します。
1つだけのテストを実行したい場合は、以下の`-run`オプションを使用してください。

```sh
gest
```

##### `-run`

`-run`の後に、テスト関数の名前を指定してください。`-run`の後の引数に渡した文字列は、Regexでマッチするテストをすべて実行します。

```sh
gest -run TestFunctionNameToRun
```

##### `-v`

`gest`の出力に加え、`go test`で実行した場合の出力も表示します。
`gest`では、基本的には余分なコンソールへの出力を抑えるため、特定の文字列を出力しない仕組みになっています。
そのため、場合によっては、テスト内でデバッグ用に`println`などを使った場合、出力がコンソールにされないことがあります。
そういった際に、`-v`オプションを付けると、デバッグ用に出力したい文字列も全て出力されます。

```sh
gest -v
```

##### `-vv`

`gest`の出力に加え、`go test -v`で実行した場合の出力も表示します。


```sh
gest -vv
```


##### `-cover`

HTMLのカバレッジファイルを`gest_coverage`ディレクトリに作成します。
カバレッジファイルを出力するディレクトリ名を指定したい場合は、`-coverprofile`オプションを使用します。

カバレッジ出力の場合のみ、`go.mod`のあるプロジェクトルートで実行することが**必須**になります。`go.mod`がないディレクトリで、`gest -cover`を実行するとエラーになります。

```sh
gest -cover
# OR 出力先のディレクトリ名を指定する場合
gest -cover -coverprofile=[any_dir_name]
```

##### `-all-dirs`

デフォルトで、Gestは`examples`フォルダのテストを実行しません。
もし`examples`フォルダの中のテストを実行したい場合、もしくは`examples`という名前をつけたサブディレクトリの中にテストがある場合は、`-all-dirs`を指定することで、全てのディレクトリ内のファイルがテストの対象になります。

```sh
gest -all-dirs
# -run と組み合わせることもできます
gest -all-dirs -run TestInExamplesDir
```


### Assertions


#### `ToBe(T)`

これで、`int`系、`bool`, `string`, `complex64`, etc、プリミティブ型と`time.Time`、カスタムの`struct`型はすべてアサートすることができます。

構造体が同じかどうかの確認は、内部の処理では、`reflect.DeepEqual()`を使って、2つの値の比較を行っています。

```go
gt.Expect(t, &r).ToBe(true)

// 構造体の場合
expected := User{Name: "john", Age: 10}
gt.Expect(t, &r).ToBe(expected)
```


#### `ToBeSamePointerAs(*T)`

ポインタが同じかテストすることができます。


```go
a := 1
p := &a
gt.Expect(t, &a).ToBeSamePointerAs(p)
```


#### `ToBeNil()`

値が`nil`かどうかをアサートします。

```go
var nilValue *int = nil
gt.Expect(t, nilValue).ToBeNil()
```

#### `ToMatchRegex(string)`

文字列が正規表現と一致するかアサートします。

```go
str := "foo"
gt.Expect(t, &str).ToMatchRegex("^foo$")
```

#### `ToContainString(string)`

Actualに指定した文字列が含まれるかアサートします。

```go
str := "hello, world"
gt.Expect(t, &str).ToContainString("world")
```

#### `ToBe_(func(T, T) bool, T)`

`ToBe`に`_`が付いた`ToBe_`は第一引数に、比較する関数を取り、actualの値とexpectedの値を比較します。
「～以上」や、「～以下」のように、2つの値を比較したい場合に使います。
`gt.GreaterThan`などの比較用関数は用意されているので、比較用関数は自作しなくても大丈夫です。
ただし、自作して、引数に渡すことも可能です。


```go
val := 10
gt.Expect(t, &val).ToBe_(gt.GreaterThan, 1)
gt.Expect(t, &val).ToBe_(gt.GreaterThanOrEq, 10)
gt.Expect(t, &val).ToBe_(gt.LessThan, 11)
gt.Expect(t, &val).ToBe_(gt.LessThanOrEq, 10)
// Betweenに指定した数値(min)と、expectedに指定した値(max)の間にある場合はpassします
gt.Expect(t, &val).ToBe_(gt.Between(1), 10)

// 時間を比較することもできます
time := time.Now()
past := time.Now().Add(-1 * time.Minute)
future := time.Now().Add(1 * time.Minute)
gt.Expect(t, &time).ToBe_(gt.After, past)
gt.Expect(t, &time).ToBe_(gt.AfterOrEq, past)
gt.Expect(t, &time).ToBe_(gt.Before, future)
gt.Expect(t, &time).ToBe_(gt.BeforeOrEq, future)
// TimeBetweenに指定した時間(from)と、expectedに指定した値(to)の間にある場合はpass
gt.Expect(t, &time).ToBe_(gt.TimeBetween(past), future)


```


#### `Not()`

`Not()`をアサーションの前につけることで、逆の意味のアサーションを行います。

```go
gt.Expect(t, &r).Not().ToBe(true)
gt.Expect(t, &x).Not().ToBeNil()
// ...etc
```


#### `LogWhenFail(string).Expect(*A)`

アサートエラーが起きた場合は、自動的に適切なメッセージを出力しますが、独自のメッセージに変更したい場合は、`LogWhenFail`関数を使います。
`LogWhenFail`を使う場合は、`t *Test`は、`Expect`ではなく、`LogWhenFail`の第一引数に指定してください。

```go
a := 1
b := 2
gt.LogWhenFail[int](t, "Fail時に出力されるカスタムメッセージ").Expect(&a).ToBe(b)
```


#### `ExpectPanic(*Test).ToHappen(func())`

`panic`が起きたかどうかをアサートする場合は、`ExpectPanic`関数を使います。
`ExpectPanic`でも、`Not()`を使い、`panic`が起きなかったことをアサートすることができます。

```go
// panicが起きたことをアサート
gt.ExpectPanic(t).ToHappen(func() {
  panic("panic")
})

// panicが起きないことをアサート
gt.ExpectPanic(t).Not().ToHappen(func () {
  someFuncMightPanic()
})
```


#### TODO:

その他のアサーションは現在実装中です。
基本的には、Jestと同じアサーションと、Go独自に必要そうなアサーションを追加していきます。


## English Documentation

WIP: I'm working on it!


