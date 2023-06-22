package block

import (
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"golang.org/x/exp/slices"
	"sync"
)

var (
	__onceBlockFiles    sync.Once
	__global_block_list = []BlockInfo{}
	__mapBlock          = map[string]BlockInfo{}
)

// BlockInfo 板块信息
type BlockInfo struct {
	Name              string   `dataframe:"name"`              // 名称
	Code              string   `dataframe:"code"`              // 代码
	Type              int      `dataframe:"type"`              // 类型
	Count             int      `dataframe:"count"`             // 个股数量
	Block             string   `dataframe:"block"`             // 通达信板块编码
	ConstituentStocks []string `dataframe:"ConstituentStocks"` // 板块成份股
}

type ConstituentStock struct {
	Code string // 股票代码
}

func loadCacheBlockInfos() {
	syncBlockFiles()
	bkFilename := cache.BlockFilename()
	list := []BlockInfo{}
	err := api.CsvToSlices(bkFilename, &list)
	if err != nil {
		return
	}
	if len(list) > 0 {
		__global_block_list = list
		for _, v := range __global_block_list {
			securityCode := proto.CorrectSecurityCode(v.Code)
			__mapBlock[securityCode] = v
		}
	}
}

// BlockList 板块列表
func BlockList() (list []BlockInfo) {
	__onceBlockFiles.Do(loadCacheBlockInfos)
	return slices.Clone(__global_block_list)
}

func GetBlockInfo(code string) *BlockInfo {
	__onceBlockFiles.Do(loadCacheBlockInfos)
	securityCode := code
	if !proto.AssertBlockBySecurityCode(&securityCode) {
		return nil
	}
	blockInfo, ok := __mapBlock[securityCode]
	if ok {
		return &blockInfo
	}
	return nil
}
