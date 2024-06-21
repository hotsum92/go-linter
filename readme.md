# go-linter

[Go で作る自作 Linter 開発入門 - emahiro/b.log](https://ema-hiro.hatenablog.com/entry/2021/08/02/035253)

[GoのAST全部見る - \*\*\*の日記](https://monpoke1.hatenablog.com/entry/2018/12/16/110943)


##### get ast

```sh
go install golang.org/x/tools/cmd/gotype
gotype -ast ./main.go
```

## unitchecker

[Goのanalysisパッケージを使った静的解析を実装する](https://zenn.dev/micin/articles/2023-12-09_ykobayashi_golang_static_analysis)

```
go build -o analysis.o ./main.go
go vet -vettool=analysis.o ./sample/main.go
```

gotype -ast ./sample/main.go
