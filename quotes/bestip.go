package quotes

import (
	"encoding/json"
	"fmt"
	v2 "gitee.com/quant1x/gotdx/proto/v2"
	"math"
	"sort"
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
{"name":"扩展市场新加双线0", "host":"119.23.127.172", "port": 7727},
]`
	// GP_HOSTS 财务数据 主机列表
	GP_HOSTS = `[
{"name":"默认财务数据线路", "host":"120.76.152.87", "port": 7709},
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

type _serers []Server

func (s _serers) Len() int { return len(s) }

func (s _serers) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s _serers) Less(i, j int) bool {
	if s[i].CrossTime < s[j].CrossTime {
		return true
	}

	return false
}

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
	//json, err := fastjson.Parse(HQ_HOSTS)
	//fmt.Printf("%+v, error=%+v\n", json, err)
	var hqServers []Server
	err := json.Unmarshal([]byte(HQ_HOSTS), &hqServers)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", hqServers)
	for i := 0; i < len(hqServers); i++ {
		v := &hqServers[i]
		fmt.Printf("%d: %+v\n", i, v)
		_ = detect(v)
		//fmt.Printf("%d: %+v, cross time: %d \n", i, v, crossTime)
		fmt.Printf("%d: %+v\n", i, v)
	}
	fmt.Println("----")
	sort.Sort(_serers(hqServers))
	fmt.Println(hqServers)
}

// 检测, 返回纳秒
func detect(srv *Server) int64 {
	var crossTime int64 = math.MaxInt64
	addr := strings.Join([]string{srv.Host, strconv.Itoa(srv.Port)}, ":")
	start := time.Now()
	//conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	cli, err := v2.NewClientForTest(addr)
	if err == nil {
		// CMD信令 1
		data := v2.CommandWithConn(cli, func() (req v2.Marshaler, resp v2.Unmarshaler, err error) {
			req, resp, err = v2.NewSetupCmd1()
			return
		})
		fmt.Printf("%+v\n", data)
		crossTime = int64(time.Since(start))
		_ = cli.Close()
	}
	srv.CrossTime = crossTime
	return crossTime
}
