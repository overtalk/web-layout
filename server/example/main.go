package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"web-layout/gate"
	"web-layout/service/demo"
	"web-layout/utils/consul"
	"web-layout/utils/version"
)

var (
	name   string
	tag    string
	commit string
	branch string
)

func main() {
	var v bool
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		git := version.Info{
			Name:   name,
			Tag:    tag,
			Commit: commit,
			Branch: branch,
		}
		fmt.Println(git.Version())
		return
	}

	g := gate.NewGate(9999)
	g.AddConsul(consulClient())
	g.Static("/", "/Users/qinhan/web-layout/static/example")
	demo.Registry(g)

	g.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down gate server...")
		g.Shutdown()
	}()

	g.Shutdown()
}

func consulClient() *consul.Client {
	consulAddr := "127.0.0.1:8500"
	r, err := consul.NewClient(consulAddr)
	if err != nil {
		panic(err)
	}

	r.ServiceRegistry(9989, &consul.RegistryConfig{
		IP:         "127.0.0.1",
		ID:         "1",
		Port:       944,
		ServerType: "Example",
		Tags:       []string{"0.98", "QQ"},
	})
	return r
}
