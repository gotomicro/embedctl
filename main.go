package main

import (
	"embed"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/embedctl/ant"
	"github.com/gotomicro/embedctl/simple"
	"io/fs"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// 设置Ant Design静态资源
	webuiAntObj := &webui{
		webuiEmbed: ant.WebUI,
		path:       "dist",
	}
	// 设置Ant Design前端访问，try file到index.html
	webuiAntIndexObj := &webuiIndex{
		webui: webuiAntObj,
	}

	// 设置简单的演示静态资源
	webuiSimpleObj := &webui{
		webuiEmbed: simple.WebUI,
		path:       "webui",
	}

	handler := gin.Default()
	// 设置路由
	{
		// 设置ant design的路径，在config.ts里配置
		handler.StaticFS("/ant/", http.FS(webuiAntObj))
		// 访问首页跳转到ant design的welcome页面
		handler.GET("/", func(ctx *gin.Context) {
			ctx.Redirect(302, "/welcome")
			return
		})
		// Ant Design前端访问，try file到index.html
		handler.GET("/welcome", func(context *gin.Context) {
			context.FileFromFS("/welcome", http.FS(webuiAntIndexObj))
		})

		// 设置hello world
		handler.GET("/api/hello", func(ctx *gin.Context) {
			ctx.JSON(200, "Hello EGO")
			return
		})

		// 设置简单的演示静态资源
		handler.StaticFS("/webui/", http.FS(webuiSimpleObj))
	}

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(listener)
	}
	httpServer := &http.Server{
		Handler: handler,
	}
	err = httpServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

// 嵌入普通的静态资源
type webui struct {
	webuiEmbed embed.FS // 静态资源
	path       string   // 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
}

// 静态资源被访问的核心逻辑
func (w *webui) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(w.path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := w.webuiEmbed.Open(fullName)
	return file, err
}

// Ant Design前端页面，需要该方式，实现刷新，访问到前端index.html
type webuiIndex struct {
	webui *webui
}

func (w *webuiIndex) Open(name string) (fs.File, error) {
	return w.webui.Open("index.html")
}
