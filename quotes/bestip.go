package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto/ext"
	"gitee.com/quant1x/gotdx/proto/std"
	"gitee.com/quant1x/gox/api"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	TDX_HOST_HQ = "HQ"
	TDX_HOST_EX = "EX"
	TDX_HOST_GP = "GP"
)

const (
	maxCrossTime = 50 // 最大耗时
)

// ServerGroup 主机组
type ServerGroup struct {
	HQ []Server `json:"HQ"`
	EX []Server `json:"EX"`
	GP []Server `json:"GP"`
}

// AllServers 全部主机
type AllServers struct {
	//Server ServerGroup `json:"Server"`
	BestIP ServerGroup `json:"BestIP"`
}

// ProfileBestIPList 测试最快的服务器
func ProfileBestIPList() *AllServers {
	var as AllServers

	// HQ-servers
	dst := cleanServers(StandardServerList, testHQ)
	as.BestIP.HQ = dst

	// EX-server, reply提示版本不一致, 扩展服务暂不可用
	dst = cleanServers(ExtensionServerList, testEX)
	as.BestIP.EX = dst

	//// SP-servers
	//dst = cleanServers(GP_HOSTS, testEX)
	//as.BestIP.GP = dst
	str, _ := json.Marshal(as)
	fmt.Println(string(str))
	return &as
}

func cleanServers(src []Server, test func(addr string) error) (dst []Server) {
	//err := json.Unmarshal([]byte(str), &src)
	//if err != nil {
	//	return src, dst
	//}
	//fmt.Printf("%+v\n", src)

	dst = slices.Clone(src)
	for i, _ := range dst {
		v := &dst[i]
		fmt.Printf("%d: %+v\n", i, v)
		_ = detect(v, test)
		fmt.Printf("%d: %+v\n", i, v)
	}

	sort.Slice(dst, func(i, j int) bool {
		return dst[i].CrossTime < dst[j].CrossTime
	})
	dst = api.Filter(dst, func(e Server) bool {
		return e.CrossTime < maxCrossTime
	})
	num := len(dst)
	if num > POOL_MAX {
		num = POOL_MAX
	}
	dst = dst[0:num]
	fmt.Println(dst)
	return
}

// 检测, 返回毫秒
func detect(srv *Server, test func(addr string) error) int64 {
	var crossTime int64 = math.MaxInt64
	addr := strings.Join([]string{srv.Host, strconv.Itoa(srv.Port)}, ":")
	start := time.Now()
	err := test(addr)
	if err != nil {
		srv.CrossTime = crossTime
		return crossTime
	}
	// 计算耗时, 纳秒
	crossTime = int64(time.Since(start))
	// 转成毫秒
	srv.CrossTime = crossTime / int64(time.Millisecond)
	return crossTime
}

// 标准服务器测试
func testHQ(addr string) error {
	cli, err := NewClientForTest(addr)
	if err != nil {
		return err
	}
	// CMD信令 1
	data, err := CommandWithConn(cli, func() (req std.Marshaler, resp std.Unmarshaler, err error) {
		req, resp, err = std.NewSetupCmd1()
		return
	})
	fmt.Printf("%+v\n", data)
	_ = cli.Close()
	return err
}

// 扩展服务器测试
func testEX(addr string) error {
	cli, err := NewClientForTest(addr)
	if err != nil {
		return err
	}
	// CMD信令 1
	data, err := CommandWithConn(cli, func() (req std.Marshaler, resp std.Unmarshaler, err error) {
		req, resp, err = ext.NewExCmd1()
		return
	})
	fmt.Printf("%+v\n", data)
	_ = cli.Close()
	return err
}
