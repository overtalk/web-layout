package consul

import (
	"context"
	"fmt"
	"time"

	consulAPI "github.com/hashicorp/consul/api"
)

var ErrAbsentServiceRegisterConfig = fmt.Errorf("service register config is absent")

// Register implements registry client interface
func (c *Registrar) Register() error {
	if c.srConfig == nil {
		return ErrAbsentServiceRegisterConfig
	}

	go func() {
		if err := c.checkServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = c.srConfig.ID
	registration.Name = c.srConfig.ServerType
	registration.Tags = c.srConfig.Tags
	registration.Address = c.srConfig.IP
	registration.Port = c.srConfig.Port
	registration.Check = &consulAPI.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, c.checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", //check失败后15秒删除本服务
	}

	return c.consulClient.Agent().ServiceRegister(registration)
}

func (c *Registrar) DeRegister() error {
	if c.srConfig == nil {
		return ErrAbsentServiceRegisterConfig
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c.checkServer.Shutdown(ctx)
	return c.consulClient.Agent().ServiceDeregister(c.srConfig.ID)
}
