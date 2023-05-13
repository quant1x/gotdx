package quotes

import (
	"encoding/json"
	"gitee.com/quant1x/gox/logger"
	"gitee.com/quant1x/gox/util/homedir"
	"os"
)

const (
	tdx_path = "~/.quant1x/tdx.json"
)

var (
	config_path     string = "/opt/ctp/tdx.json"
	DefaultHQServer Server
	DefaultEXServer Server
)

func init() {
	_path, err := homedir.Expand(tdx_path)
	if err == nil {
		config_path = _path
	}

	DefaultHQServer = Server{
		Name:      "临时主机",
		Host:      "119.147.212.81",
		Port:      7709,
		CrossTime: 0,
	}
	DefaultEXServer = DefaultHQServer
}

func OpenConfig() *AllServers {
	f, err := os.Open(config_path)
	if err != nil {
		return nil
	}
	decoder := json.NewDecoder(f)
	var as AllServers
	err = decoder.Decode(&as)
	if err != nil {
		return nil
	}
	return &as
}

func CacheServers(as AllServers) error {
	data, err := json.Marshal(as)
	if err != nil {
		return err
	}
	err = os.WriteFile(config_path, data, 0644)
	return err
}

func GetFastHost(key string) []Server {
	as := OpenConfig()
	if as == nil {
		logger.Infof("首次执行通达信数据接口, 正在进行服务器测速")
		BestIP()
	}
	as = OpenConfig()
	if as == nil && key == TDX_HOST_HQ {
		return []Server{DefaultHQServer}
	}
	if as == nil && key == TDX_HOST_EX {
		return []Server{DefaultHQServer}
	}

	bestIp := as.BestIP
	if key == TDX_HOST_HQ {
		if len(bestIp.HQ) > 0 {
			return bestIp.HQ
		} else {
			return []Server{DefaultHQServer}
		}
	} else if key == TDX_HOST_EX {
		if len(bestIp.EX) > 0 {
			return bestIp.EX
		} else {
			return []Server{DefaultHQServer}
		}
	}
	return []Server{DefaultHQServer}
}
