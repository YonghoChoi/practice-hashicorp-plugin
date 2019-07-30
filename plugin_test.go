package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestServicePlugin(t *testing.T) {
	// 플러그인 내부 로깅을 위한 Hashcorp Logger 생성
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// 바이너리 파일로 RPC Server 프로세스 구동
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin-example.exe"),
		Logger:          logger,
	})
	defer client.Kill()

	// RPC 서버로 연결
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// RPC Client 인스턴스 반환
	raw, err := rpcClient.Dispense("example")
	if err != nil {
		log.Fatal(err)
	}

	// Clinet 인터페이스를 통해 함수 호출
	p := raw.(*PluginRPCClient)
	fmt.Println(p.Hello())
}

var pluginMap = map[string]plugin.Plugin{
	"example": &Plugin{},
}
