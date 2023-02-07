package main

import (
	"github.com/hashicorp/consul/api"
)

func RegisterHttpServer(id, name, address string, port int, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "0.0.0.0:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 配置检查项目
	check := &api.AgentServiceCheck{
		HTTP:                           "http://0.0.0.0:8010/health",
		Timeout:                        "10s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "1m",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	_ = RegisterHttpServer("user-web", "user-web", "0.0.0.0", 8010, []string{"spark", "carl"})

}
