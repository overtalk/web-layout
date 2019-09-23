package consul

import (
	"log"
)

// Watch listening to the service in Consul
func (client *Client) Watch() <-chan AvailableServers {
	if len(client.sdConfigs) == 0 {
		return nil
	}

	for _, sdConfig := range client.sdConfigs {
		go func(sdConfig *DiscoveryConfig) {
			if err := sdConfig.plan.Run(client.consulAddr); err != nil {
				log.Printf("Consul Watch Err: %+v\n", err)
			}
		}(sdConfig)
	}
	return client.watchChan
}
