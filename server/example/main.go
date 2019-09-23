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
	flag.BoolVar(&v, "version", false, "show version")
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
