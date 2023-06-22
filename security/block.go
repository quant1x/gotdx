package security

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

func loadCacheBlockInfos() {
	syncBlockFiles()
	bkFilename := cache.BlockFilename()
	list := []BlockInfo{}
	err := api.CsvToSlices(bkFilename, &list)
	if err != nil {
		return
	}
	if len(list) > 0 {
		for _, v := range list {
			// 对齐板块代码
			blockCode := proto.CorrectSecurityCode(v.Code)
			v.Code = blockCode
			for i := 0; i < len(v.ConstituentStocks); i++ {
				// 对齐个股代码
				stockCode := proto.CorrectSecurityCode(v.ConstituentStocks[i])
				v.ConstituentStocks[i] = stockCode
			}
			// 缓存列表
			__global_block_list = append(__global_block_list, v)
			// 缓存板块映射关系
			__mapBlock[v.Code] = v
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
