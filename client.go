package main

import "net/rpc"

// ProcessPluginRPCClient ProcessPluginRPCClient 정의
type PluginRPCClient struct{ client *rpc.Client }

// Hello RPC 테스트 함수
func (s *PluginRPCClient) Hello() string {
	var resp string
	err := s.client.Call("Plugin.Hello", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}
