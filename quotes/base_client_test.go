package quotes

import (
	"testing"
	"time"
)

func TestTcpClient_heartbeat(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	// 休眠20秒触发超时流程
	time.Sleep(time.Second * 2000)
}
