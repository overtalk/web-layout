package consul

import (
	"fmt"
	"net/http"

	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"
)

// ConsulClient : 注册器
type Registrar struct {
	// consul 地址
	consulAddr string // consul地址（127.0.0.1:8500）

	// 服务注册相关
	srConfig     *SRConfig
	checkPort    int // 服务注册 check 的端口
	checkServer  *http.Server
	consulClient *consulAPI.Client // consul client

	// 服务发现相关
	sdConfigs []*SDConfig
	watchChan chan AvailableServers
}

// NewRegistrar ： 构造函数
func NewRegistrar(checkPort int, srConfig *SRConfig, consulAddr string, sdConfigs ...*SDConfig) (*Registrar, error) {
	// service registry
	client, err := consulAPI.NewClient(&consulAPI.Config{Address: consulAddr})
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

		// 绑定 plan 到 SDConfig 上
		sdConfig.watchChan = watchChan
		sdConfig.plan = plan
	}

	// 构建 check sever
	mux := http.NewServeMux()
	mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	checkServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", checkPort),
		Handler: mux,
	}

	return &Registrar{
		checkPort:    checkPort,
		checkServer:  checkServer,
		consulAddr:   consulAddr,
		srConfig:     srConfig,
		sdConfigs:    sdConfigs,
		consulClient: client,
		watchChan:    watchChan,
	}, nil
}
