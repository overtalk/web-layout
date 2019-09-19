package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"web-layout/utils/consul"
)

func main() {
	go listen()
	go service()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func service() {
	checkPort := 9000
	consulAddr := "127.0.0.1:8500"
	service := &consul.SRConfig{
		IP:         "127.0.0.1",
		ID:         "service1.1",
		Port:       9994,
		ServerType: "service1",
		Tags:       []string{"0.98", "QQ"},
	}
	r, err := consul.NewRegistrar(checkPort, service, consulAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("开始注册服务 --- service1")
	if err := r.Register(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("注销服务 --- service1")
	r.DeRegister()
}

func listen() {
	fmt.Println("开始监听服务 --- service1")
	consulAddr := "127.0.0.1:8500"
	r, err := consul.NewRegistrar(0, nil, consulAddr, &consul.SDConfig{
		ServerType: "service1",
		Tags:       []string{"0.98", "QQ"},
	})
	if err != nil {
		log.Fatal(err)
	}

	queue := r.Watch()
	for {
		select {
		case data := <-queue:
			fmt.Println("服务状态发送变化了")
			fmt.Println(data)
		default:
			time.Sleep(time.Second)
		}
	}

}
