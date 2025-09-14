package securities

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/quant1x/exchange/cache"
	"github.com/quant1x/x/api"
	"github.com/quant1x/x/text/encoding"
)

// IndustryInfo 行业板块对应
type IndustryInfo struct {
	MarketId int    // 市场代码
	Code     string // 股票代码
	Block    string // 行业板块代码
	Block5   string // 二级行业板块代码
	XBlock   string // x行业代码
	XBlock5  string // x二级行业代码
}

// 获取行业板块
func loadIndustryBlocks() []IndustryInfo {
	hyfile := "tdxhy.cfg"
	name := hyfile
	cacheFilename := cache.GetBlockPath() + "/" + name
	if !api.FileExist(cacheFilename) {
		// 如果文件不存在, 导出内嵌资源
		err := export(cacheFilename, name)
		if err != nil {
			return nil
		}
	}
	file, err := os.Open(cacheFilename)
	if err != nil {
		return nil
	}
	defer api.CloseQuietly(file)
	reader := bufio.NewReader(file)
	// 按行处理txt
	decoder := encoding.NewDecoder("GBK")
	var hys = []IndustryInfo{}
	for {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := decoder.ConvertString(string(data))
		arr := strings.Split(line, "|")
		bc := arr[2]
		bc5 := bc
		if len(bc5) >= 5 {
			bc5 = bc5[0:5]
		}
		var xbc, xbc5 string
		if len(arr) >= 6 {
			xbc5 = arr[5]
			if len(xbc5) >= 6 {
				xbc = xbc5[:5]
			}
		}

		hy := IndustryInfo{
			MarketId: int(api.ParseInt(arr[0])),
			Code:     arr[1],
			Block:    bc,
			Block5:   bc5,
			XBlock:   xbc,
			XBlock5:  xbc5,
		}
		hys = append(hys, hy)
	}
	return hys
}

// 从行业信息中提取股票代码列表
func industryConstituentStockList(hys []IndustryInfo, block string) []string {
	list := []string{}
	for _, v := range hys {
		if strings.HasPrefix(v.Block5, block) || strings.HasPrefix(v.XBlock5, block) {
			list = append(list, v.Code)
		} else if v.Block5 == block || v.Block == block || v.XBlock5 == block || v.XBlock == block {
			list = append(list, v.Code)
		}
	}
	if len(list) > 0 {
		slices.Sort(list)
	}
	return list
}
