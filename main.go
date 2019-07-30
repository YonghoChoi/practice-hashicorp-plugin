package main

import (
	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	// ToDo : Plugin에 필요한 기본 정보(projectID, hostGroupID 등)를 초기화할 수 있는 함수 및 설정파일 필요
	p := new(Plugin)
	// pluginMap에 등록된 Plugin들은 클라이언트에서 이름으로 참조하여 dispense 할 수 있음
	var pluginMap = map[string]plugin.Plugin{
		"example": p,
	}

	// 플러그인 실행
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
