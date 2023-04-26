# 简单 web server

## index
<http://localhost:8099>

`print: hello, index page`

## hello
<http://localhost:8099/hello?name=007>

`print: hello, 007`

## html
<http://localhost:8099/html>

> 打开html/index.html页面

`这里是Go Project学习站`

## Iris
<http://localhost:8098/ping>

## Iris 路由
- app.Get("/someGet", getting)
- app.Post("/somePost", posting)
- app.Put("/somePut", putting)
- app.Delete("/someDelete", deleting)
- app.Patch("/somePatch", patching)
- app.Head("/someHead", head)
- app.Options("/someOptions", options)

## Iris 页面

> 绑定 views 目录

- 输出消息 messag