package block

import (
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
)

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
				marketId, _, _ := proto.DetectMarket(sc.Code)
				if marketId == proto.MarketIdBeiJing {
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
		filename := cache.BlockFilename()
		_ = api.SlicesToCsv(filename, blockInfos)
	}

	//for _, v := range blockInfos {
	//	if v.Count == 0 {
	//		continue
	//	}
	//	bk_stock := v.ConstituentStocks
	//	codes := pandas.NewSeries(stat.SERIES_TYPE_STRING, "code", bk_stock)
	//	tmp := pandas.NewDataFrame(codes)
	//	if tmp.Nrow() > 0 {
	//		_ = tmp.WriteCSV(cache.GetBlockPath() + "/" + v.Code + ".csv")
	//	}
	//	bk_code = append(bk_code, v.Code)
	//	bk_name = append(bk_name, v.Name)
	//	bk_type = append(bk_type, v.Type)
	//}
	//bkc := pandas.NewSeries(stat.SERIES_TYPE_STRING, "code", bk_code)
	//bkn := pandas.NewSeries(stat.SERIES_TYPE_STRING, "name", bk_name)
	//bkt := pandas.NewSeries(stat.SERIES_TYPE_STRING, "type", bk_type)
	//dfBk := pandas.NewDataFrame(bkc, bkn, bkt)
	//if dfBk.Nrow() > 0 {
	//	_ = dfBk.WriteCSV(cache.BlockFilename())
	//}
}
