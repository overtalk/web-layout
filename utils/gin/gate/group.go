package gate

import (
	"github.com/gin-gonic/gin"
)

// 一个 api
type api struct {
	url     string
	method  string
	handler interface{}
}

// Group 相当于 gin 中的 Group
type group struct {
	relativeURL string
	middlewares []gin.HandlerFunc
	subGroups   []*group
	apis        []*api
}

// newGroup 构造函数
func newGroup(r string, m ...gin.HandlerFunc) *group {
	return &group{
		relativeURL: r,
		middlewares: m,
	}
}

func (g *group) Group(relativeURL string, middlewares ...gin.HandlerFunc) Group {
	sub := newGroup(relativeURL, middlewares...)
	g.subGroups = append(g.subGroups, sub)
	return sub
}

func (g *group) API(url, method string, handler interface{}) {
	g.apis = append(g.apis, &api{
		url:     url,
		method:  method,
		handler: handler,
	})
}

func (g *group) add2gin(group *gin.RouterGroup) {
	// 将 group 加入 gin engine, 首先新建 *gin.RouterGroup
	ginGroup := group.Group(g.relativeURL, g.middlewares...)
	// 加入 apis
	g.addApis(ginGroup)
	// 加入 sub group
	g.addGroups(ginGroup)
}

func (g *group) addApis(ginGroup *gin.RouterGroup) {
	// add api
	for _, api := range g.apis {
		if api != nil {
			switch api.method {
			case "POST":
				ginGroup.POST(api.url, Wrap(api.handler))
			case "GET":
				ginGroup.GET(api.url, Wrap(api.handler))
			default:
				// todo: invalid http method
				panic(" invalid http method")
			}
		} else {
			// todo: nil api
			panic("nil api")
		}

	}
}

func (g *group) addGroups(ginGroup *gin.RouterGroup) {
	// add groups
	for _, group := range g.subGroups {
		group.add2gin(ginGroup)
	}
}
