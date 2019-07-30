package main

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Plugin struct {
}

// Server RPC Server
func (s *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ProcessPluginRPCServer{Impl: s}, nil
}

// Client RPC Client
func (Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPCClient{client: c}, nil
}

// Hello RPC 테스트 함수
func (s Plugin) Hello() string {
	return "hello plugin"
}
