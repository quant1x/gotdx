package quotes

import (
	"encoding/json"
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/exchange/cache"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/coroutine"
	"os"
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
	sortedServerListFileName = "tdx.json"
)

var (
	serverType             string
	onceSortServers        coroutine.RollingOnce
	cachedSortedServerList []Server
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

func GetFastHost(inKey string) []Server {
	serverType = inKey
	onceSortServers.Do(lazyCachedSortedServerList)
	return cachedSortedServerList
}

func getServerListByKey(bestIp ServerGroup, key string) *[]Server {
	if key == TDX_HOST_HQ {
		if len(bestIp.HQ) > 0 {
			return &bestIp.HQ
		} else {
			return &[]Server{DefaultHQServer}
		}
	} else if key == TDX_HOST_EX {
		if len(bestIp.EX) > 0 {
			return &bestIp.EX
		} else {
			return &[]Server{DefaultEXServer}
		}
	}

	// Should not reach
	return nil
}

func lazyCachedSortedServerList() {
	target := cache.GetMetaPath() + "/" + sortedServerListFileName
	var allServers *AllServers

	// 检查缓存文件是否存在
	fs, err := api.GetFileStat(target)
	if err == nil {
		lastModified := fs.LastWriteTime
		cacheLastDay := lastModified.Format(exchange.TradingDayDateFormat)

		if cacheLastDay < exchange.LastTradeDate() {
			// 缓存过时，重新生成
			allServers = ProfileBestIPList()
		} else {
			// 缓存有效，尝试加载
			allServers = loadSortedServerList(target)
		}
	} else {
		// 缓存文件不存在，重新生成
		allServers = ProfileBestIPList()
	}

	if allServers != nil && len(allServers.BestIP.HQ) > 0 && len(allServers.BestIP.EX) > 0 {
		// 保存有效缓存
		_ = saveSortedServerList(allServers, target)
	}

	cachedSortedServerList = *getServerListByKey(allServers.BestIP, serverType)
}
