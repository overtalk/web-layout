package consul

import (
	"fmt"

	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"
)

// ConsulClient : consul 客户端
// 向 Consul 注册/取消注册服务
// 监听 Consul 上的服务
type ConsulClient interface {
	Wait()                          // 等待特定的服务上线
	Register() error                // 向consul注册服务
	DeRegister() error              // 向consul取消注册
	Watch() <-chan AvailableServers // 服务watch
}

// AvailableServers ： 在线的可用服务
type AvailableServers struct {
	ServerType string
	Servers    []string
}

// SRConfig : service registry config
// 服务发现配置
type SRConfig struct {
	ID         string   // 服务ID
	IP         string   // 服务地址
	Port       int      // 服务端口
	ServerType string   // 服务类型
	Tags       []string // 服务Tags
}

// SDConfig ：service discovery config
// 服务发现配置
type SDConfig struct {
	ServerType string   // 目标服务类型
	Tags       []string // 目标服务的tags
	// others
	watchChan chan AvailableServers
	plan      *consulWatch.Plan
}

func (c *SDConfig) handler(index uint64, raw interface{}) {
	if raw == nil {
		return
	}
	if entries, ok := raw.([]*consulAPI.ServiceEntry); ok {
		var servers []string
		for _, entry := range entries {
			// 如果服务没有通过健康检查，直接continue
			if entry.Checks.AggregatedStatus() != consulAPI.HealthPassing {
				continue
			}
			servers = append(servers, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
		}
		c.watchChan <- AvailableServers{
			ServerType: c.ServerType,
			Servers:    servers,
		}
	}
}
