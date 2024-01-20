package quotes

import (
	"encoding/json"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/exchange/cache"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"os"
	"path/filepath"
	"time"
)

var (
	DefaultHQServer = Server{
		Name:      "临时主机",
		Host:      "119.147.212.81",
		Port:      7709,
		CrossTime: 0,
	}
	DefaultEXServer = DefaultHQServer
)

const (
	serverListFilename = "tdx.json"
)

var (
	//serverType             string
	onceSortServers coroutine.RollingOnce
	//cachedSortedServerList []Server
	cacheAllServers AllServers
)

func loadSortedServerList(configPath string) *AllServers {
	f, err := os.Open(configPath)
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

func saveSortedServerList(as *AllServers, configPath string) error {
	data, err := json.Marshal(as)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	return err
}

func GetFastHost(key string) []Server {
	onceSortServers.Do(lazyCachedSortedServerList)
	bestIp := cacheAllServers.BestIP
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

func lazyCachedSortedServerList() {
	// 1. 组织文件路径
	filename := filepath.Join(cache.GetMetaPath(), serverListFilename)

	// 2. 检查缓存文件是否存在
	var lastModified time.Time
	fs, err := api.GetFileStat(filename)
	if err == nil {
		lastModified = fs.LastWriteTime
	}
	// 2.2 转换缓存文件最后修改日期, 时间格式和日历格式对齐
	cacheLastDay := lastModified.Format(exchange.TradingDayDateFormat)

	var allServers *AllServers
	// 3. 比较缓存日期和最后一个交易日
	if cacheLastDay < exchange.LastTradeDate() {
		// 缓存过时，重新生成
		allServers = ProfileBestIPList()
	} else {
		// 缓存有效，尝试加载
		allServers = loadSortedServerList(filename)
	}
	// 4. 数据有效, 则缓存文件
	if allServers != nil && len(allServers.BestIP.HQ) > 0 /*&& len(allServers.BestIP.EX) > 0*/ {
		// 保存有效缓存
		_ = saveSortedServerList(allServers, filename)
	} else {
		panic("not found")
	}
	// 5. 更新缓存
	cacheAllServers = *allServers
}
