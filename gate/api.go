package gate

import (
	"fmt"

	"web-layout/utils/consul"
)

func (gate *Gate) Start() {
	if gate.consulClient != nil {
		if err := gate.consulClient.Register(); err != nil && err != consul.ErrAbsentServiceRegisterConfig {
			panic(err)
		}
		gate.consulClient.Wait()
	}

	gate.Gate.Start()
}

func (gate *Gate) Shutdown() {
	if gate.consulClient != nil {
		fmt.Println("加了consul")
		if err := gate.consulClient.DeRegister(); err != nil && err != consul.ErrAbsentServiceRegisterConfig {
			panic(err)
		}
	}

	gate.Gate.Shutdown()
}
