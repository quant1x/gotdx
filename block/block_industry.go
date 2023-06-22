package block

import (
	"bufio"
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/text/encoding"
	"golang.org/x/exp/slices"
	"io"
	"os"
	"strings"
)

// IndustryInfo 行业板块对应
type IndustryInfo struct {
	Code   string // 股票代码
	Block  string // 通达信板块代码
	Block5 string // 通达信板块代码
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
		hy := IndustryInfo{
			Code:   arr[1],
			Block:  bc,
			Block5: bc5,
		}
		hys = append(hys, hy)
	}
	return hys
}

// 从行业信息中提取股票代码列表
func industryConstituentStockList(hys []IndustryInfo, block string) []string {
	list := []string{}
	for _, v := range hys {
		if len(block) == 5 {
			if v.Block5 == block {
				list = append(list, v.Code)
			}
		} else {
			if v.Block == block {
				list = append(list, v.Code)
			}
		}
	}
	if len(list) > 0 {
		slices.Sort(list)
	}
	return list
}
