package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	. "web-layout/utils/gin/gate"
)

func main() {
	g := NewGate(8080)
	g.API("/test", "POST", test)

	group1 := g.Group("/group1")
	group1.API("/test1", "GET", test1)
	group1.API("/test2", "GET", test2)

	group2 := g.Group("/group2")
	group2.API("/test3", "GET", test3)
	group2.API("/test4", "GET", test4)

	g.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	g.Stop()
}

func test(ctx context.Context, req *map[string]interface{}) map[string]interface{} {
	fmt.Println(req)
	return map[string]interface{}{"func": "test"}
}
func test1(ctx context.Context) map[string]interface{} { return map[string]interface{}{"func": "test1"} }
func test2(ctx context.Context) map[string]interface{} { return map[string]interface{}{"func": "test2"} }
func test3(ctx context.Context) map[string]interface{} { return map[string]interface{}{"func": "test3"} }
func test4(ctx context.Context) map[string]interface{} { return map[string]interface{}{"func": "test4"} }
