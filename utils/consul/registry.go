package consul

import (
	"context"
	"fmt"
	"time"

	consulAPI "github.com/hashicorp/consul/api"
)

var ErrAbsentServiceRegisterConfig = fmt.Errorf("service register config is absent")

// Register implements registry Client interface
func (client *Client) Register() error {
	if client.srConfig == nil {
		return ErrAbsentServiceRegisterConfig
	}

	go func() {
		if err := client.checkServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = client.srConfig.ID
	registration.Name = client.srConfig.ServerType
	registration.Tags = client.srConfig.Tags
	registration.Address = client.srConfig.IP
	registration.Port = client.srConfig.Port
	registration.Check = &consulAPI.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, client.checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", // del this service in 15s after check fail
	}

	return client.consulClient.Agent().ServiceRegister(registration)
}

func (client *Client) DeRegister() error {
	if client.srConfig == nil {
		return ErrAbsentServiceRegisterConfig
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client.checkServer.Shutdown(ctx)
	return client.consulClient.Agent().ServiceDeregister(client.srConfig.ID)
}
