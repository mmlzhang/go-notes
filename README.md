
## learning golang

goland 常用工具  (goland)
goimport 该工具在 gofmt 的基础上增加了自动删除和引入包
`go get golang.org/x/tools/cmd/goimports`

go vet vet工具可以帮我们静态分析我们的源码存在的各种问题，例如多余的代码，提前return的逻辑，struct的tag是否符合标准等
`go get golang.org/x/tools/cmd/vet`
使用　　`go vet .`