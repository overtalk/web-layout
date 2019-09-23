package gate

import (
	"web-layout/utils/consul"
)

func (gate *Gate) AddConsul(client *consul.Client) {
	gate.consulClient = client
}
