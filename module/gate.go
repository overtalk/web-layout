package module

import (
	"github.com/gin-gonic/gin"
)

type Gate interface {
	Start()
	Shutdown()
	AddGroup(relativePath string, handlers ...gin.HandlerFunc)
	GET(relativePath string, handler interface{}, group ...string)
	POST(relativePath string, handler interface{}, group ...string)
	Static(urlPrefix, root string)
}
