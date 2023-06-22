package security

type BlockType = int

const (
	BK_UNKNOWN BlockType = 0  // 未知类型
	BK_HANGYE  BlockType = 2  // 行业
	BK_DIQU    BlockType = 3  // 地区
	BK_GAINIAN BlockType = 4  // 概念
	BK_FENGGE  BlockType = 5  // 风格
	BK_ZHISHU  BlockType = 6  // 指数
	BK_YJHY    BlockType = 12 // 研究行业

	BKN_HANGYE  = "行业"
	BKN_DIQU    = "地区"
	BKN_GAINIAN = "概念"
	BKN_FENGGE  = "风格"
	BKN_ZHISHU  = "指数"
	BKN_YJHY    = "研究行业"
)

var (
	kMapBlock = map[BlockType]string{
		BK_HANGYE:  BKN_HANGYE,
		BK_DIQU:    BKN_DIQU,
		BK_GAINIAN: BKN_GAINIAN,
		BK_FENGGE:  BKN_FENGGE,
		BK_ZHISHU:  BKN_ZHISHU,
		BK_YJHY:    BKN_YJHY,
	}
)

// BlockTypeNameByCode 通过板块类型代码获取板块类型名称
func BlockTypeNameByCode(blockCode int) (name string, ok bool) {
	blockType := BlockType(blockCode)
	return BlockTypeNameByTypeCode(blockType)
}

// BlockTypeNameByTypeCode 通过板块类型代码获取板块类型名称
func BlockTypeNameByTypeCode(blockType BlockType) (string, bool) {
	bkName, found := kMapBlock[blockType]
	return bkName, found
}
