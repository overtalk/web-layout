package gate

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// 一个 staticFile
type staticFile struct {
	urlPrefix string
	root      string
	indexes   bool
}

type gate struct {
	*group
	port   int
	engine *gin.Engine
	srv    *http.Server
	static *staticFile
}

func NewGate(p int, m ...gin.HandlerFunc) Gate {
	g := newGroup("", m...)
	return &gate{
		port:  p,
		group: g,
	}
}

func (g *gate) Start() {
	// 新建 engine
	g.engine = gin.New()
	// 加入 staticFile
	if g.static != nil {
		g.engine.Use(static.Serve(g.static.urlPrefix, static.LocalFile(g.static.root, g.static.indexes)))
	}
	// 添加路由
	g.add2gin(&g.engine.RouterGroup)

	go func(g *gate) {
		g.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", g.port),
			Handler: g.engine,
		}

		fmt.Printf("Server Start on port : %d\n", g.port)
		if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("Server Start Error : " + err.Error())
		}
	}(g)
}

func (g *gate) Stop() {
	// prevent call stop without start
	if g.srv != nil {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := g.srv.Shutdown(ctx); err != nil {
			panic("Server Shutdown Error : " + err.Error())
		}
	}
	fmt.Println("Server Exit")
}

func (g *gate) Static(urlPrefix, root string, indexes bool) {
	g.static = &staticFile{
		urlPrefix: urlPrefix,
		root:      root,
		indexes:   indexes,
	}
}
