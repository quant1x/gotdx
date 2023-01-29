package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto/v1"
	v2 "gitee.com/quant1x/gotdx/proto/v2"
	"github.com/mymmsc/gox/util/lambda"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	// HQ_HOSTS 标准市场 主机列表
	HQ_HOSTS = `[
{"name":"深圳双线主站1", "host":"110.41.147.114", "port": 7709},
{"name":"深圳双线主站2", "host":"8.129.13.54", "port": 7709},
{"name":"深圳双线主站3", "host":"120.24.149.49", "port": 7709},
{"name":"深圳双线主站4", "host":"47.113.94.204", "port": 7709},
{"name":"深圳双线主站5", "host":"8.129.174.169", "port": 7709},
{"name":"深圳双线主站6", "host":"110.41.154.219", "port": 7709},
{"name":"上海双线主站1", "host":"124.70.176.52", "port": 7709},
{"name":"上海双线主站2", "host":"47.100.236.28", "port": 7709},
{"name":"上海双线主站3", "host":"101.133.214.242", "port": 7709},
{"name":"上海双线主站4", "host":"47.116.21.80", "port": 7709},
{"name":"上海双线主站5", "host":"47.116.105.28", "port": 7709},
{"name":"上海双线主站6", "host":"124.70.199.56", "port": 7709},
{"name":"北京双线主站1", "host":"121.36.54.217", "port": 7709},
{"name":"北京双线主站2", "host":"121.36.81.195", "port": 7709},
{"name":"北京双线主站3", "host":"123.249.15.60", "port": 7709},
{"name":"广州双线主站1", "host":"124.71.85.110", "port": 7709},
{"name":"广州双线主站2", "host":"139.9.51.18", "port": 7709},
{"name":"广州双线主站3", "host":"139.159.239.163", "port": 7709},
{"name":"上海双线主站7", "host":"106.14.201.131", "port": 7709},
{"name":"上海双线主站8", "host":"106.14.190.242", "port": 7709},
{"name":"上海双线主站9", "host":"121.36.225.169", "port": 7709},
{"name":"上海双线主站10", "host":"123.60.70.228", "port": 7709},
{"name":"上海双线主站11", "host":"123.60.73.44", "port": 7709},
{"name":"上海双线主站12", "host":"124.70.133.119", "port": 7709},
{"name":"上海双线主站13", "host":"124.71.187.72", "port": 7709},
{"name":"上海双线主站14", "host":"124.71.187.122", "port": 7709},
{"name":"武汉电信主站1", "host":"119.97.185.59", "port": 7709},
{"name":"深圳双线主站7", "host":"47.107.64.168", "port": 7709},
{"name":"北京双线主站4", "host":"124.70.75.113", "port": 7709},
{"name":"广州双线主站4", "host":"124.71.9.153", "port": 7709},
{"name":"上海双线主站15", "host":"123.60.84.66", "port": 7709},
{"name":"深圳双线主站8", "host":"47.107.228.47", "port": 7719},
{"name":"北京双线主站5", "host":"120.46.186.223", "port": 7709},
{"name":"北京双线主站6", "host":"124.70.22.210", "port": 7709},
{"name":"北京双线主站7", "host":"139.9.133.247", "port": 7709},
{"name":"广州双线主站5", "host":"116.205.163.254", "port": 7709},
{"name":"广州双线主站6", "host":"116.205.171.132", "port": 7709},
{"name":"广州双线主站7", "host":"116.205.183.150", "port": 7709}
]`

	// EX_HOSTS 扩展市场主机列表
	EX_HOSTS = `[
{"name":"扩展市场深圳双线1", "host":"112.74.214.43", "port": 7727},
{"name":"扩展市场深圳双线2", "host":"120.24.0.77", "port": 7727},
{"name":"扩展市场深圳双线3", "host":"47.107.75.159", "port": 7727},
{"name":"扩展市场武汉主站1", "host":"119.97.185.5", "port": 7727},
{"name":"扩展市场武汉主站2", "host":"202.103.36.71", "port": 7727},
{"name":"扩展市场武汉主站3", "host":"59.175.238.38", "port": 7727},
{"name":"扩展市场北京双线0", "host":"47.92.127.181", "port": 7727},
{"name":"扩展市场上海双线0", "host":"106.14.95.149", "port": 7727},
{"name":"扩展市场新加双线0", "host":"119.23.127.172", "port": 7727}
]`
	// GP_HOSTS 财务数据 主机列表
	GP_HOSTS = `[
{"name":"默认财务数据线路", "host":"120.76.152.87", "port": 7709}
]`

	CONFIG = `{
"SERVER": {"HQ": HQ_HOSTS, "EX": EX_HOSTS, "GP": GP_HOSTS},
"BESTIP": {"HQ": "", "EX": "", "GP": ""},
"TDXDIR": "C:/new_tdx",
}`
)

