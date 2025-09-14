package securities

import (
	"fmt"

	"github.com/quant1x/exchange"
	"github.com/quant1x/exchange/cache"
	"github.com/quant1x/x/api"
)

// SectorFilename 板块缓存文件名
func SectorFilename(date ...string) string {
	name := "blocks"
	cacheDate := exchange.LastTradeDate()
	if len(date) > 0 {
		cacheDate = exchange.FixTradeDate(date[0])
	}
	filename := fmt.Sprintf("%s/%s.%s", cache.GetMetaPath(), name, cacheDate)
	return filename
}

// 读取板块数据
func parseAndGenerateBlockFile() {
	blockInfos := loadIndexBlockInfos()
	block2Name := map[string]string{}
	for _, v := range blockInfos {
		block2Name[v.Block] = v.Name
	}
	bks := []string{"block.dat", "block_gn.dat", "block_fg.dat", "block_zs.dat"}
	//bks := []string{"block_gn.dat", "block_fg.dat", "block_zs.dat"}
	name2block := map[string]__raw_block_info{}
	for _, v := range bks {
		bi := parseRawBlockData(v)
		if bi != nil {
			for _, bk := range (*bi).Data {
				blockName := bk.BlockName
				if bn, ok := block2Name[blockName]; ok {
					blockName = bn
				}
				name2block[blockName] = bk
			}
		}
	}
	// 行业代码映射
	code2hy := map[string]string{}
	for _, v := range blockInfos {
		if v.Name != v.Block {
			code2hy[v.Block] = v.Name
		}
	}
	// 行业板块数据
	hys := loadIndustryBlocks()
	for i, v := range blockInfos {
		bn := v.Name
		__info, ok := name2block[bn]
		if ok {
			list := []string{}
			for _, sc := range __info.List {
				if len(sc.Code) < 5 {
					continue
				}
				marketId, _, _ := exchange.DetectMarket(sc.Code)
				if marketId == exchange.MarketIdBeiJing {
					continue
				}
				list = append(list, sc.Code)
			}
			blockInfos[i].Count = int(__info.Num)
			blockInfos[i].ConstituentStocks = list
			continue
		}
		bc := v.Block
		stockList := industryConstituentStockList(hys, bc)
		if len(stockList) > 0 {
			blockInfos[i].Count = len(stockList)
			blockInfos[i].ConstituentStocks = stockList
		}
	}
	blockInfos = api.Filter(blockInfos, func(info BlockInfo) bool {
		return len(info.ConstituentStocks) > 0
	})
	if len(blockInfos) > 0 {
		filename := SectorFilename()
		_ = api.SlicesToCsv(filename, blockInfos)
	}
}
