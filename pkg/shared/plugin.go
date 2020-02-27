package shared

import (
	"context"

	"github.com/wukaiying/goplugin-learn/pkg/plugins/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

/**
定义plugin和host对接的接口
 */

//plugin和host握手协议
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	MagicCookieKey: "WKY_PLUGIN",
	MagicCookieValue: "wukaiying",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"kv_grpc": &KVGRPCPlugin{},
}

// KV is the interface that we're exposing as a plugin.
type KV interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type KVGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl KV
}

func (p *KVGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterKVServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *KVGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewKVClient(c)}, nil
}



