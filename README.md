# embedctl
演示Go命令行嵌入前端资源

## 需要Go版本 >= 1.16

## 安装
```go
go install github.com/gotomicro/embedctl@latest
```

## 运行
```
➜  ~ embedctl
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ant/*filepath            --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /ant/*filepath            --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] GET    /welcome                  --> main.main.func2 (3 handlers)
[GIN-debug] GET    /api/hello                --> main.main.func3 (3 handlers)
[GIN-debug] GET    /webui/*filepath          --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /webui/*filepath          --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
```

## 访问
浏览器输入localhost:8888


