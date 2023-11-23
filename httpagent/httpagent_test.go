package httpagent_test

import (
	"SDCS/httpagent"
	"testing"
)

func TestHttpAgent(t *testing.T) {
	// 1. 初始化agent
	agent := httpagent.NewHttpAgent(0, "9527")
	// 2. 启动agent
	agent.StartHttpAgent()
}