// Server 主机信息
type Server struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	CrossTime int64  `json:"crossTime"`
}

//type _serers []Server
//
//func (s _serers) Len() int { return len(s) }
//
//func (s _serers) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
//
//func (s _serers) Less(i, j int) bool {
//	if s[i].CrossTime < s[j].CrossTime {
//		return true
//	}
//
//	return false
//}

const (
	TDX_HOST_HQ = "HQ"
	TDX_HOST_EX = "EX"
	TDX_HOST_GP = "GP"
)

// ServerGroup 主机组
type ServerGroup struct {
	HQ []Server `json:"HQ"`
	EX []Server `json:"EX"`
	GP []Server `json:"GP"`
}

// AllServers 全部主机
type AllServers struct {
	Server ServerGroup `json:"Server"`
	BestIP ServerGroup `json:"BestIP"`
}

// BestIP 测试最快的服务器
func BestIP() {
	var as AllServers

	// HQ-servers
	//var hqServers []Server
	//err := json.Unmarshal([]byte(HQ_HOSTS), &hqServers)
	//if err != nil {
	//	fmt.Printf("没有HQ服务器\n")
	//} else {
	//	as.Server.HQ = hqServers
	//	fmt.Printf("%+v\n", hqServers)
	//	for i := 0; i < len(hqServers); i++ {
	//		v := &hqServers[i]
	//		fmt.Printf("%d: %+v\n", i, v)
	//		_ = detect(v)
	//		fmt.Printf("%d: %+v\n", i, v)
	//	}
	//	fmt.Println("----")
	//	hqS := lambda.LambdaArray(hqServers).Sort(func(a Server, b Server) bool {
	//		return a.CrossTime < b.CrossTime
	//	}).Filter(func(e Server) bool { return e.CrossTime < 100 }).Pointer().([]Server)
	//	fmt.Println(hqS)
	//}

	// HQ-servers
	src, dst := cleanServers(HQ_HOSTS, testHQ)
	as.Server.HQ = src
	as.BestIP.HQ = dst
	// EX-server, reply提示版本不一致, 扩展服务暂不可用
	src, dst = cleanServers(EX_HOSTS, testEX)
	as.Server.EX = src
	as.BestIP.EX = dst

	// SP-servers
	src, dst = cleanServers(GP_HOSTS, testEX)
	as.Server.GP = src
	as.BestIP.GP = dst

	str, _ := json.Marshal(as)
	fmt.Println(string(str))
	_ = CacheServers(as)
}

func cleanServers(str string, test func(addr string)) (src, dst []Server) {
	err := json.Unmarshal([]byte(str), &src)
	if err != nil {
		return src, dst
	}
	fmt.Printf("%+v\n", src)
	for i := 0; i < len(src); i++ {
		v := &src[i]
		fmt.Printf("%d: %+v\n", i, v)
		_ = detect(v, test)
		fmt.Printf("%d: %+v\n", i, v)
	}
	//dst = lambda.LambdaArray(src).Sort(func(a Server, b Server) bool {
	//	return a.CrossTime < b.CrossTime
	//}).Filter(func(e Server) bool { return e.CrossTime < 100 }).Pointer().([]Server)
	dst1 := lambda.LambdaArray(src).Sort(func(a Server, b Server) bool {
		return a.CrossTime < b.CrossTime
	})
	dst2 := dst1.Filter(func(e Server) bool { return e.CrossTime < 100 })
	dst = dst2.Take(0, 2).Pointer().([]Server)
	fmt.Println(dst)
	return
}

// 检测, 返回毫秒
func detect(srv *Server, test func(addr string)) int64 {
	var crossTime int64 = math.MaxInt64
	addr := strings.Join([]string{srv.Host, strconv.Itoa(srv.Port)}, ":")
	start := time.Now()
	test(addr)
	// 计算耗时, 纳秒
	crossTime = int64(time.Since(start))
	// 转成毫秒
	srv.CrossTime = crossTime / int64(time.Millisecond)
	return crossTime
}

// 标准服务器测试
func testHQ(addr string) {
	cli, err := v1.NewClientForTest(addr)
	if err == nil {
		// CMD信令 1
		data := CommandWithConn(cli, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
			req, resp, err = v1.NewSetupCmd1()
			return
		})
		fmt.Printf("%+v\n", data)
		_ = cli.Close()
	}
}

// 扩展服务器测试
func testEX(addr string) {
	cli, err := v1.NewClientForTest(addr)
	if err == nil {
		// CMD信令 1
		data := CommandWithConn(cli, func() (req v1.Marshaler, resp v1.Unmarshaler, err error) {
			req, resp, err = v2.NewExCmd1()
			return
		})
		fmt.Printf("%+v\n", data)
		_ = cli.Close()
	}
}
