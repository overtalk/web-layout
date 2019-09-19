package consul

import (
	"log"
)

func (c *Registrar) Watch() <-chan AvailableServers {
	if len(c.sdConfigs) == 0 {
		return nil
	}

	for _, sdConfig := range c.sdConfigs {
		go func(sdConfig *SDConfig) {
			if err := sdConfig.plan.Run(c.consulAddr); err != nil {
				log.Printf("Consul Watch Err: %+v\n", err)
			}
		}(sdConfig)
	}
	return c.watchChan
}
