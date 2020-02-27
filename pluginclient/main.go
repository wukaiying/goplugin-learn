package main

import (
	"fmt"
	"goplugin-learn/pkg/shared"
	"io/ioutil"

	"github.com/hashicorp/go-plugin"
)

type KV struct {}

func (KV) Put(key string, value []byte) error {
	value = []byte(fmt.Sprintf("%s\n\nWritten from plugin-go-grpc", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(key string) ([]byte, error) {
	return ioutil.ReadFile("kv_"+key)
}

func main()  {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,			//需要引入host上面定义好的一些变量
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVGRPCPlugin{Impl: &KV{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}