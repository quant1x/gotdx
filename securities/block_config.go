package securities

import (
	"bufio"
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/text/encoding"
	"io"
	"os"
	"strings"
)

// 加载板块和板块名称对应
func loadIndexBlockInfos() []BlockInfo {
	bks := []string{"tdxzs.cfg", "tdxzs3.cfg"}
	bis := []BlockInfo{}
	tmpMap := map[string]bool{}
	for _, v := range bks {
		bi := getBlockInfoFromConfig(v)
		if len(bi) == 0 {
			continue
		}
		for _, info := range bi {
			if _, ok := tmpMap[info.Code]; !ok {
				bis = append(bis, info)
				tmpMap[info.Code] = true
			}
		}
	}
	return bis
}

func getBlockInfoFromConfig(name string) []BlockInfo {
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
	var blocks = []BlockInfo{}
	for {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := decoder.ConvertString(string(data))
		arr := strings.Split(line, "|")
		bk := BlockInfo{
			Name:  arr[0],
			Code:  arr[1],
			Type:  int(api.ParseInt(arr[2])),
			Block: arr[5],
		}
		blocks = append(blocks, bk)
	}
	return blocks
}
