## 从网络获取数据

### 运行方式
main.go调用了其他同包(main)文件的函数
```
cd project01/main
go run .
```
或者
```
go run main.go test01.go test02.go
```

### 加载依赖
`go mod download -v`

### 清理
`go mod tidy -v`


**timeout 异常**

- 强制开启GO111MODULE

    > 确保你的go版本大于1.11go version(最好大于1.13)

    `go env -w GO111MODULE=on`
- 使用中国代理

    `go env -w GOPROXY=https://goproxy.cn,direct`

### [Test01](./main/test01.go)

- 使用 Http GET, 并设置请求头, 否则会被服务器拒绝请求
- 将获得的数据转换为json(非结构化), 从json中获取指定key数值

### [Test02](./main/test02.go)