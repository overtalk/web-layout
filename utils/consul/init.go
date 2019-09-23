package consul

import (
	"fmt"
	"net/http"

	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"
)

// Client defines the consul client
type Client struct {
	consulAddr string // consul address（127.0.0.1:8500）

	// service registration related
	srConfig     *RegistryConfig
	checkPort    int // 服务注册 check 的端口
	checkServer  *http.Server
	consulClient *consulAPI.Client // consul Client

	// service discovery related
	sdConfigs []*DiscoveryConfig
	watchChan chan AvailableServers
}

// NewClient is the constructor of consul Client
func NewClient(checkPort int, srConfig *RegistryConfig, consulAddr string, sdConfigs ...*DiscoveryConfig) (*Client, error) {
	// service registry
	c, err := consulAPI.NewClient(&consulAPI.Config{Address: consulAddr})
	if err != nil {
		return nil, err
	}

	// service discovery channel
	watchChan := make(chan AvailableServers, 100)

	// service discovery
	for _, sdConfig := range sdConfigs {
		// 构建plan
		params := make(map[string]interface{})
		params["type"] = "service"
		params["service"] = sdConfig.ServerType
		params["tag"] = sdConfig.Tags
		plan, err := consulWatch.Parse(params)
		if err != nil {
			return nil, err
		}
		plan.Handler = sdConfig.handler

		// bind plan to DiscoveryConfig
		sdConfig.watchChan = watchChan
		sdConfig.plan = plan
	}

	// construct check sever
	mux := http.NewServeMux()
	mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	checkServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", checkPort),
		Handler: mux,
	}

	return &Client{
		checkPort:    checkPort,
		checkServer:  checkServer,
		consulAddr:   consulAddr,
		srConfig:     srConfig,
		sdConfigs:    sdConfigs,
		consulClient: c,
		watchChan:    watchChan,
	}, nil
}
