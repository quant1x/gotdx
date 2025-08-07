package securities

import (
	"os"

	"gitee.com/quant1x/exchange"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/gox/api"
)

// 同步板块数据
func syncBlockFiles() {
	downloadBlockRawData(quotes.BLOCK_DEFAULT)
	downloadBlockRawData(quotes.BLOCK_GAINIAN)
	downloadBlockRawData(quotes.BLOCK_FENGGE)
	downloadBlockRawData(quotes.BLOCK_ZHISHU)
	updateCacheBlockFile()
}

// 更新缓存csv数据文件
func updateCacheBlockFile() {
	// 如果板块数据不存在, 从应用内导出
	blockFile := SectorFilename()
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
