package gate

import (
	"github.com/gin-gonic/gin"
)

// Gate 接口
type Gate interface {
	Group
	Static(urlPrefix, root string, indexes bool)
	Start()
	Stop()
}

// Group 接口
type Group interface {
	Group(relativeURL string, middlewares ...gin.HandlerFunc) Group
	API(url, method string, handler interface{})
}
