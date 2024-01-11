package securities

import (
	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/exchange/cache"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/gox/api"
	"os"
)

// 同步板块数据
func syncBlockFiles() {
	downloadBlockRawData(quotes.BLOCK_DEFAULT)
	downloadBlockRawData(quotes.BLOCK_GAINIAN)
	downloadBlockRawData(quotes.BLOCK_FENGGE)
	downloadBlockRawData(quotes.BLOCK_ZHISHU)
	updateCacheBlockFile()
}

// 更新本地板块文件
func v1UpdateBlockFile() {
	// 如果板块数据不存在, 从应用内导出
	blockFile := cache.BlockFilename()
	createOrUpdate := false
	if !api.FileExist(blockFile) {
		createOrUpdate = true
	}
	if !createOrUpdate {
		blockData := cache.GetBlockPath() + "/" + quotes.BLOCK_DEFAULT
		dataStat, err := os.Stat(blockData)
		if err == nil || os.IsExist(err) {
			dataModifyTime := dataStat.ModTime()
			// 检查通达信热股880818板块文件
			bk880818 := cache.GetBlockPath() + "/" + "880818.csv"
			bkStat, err := os.Stat(bk880818)
			if err == nil || os.IsExist(err) {
				if bkStat.ModTime().Before(dataModifyTime) {
					createOrUpdate = true
				}
			} else {
				createOrUpdate = true
			}
		} else {
			createOrUpdate = true
		}
	}
	if createOrUpdate {
		parseAndGenerateBlockFile()
	}
}

// 更新缓存csv数据文件
func updateCacheBlockFile() {
	// 如果板块数据不存在, 从应用内导出
	blockFile := cache.BlockFilename()
	createOrUpdate := false
	if !api.FileExist(blockFile) {
		createOrUpdate = true
	} else {
		dataStat, err := os.Stat(blockFile)
		if err == nil || os.IsExist(err) {
			dataModifyTime := dataStat.ModTime()
			toInit := exchange.CanInitialize(dataModifyTime)
			if toInit {
				createOrUpdate = true
			}
		} else {
			createOrUpdate = true
		}
	}
	if createOrUpdate {
		parseAndGenerateBlockFile()
	}
}
