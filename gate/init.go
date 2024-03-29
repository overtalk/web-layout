package gate

import (
	"github.com/gin-gonic/gin"

	"web-layout/utils/consul"
	"web-layout/utils/gin/gate"
)

type Gate struct {
	consulClient *consul.Client
	*gate.Gate
}

func NewGate(port int, m ...gin.HandlerFunc) *Gate {
	g := &Gate{
		Gate: gate.NewGate(port, m...),
	}
	return g
}
