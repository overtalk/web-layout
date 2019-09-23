package demo

import (
	"context"

	"web-layout/module"
)

type Demo struct{}

func Registry(gate module.Gate) {
	d := Demo{}
	gate.GET("/test", d.DemoTest)
}

func (demo *Demo) DemoTest(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"key": "value",
	}
}
