package util


import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)
import "fmt"

func ConsulKVTest() {
	// Get a new client
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &consulapi.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

func RegisterUserService()  {
	// Get a new client
	consulConfig := consulapi.Config{
		Address:    "127.0.0.1:8500",
	}
	client, err := consulapi.NewClient(&consulConfig)
	if err != nil {
		panic(err)
	}

	serviceConfig := consulapi.AgentServiceRegistration{
		ID:                "userservice",
		Name:              "userservice",
		Port:              8080,
		Address:           "127.0.0.1",
		Tags:[]string{"primary"},
		Check:nil,
	}

	check := consulapi.AgentServiceCheck{
		HTTP: "http://127.0.0.1:8080/health",
		Timeout:"10s",
		Interval:"5s",
	}
	serviceConfig.Check = &check

	if err := client.Agent().ServiceRegister(&serviceConfig); err != nil {
		panic(err)
	}
}


func DeregisterUserService()  {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	if err := client.Agent().ServiceDeregister("userservice"); err != nil {
		log.Println(err)
	}
}