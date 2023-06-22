package security

import (
	"gitee.com/quant1x/gotdx"
	"gitee.com/quant1x/gotdx/internal/cache"
	"gitee.com/quant1x/gotdx/trading"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/encoding/binary/struc"
	"gitee.com/quant1x/gox/text/encoding"
	"os"
	"strings"
)

// 下载板块原始数据文件
func downloadBlockRawData(filename string) {
	tdxApi := gotdx.GetTdxApi()
	fn := cache.GetBlockPath() + "/" + filename
	fileInfo, err := os.Stat(fn)
	if err == nil || os.IsExist(err) {
		toInit := trading.CanInitialize(fileInfo.ModTime())
		if !toInit {
			return
		}
	}
	resp, err := tdxApi.GetBlockInfo(filename)
	if err == nil {
		fn := cache.GetBlockPath() + "/" + filename
		_ = api.CheckFilepath(fn, true)
		fp, err := os.Create(fn)
		if err == nil {
			_, _ = fp.Write(resp.Data)
			_ = fp.Close()
		}
	}
}

type __raw_block_info struct {
	BlockName string             `struc:"[9]byte,little"`             // 板块名称
	Num       uint16             `struc:"uint16,little"`              // 个股数量
	BlockType uint16             `struc:"uint16,little"`              // 板块类型
	List      [400]__block_stock `struct:"[400]__block_stock,little"` // 个股列表
}

type __block_stock struct {
	Code string `struc:"[7]byte,little"` // 证券代码
}

type __raw_block_data struct {
	//Header blockHeader `struc:"[386]byte,little"`
	Unknown [384]byte          `struc:"[384]byte,little"`          // 头信息, 忽略
	Count   uint16             `struc:"uint16,little,sizeof=Data"` // 板块数量
	Data    []__raw_block_info `struc:"[2813]byte, little"`        // 板块数据
}

func parseRawBlockData(blockFilename string) *__raw_block_data {
	fn := cache.GetBlockPath() + "/" + blockFilename
	_ = api.CheckFilepath(fn, true)
	file, err := os.Open(fn)
	if err != nil {
		return nil
	}
	defer api.CloseQuietly(file)
	var block __raw_block_data
	err = struc.Unpack(file, &block)
	if err != nil {
		return nil
	}
	decoder := encoding.NewDecoder("GBK")
	for i, v := range block.Data {
		name := decoder.ConvertString(v.BlockName)
		block.Data[i].BlockName = strings.ReplaceAll(name, string([]byte{0x00}), "")
		for j, s := range v.List {
			block.Data[i].List[j].Code = strings.ReplaceAll(s.Code, string([]byte{0x00}), "")
		}
	}
	return &block
}
