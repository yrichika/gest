# Gest

日本語の説明は、英語の説明の後にあります



You can write in Jest manner with Go test.
It also has 100% compatibility with go standard test. It just extends writing style from Go test. You don't lose anything from go standard test, and no problem having both gest and go test in one project.
Under the hood, it just uses go test.

If you are a javascript/typescript programmer and already used Jest, it's really easy to learn.

```
Go >= 1.21.5
```

gest command installation

```sh
go install github.com/yrichika/gest/cmd/gest@latest
```

make sure `GOBIN` is set and it's included in your `PATH`.
`GOBIN` is usually `~/go/bin`. add it to the `PATH` environment variable

gest framework installation
```sh
go get # TODO:
```

