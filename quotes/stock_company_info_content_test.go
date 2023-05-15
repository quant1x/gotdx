package quotes

import (
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/util/linkedhashmap"
	"strings"
	"testing"
)

func TestCompanyInfoContentPackage(t *testing.T) {
	stdApi, err := NewStdApi()
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	reply, err := stdApi.GetCompanyInfoContent(proto.MarketIdShangHai, "600977", "资金动向")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Printf("%+v\n", reply)
	//data, _ := json.Marshal(reply)
	//text := api.Bytes2String(data)
	//fmt.Println(text)
	dict := reply.Map("基本资料")
	dict.Each(func(key interface{}, value interface{}) {
		fmt.Println(key, value)
	})

}

func TestParseCompanyInfo(t *testing.T) {
	content := `公司概况☆ ◇600105 永鼎股份 更新日期：2023-05-03◇ 通达信沪深京F10\r\n★本栏包括【1.基本资料】【2.发行和交易】【3.员工效益】【4.研发投入】\r\n          【5.参股控股】\r\n\r\n【1.基本资料】\r\n┌───────┬───────────────────────────────┐\r\n│公司名称      │江苏永鼎股份有限公司                                          │\r\n├───────┼───────────────────────────────┤\r\n│英文全称      │Jiangsu Etern Company Limited                                 │\r\n├───────┼───────────┬───────┬───────────┤\r\n│证券简称      │永鼎股份              │证券代码      │600105                │\r\n├───────┼───────────┴───────┴───────────┤\r\n│曾用简称      │永鼎光缆-\u003eG永鼎-\u003e永鼎光缆                                     │\r\n├───────┼───────────────────────────────┤\r\n│关联上市      │---                                                           │\r\n├───────┼───────────┬───────┬───────────┤\r\n│通达信研究行业│通信-通信设备         │证监会行业    │电气机械和器材制造业  │\r\n├───────┼───────────┼───────┼───────────┤\r\n│证券类别      │上交所A股             │上市日期      │1997-09-29            │\r\n├───────┼───────────┼───────┼───────────┤\r\n│法人代表      │路庆海                │总经理        │路庆海                │\r\n├───────┼───────────┼───────┼───────────┤\r\n│公司董秘      │张国栋                │证券事务代表  │陈海娟                │\r\n├───────┼───────────┴───────┴───────────┤\r\n│会计事务所    │亚太(集团)会计师事务所(特殊普通合伙)                          │\r\n├───────┼───────────┬───────┬───────────┤\r\n│联系电话      │0512-63271201         │传真          │0512-63271866         │\r\n├───────┼───────────┴───────┴───────────┤\r\n│公司网址      │www.yongding.com.cn                                           │\r\n├───────┼───────────────────────────────┤\r\n│电子邮箱      │zqb@yongding.com.cn                                           │\r\n├───────┼───────────────────────────────┤\r\n│注册地址      │江苏省苏州市吴江区黎里镇318国道74K处芦墟段北侧                │\r\n├───────┼───────────────────────────────┤\r\n│办公地址      │江苏省苏州市吴江区黎里镇318国道74K处芦墟段北侧                │\r\n├───────┼───────────────────────────────┤\r\n│经营范围      │电线、电缆、光纤预制棒、光纤、光缆、配电开关控制设备、电子产品│\r\n│              │、通信设备、汽车及零部件的研究、制造，国内贸易，实业投资，实物│\r\n│              │租赁，自营和代理各类商品及技术的进出口业务，机电工程技术服务，│\r\n│              │企业管理咨询，铜制材和铜加工（冷加工）及其铜产品的销售，移动通│\r\n│              │信设备开发生产及销售，计算机系统及网络技术服务，通信信息网络系│\r\n│              │统集成，承包境外与出口自产设备相关的工程和境内国际招标工程，新│\r\n│              │能源汽车线束的研发生产及销售，信息科技领域内光电器件技术研发、│\r\n│              │生产、销售和相关技术服务，对外派遣实施上述境外工程所需的劳务人│\r\n│              │员。（依法须经批准的项目，经相关部门批准后方可开展经营活动）。│\r\n├───────┼───────────────────────────────┤\r\n│主营业务      │通信光缆、电缆。                                              │\r\n├───────┼───────────────────────────────┤\r\n│公司简介      │公司系由上海贝尔电话设备制造有限公司于1994年6月30日发起设立， │\r\n│              │始经发起人净资产及现金投入折为发起人股3600万股，并定向募集法人│\r\n│              │股1275万股、内部职工股125万股，经1997年9月15日发行后，上市时总│\r\n│              │股本达13500万股，其内部职工股250万股将于公众股3500万股1997年9 │\r\n│              │月29日在上交所上市交易期满三年后上市。                        │\r\n└───────┴───────────────────────────────┘\r\n\r\n【2.发行和交易】\r\n┌─────────┬─────────┬─────────┬─────────┐\r\n│股票类别          │A股               │发行制度          │核准制            │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│网上发行日期      │1997-09-15        │上市日期          │1997-09-29        │\r\n├─────────┼─────────┴─────────┴─────────┤\r\n│发行方式          │上网定价发行                                              │\r\n├─────────┼─────────┬─────────┬─────────┤\r\n│发行量(万股)      │3500.00           │发行价格(元)      │7.10              │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│发行费用(万)      │853.98            │发行总市值(万)    │24850.00          │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│每股面值          │1.00元            │募集资金净额(万)  │23996.03          │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│上市首日开盘价    │13.13             │上市首日收盘价    │12.10             │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│摊薄发行市盈率    │12.46             │加权发行市盈率    │---               │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│发行前每股净资产  │1.32              │发行后每股净资产  │2.75              │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│网上定价中签率%   │0.23              │网下配售中签率%   │---               │\r\n├─────────┼─────────┴─────────┴─────────┤\r\n│主承销商          │国泰证券有限公司                                          │\r\n├─────────┼─────────────────────────────┤\r\n│保荐人            │国泰证券有限公司,南方证券股份有限公司                     │\r\n└─────────┴─────────────────────────────┘\r\n\r\n【3.员工效益】 截止日期:2022-12-31，本期在职员工总数:4053，较上期变动:-2.97%\r\n┌─────────┬─────┬─────┬─────┬─────┬─────┐\r\n│指标/日期         │    2022年│    2021年│    2020年│    2019年│    2018年│\r\n├─────────┼─────┼─────┼─────┼─────┼─────┤\r\n│人均扣非净利润(元)│    1.81万│   6078.39│  -14.54万│  -5113.14│    3.16万│\r\n│人均营业总收入(元)│  104.31万│   93.60万│   79.20万│   77.94万│   75.39万│\r\n│人均薪酬(元)      │   14.22万│   13.44万│   11.76万│   11.72万│   11.06万│\r\n└─────────┴─────┴─────┴─────┴─────┴─────┘\r\n\r\n【4.研发投入】\r\n┌─────────┬─────┬─────┬─────┬─────┬─────┐\r\n│指标/日期         │2022-12-31│2021-12-31│2020-12-31│2019-12-31│2018-12-31│\r\n├─────────┼─────┼─────┼─────┼─────┼─────┤\r\n│研发人数(人)      │       736│       768│       740│       658│       718│\r\n│较上期变动(人)    │       -32│        28│        82│       -60│       105│\r\n│研发人员占比(%)   │     18.16│     18.39│     16.50│     15.21│     16.80│\r\n│研发投入:         │          │          │          │          │          │\r\n│   金额(元)       │    2.11亿│    1.83亿│    1.91亿│    1.78亿│    1.68亿│\r\n│   占营收比(%)    │      5.00│      4.69│      5.80│      5.29│      5.22│\r\n│资本化研发投入:   │          │          │          │          │          │\r\n│   金额(元)       │ 2689.61万│ 2746.48万│ 2361.57万│ 2019.21万│ 1962.64万│\r\n│   占研发投入比(%)│     12.72│     14.98│     12.38│     11.33│     11.67│\r\n│费用化研发投入:   │          │          │          │          │          │\r\n│   金额(元)       │    1.85亿│    1.56亿│    1.67亿│    1.58亿│    1.49亿│\r\n│   占研发投入比(%)│     87.28│     85.02│     87.62│     88.67│     88.33│\r\n└─────────┴─────┴─────┴─────┴─────┴─────┘\r\n\r\n【5.参股控股】(前30) 截止日期:2022-12-31 共30家\r\n┌────────────┬────┬──────┬──────┬───────┐\r\n│关联方名称              │参控关系│ 持股比例(%)│投资金额(元)│      主营业务│\r\n├────────────┼────┼──────┼──────┼───────┤\r\n│上海东昌投资发展有限公司│联营企业│       50.00│     11.82亿│  汽车、房地产│\r\n│上海金亭汽车线束有限公司│子公司  │      100.00│      6.29亿│        制造业│\r\n│北京永鼎致远网络科技有限│子公司  │      100.00│      5.34亿│        软件业│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎投资有限公司    │子公司  │      100.00│   8000.00万│      实业投资│\r\n│江苏永鼎泰富工程有限公司│子公司  │       51.00│   7011.82万│      工程施工│\r\n│东部超导科技（苏州）有限│子公司  │      100.00│   6468.92万│        制造业│\r\n│公司                    │        │            │            │              │\r\n│上海数码通宽带网络有限公│子公司  │      100.00│   5992.17万│      网络服务│\r\n│司                      │        │            │            │              │\r\n│江苏永鼎光纤科技有限公司│子公司  │      100.00│   5046.18万│        制造业│\r\n│苏州永鼎线缆科技有限公司│子公司  │      100.00│   5042.23万│        制造业│\r\n│武汉永鼎汇谷科技有限公司│子公司  │      100.00│   5033.39万│        制造业│\r\n│北京永鼎欣益信息技术有限│子公司  │      100.00│   4515.27万│          商业│\r\n│公司                    │        │            │            │              │\r\n│江苏永鼎电气有限公司    │子公司  │      100.00│   4459.26万│        制造业│\r\n│武汉永鼎光电子集团有限公│子公司  │       75.00│   4030.57万│        制造业│\r\n│司                      │        │            │            │              │\r\n│上海永鼎光电子技术有限公│子公司  │      100.00│   3566.54万│        制造业│\r\n│司                      │        │            │            │              │\r\n│江苏永鼎盛达电缆有限公司│子公司  │       70.00│   2121.77万│        制造业│\r\n│北京新碳和能源有限公司  │联营企业│         ---│   1482.45万│           ---│\r\n│苏州永鼎源臻股权投资管理│子公司  │      100.00│   1138.21万│          投资│\r\n│有限公司                │        │            │            │              │\r\n│北京永鼎致远信息技术有限│孙公司  │      100.00│   1000.00万│        软件业│\r\n│公司                    │        │            │            │              │\r\n│上海巍尼电气工程有限公司│子公司  │      100.00│    990.00万│          贸易│\r\n│苏州鼎诚汽车零部件有限公│子公司  │       60.00│    619.27万│        制造业│\r\n│司                      │        │            │            │              │\r\n│北京永鼎科技发展有限公司│子公司  │      100.00│    362.86万│          贸易│\r\n│永鼎寰宇（国际）有限公司│子公司  │      100.00│    335.41万│          商业│\r\n│武汉永鼎光通科技有限公司│联营企业│         ---│    156.20万│           ---│\r\n│永鼎海缆（南通）有限公司│子公司  │      100.00│     50.00万│          贸易│\r\n│金亭汽车线束（武汉）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│金亭汽车线束（苏州）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│华东超导检测（江苏）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎国际贸易有限公司│孙公司  │      100.00│         ---│          贸易│\r\n│苏州永鼎一园物业管理有限│子公司  │      100.00│         ---│      物业管理│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎智在云科技有限公│孙公司  │      100.00│         ---│        制造业│\r\n│司                      │        │            │            │              │\r\n└────────────┴────┴──────┴──────┴───────┘\r\n\r\n〖免责条款〗\r\n 1、本公司力求但不保证提供的任何信息的真实性、准确性、完整性及原创性等，投资者使\r\n 用前请自行予以核实，如有错漏请以中国证监会指定上市公司信息披露媒体为准，本公司\r\n 不对因上述信息全部或部分内容而引致的盈亏承担任何责任。\r\n 2、本公司无法保证该项服务能满足用户的要求，也不担保服务不会受中断，对服务的及时\r\n 性、安全性以及出错发生都不作担保。\r\n 3、本公司提供的任何信息仅供投资者参考，不作为投资决策的依据，本公司不对投资者依\r\n 据上述信息进行投资决策所产生的收益和损失承担任何责任。投资有风险，应谨慎至上。`
	c := content
	c = strings.ReplaceAll(c, "┌", "") // 左上角
	c = strings.ReplaceAll(c, "┬", "") // 中上

	c = strings.ReplaceAll(c, "┐", "") // 右上角

	//c = strings.ReplaceAll(c, "│", "") // 左边
	c = strings.ReplaceAll(c, "└", "") // 左下角
	c = strings.ReplaceAll(c, "┘", "") // 右下角
	c = strings.ReplaceAll(c, "┼", "")
	c = strings.ReplaceAll(c, "┴", "")
	c = strings.ReplaceAll(c, "├", "")
	c = strings.ReplaceAll(c, "┤", "")
	c = strings.ReplaceAll(c, "─", "")
	c = strings.ReplaceAll(c, "-\\u003e", "->")
	//c = strings.ReplaceAll(c, " ", "")
	c = strings.ReplaceAll(c, "\\r\\n\\r\\n", ";")
	arr := strings.Split(c, ";")
	for _, v := range arr {
		v = strings.ReplaceAll(v, " ", "")
		v = strings.ReplaceAll(v, "│\\r\\n││", "")
		v = strings.Trim(v, "│")
		fmt.Println(v)
	}
}

func TestParseCompanyInfoParse(t *testing.T) {
	content := `公司概况☆ ◇600105 永鼎股份 更新日期：2023-05-03◇ 通达信沪深京F10\r\n★本栏包括【1.基本资料】【2.发行和交易】【3.员工效益】【4.研发投入】\r\n          【5.参股控股】\r\n\r\n【1.基本资料】\r\n┌───────┬───────────────────────────────┐\r\n│公司名称      │江苏永鼎股份有限公司                                          │\r\n├───────┼───────────────────────────────┤\r\n│英文全称      │Jiangsu Etern Company Limited                                 │\r\n├───────┼───────────┬───────┬───────────┤\r\n│证券简称      │永鼎股份              │证券代码      │600105                │\r\n├───────┼───────────┴───────┴───────────┤\r\n│曾用简称      │永鼎光缆-\u003eG永鼎-\u003e永鼎光缆                                     │\r\n├───────┼───────────────────────────────┤\r\n│关联上市      │---                                                           │\r\n├───────┼───────────┬───────┬───────────┤\r\n│通达信研究行业│通信-通信设备         │证监会行业    │电气机械和器材制造业  │\r\n├───────┼───────────┼───────┼───────────┤\r\n│证券类别      │上交所A股             │上市日期      │1997-09-29            │\r\n├───────┼───────────┼───────┼───────────┤\r\n│法人代表      │路庆海                │总经理        │路庆海                │\r\n├───────┼───────────┼───────┼───────────┤\r\n│公司董秘      │张国栋                │证券事务代表  │陈海娟                │\r\n├───────┼───────────┴───────┴───────────┤\r\n│会计事务所    │亚太(集团)会计师事务所(特殊普通合伙)                          │\r\n├───────┼───────────┬───────┬───────────┤\r\n│联系电话      │0512-63271201         │传真          │0512-63271866         │\r\n├───────┼───────────┴───────┴───────────┤\r\n│公司网址      │www.yongding.com.cn                                           │\r\n├───────┼───────────────────────────────┤\r\n│电子邮箱      │zqb@yongding.com.cn                                           │\r\n├───────┼───────────────────────────────┤\r\n│注册地址      │江苏省苏州市吴江区黎里镇318国道74K处芦墟段北侧                │\r\n├───────┼───────────────────────────────┤\r\n│办公地址      │江苏省苏州市吴江区黎里镇318国道74K处芦墟段北侧                │\r\n├───────┼───────────────────────────────┤\r\n│经营范围      │电线、电缆、光纤预制棒、光纤、光缆、配电开关控制设备、电子产品│\r\n│              │、通信设备、汽车及零部件的研究、制造，国内贸易，实业投资，实物│\r\n│              │租赁，自营和代理各类商品及技术的进出口业务，机电工程技术服务，│\r\n│              │企业管理咨询，铜制材和铜加工（冷加工）及其铜产品的销售，移动通│\r\n│              │信设备开发生产及销售，计算机系统及网络技术服务，通信信息网络系│\r\n│              │统集成，承包境外与出口自产设备相关的工程和境内国际招标工程，新│\r\n│              │能源汽车线束的研发生产及销售，信息科技领域内光电器件技术研发、│\r\n│              │生产、销售和相关技术服务，对外派遣实施上述境外工程所需的劳务人│\r\n│              │员。（依法须经批准的项目，经相关部门批准后方可开展经营活动）。│\r\n├───────┼───────────────────────────────┤\r\n│主营业务      │通信光缆、电缆。                                              │\r\n├───────┼───────────────────────────────┤\r\n│公司简介      │公司系由上海贝尔电话设备制造有限公司于1994年6月30日发起设立， │\r\n│              │始经发起人净资产及现金投入折为发起人股3600万股，并定向募集法人│\r\n│              │股1275万股、内部职工股125万股，经1997年9月15日发行后，上市时总│\r\n│              │股本达13500万股，其内部职工股250万股将于公众股3500万股1997年9 │\r\n│              │月29日在上交所上市交易期满三年后上市。                        │\r\n└───────┴───────────────────────────────┘\r\n\r\n【2.发行和交易】\r\n┌─────────┬─────────┬─────────┬─────────┐\r\n│股票类别          │A股               │发行制度          │核准制            │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│网上发行日期      │1997-09-15        │上市日期          │1997-09-29        │\r\n├─────────┼─────────┴─────────┴─────────┤\r\n│发行方式          │上网定价发行                                              │\r\n├─────────┼─────────┬─────────┬─────────┤\r\n│发行量(万股)      │3500.00           │发行价格(元)      │7.10              │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│发行费用(万)      │853.98            │发行总市值(万)    │24850.00          │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│每股面值          │1.00元            │募集资金净额(万)  │23996.03          │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│上市首日开盘价    │13.13             │上市首日收盘价    │12.10             │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│摊薄发行市盈率    │12.46             │加权发行市盈率    │---               │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│发行前每股净资产  │1.32              │发行后每股净资产  │2.75              │\r\n├─────────┼─────────┼─────────┼─────────┤\r\n│网上定价中签率%   │0.23              │网下配售中签率%   │---               │\r\n├─────────┼─────────┴─────────┴─────────┤\r\n│主承销商          │国泰证券有限公司                                          │\r\n├─────────┼─────────────────────────────┤\r\n│保荐人            │国泰证券有限公司,南方证券股份有限公司                     │\r\n└─────────┴─────────────────────────────┘\r\n\r\n【3.员工效益】 截止日期:2022-12-31，本期在职员工总数:4053，较上期变动:-2.97%\r\n┌─────────┬─────┬─────┬─────┬─────┬─────┐\r\n│指标/日期         │    2022年│    2021年│    2020年│    2019年│    2018年│\r\n├─────────┼─────┼─────┼─────┼─────┼─────┤\r\n│人均扣非净利润(元)│    1.81万│   6078.39│  -14.54万│  -5113.14│    3.16万│\r\n│人均营业总收入(元)│  104.31万│   93.60万│   79.20万│   77.94万│   75.39万│\r\n│人均薪酬(元)      │   14.22万│   13.44万│   11.76万│   11.72万│   11.06万│\r\n└─────────┴─────┴─────┴─────┴─────┴─────┘\r\n\r\n【4.研发投入】\r\n┌─────────┬─────┬─────┬─────┬─────┬─────┐\r\n│指标/日期         │2022-12-31│2021-12-31│2020-12-31│2019-12-31│2018-12-31│\r\n├─────────┼─────┼─────┼─────┼─────┼─────┤\r\n│研发人数(人)      │       736│       768│       740│       658│       718│\r\n│较上期变动(人)    │       -32│        28│        82│       -60│       105│\r\n│研发人员占比(%)   │     18.16│     18.39│     16.50│     15.21│     16.80│\r\n│研发投入:         │          │          │          │          │          │\r\n│   金额(元)       │    2.11亿│    1.83亿│    1.91亿│    1.78亿│    1.68亿│\r\n│   占营收比(%)    │      5.00│      4.69│      5.80│      5.29│      5.22│\r\n│资本化研发投入:   │          │          │          │          │          │\r\n│   金额(元)       │ 2689.61万│ 2746.48万│ 2361.57万│ 2019.21万│ 1962.64万│\r\n│   占研发投入比(%)│     12.72│     14.98│     12.38│     11.33│     11.67│\r\n│费用化研发投入:   │          │          │          │          │          │\r\n│   金额(元)       │    1.85亿│    1.56亿│    1.67亿│    1.58亿│    1.49亿│\r\n│   占研发投入比(%)│     87.28│     85.02│     87.62│     88.67│     88.33│\r\n└─────────┴─────┴─────┴─────┴─────┴─────┘\r\n\r\n【5.参股控股】(前30) 截止日期:2022-12-31 共30家\r\n┌────────────┬────┬──────┬──────┬───────┐\r\n│关联方名称              │参控关系│ 持股比例(%)│投资金额(元)│      主营业务│\r\n├────────────┼────┼──────┼──────┼───────┤\r\n│上海东昌投资发展有限公司│联营企业│       50.00│     11.82亿│  汽车、房地产│\r\n│上海金亭汽车线束有限公司│子公司  │      100.00│      6.29亿│        制造业│\r\n│北京永鼎致远网络科技有限│子公司  │      100.00│      5.34亿│        软件业│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎投资有限公司    │子公司  │      100.00│   8000.00万│      实业投资│\r\n│江苏永鼎泰富工程有限公司│子公司  │       51.00│   7011.82万│      工程施工│\r\n│东部超导科技（苏州）有限│子公司  │      100.00│   6468.92万│        制造业│\r\n│公司                    │        │            │            │              │\r\n│上海数码通宽带网络有限公│子公司  │      100.00│   5992.17万│      网络服务│\r\n│司                      │        │            │            │              │\r\n│江苏永鼎光纤科技有限公司│子公司  │      100.00│   5046.18万│        制造业│\r\n│苏州永鼎线缆科技有限公司│子公司  │      100.00│   5042.23万│        制造业│\r\n│武汉永鼎汇谷科技有限公司│子公司  │      100.00│   5033.39万│        制造业│\r\n│北京永鼎欣益信息技术有限│子公司  │      100.00│   4515.27万│          商业│\r\n│公司                    │        │            │            │              │\r\n│江苏永鼎电气有限公司    │子公司  │      100.00│   4459.26万│        制造业│\r\n│武汉永鼎光电子集团有限公│子公司  │       75.00│   4030.57万│        制造业│\r\n│司                      │        │            │            │              │\r\n│上海永鼎光电子技术有限公│子公司  │      100.00│   3566.54万│        制造业│\r\n│司                      │        │            │            │              │\r\n│江苏永鼎盛达电缆有限公司│子公司  │       70.00│   2121.77万│        制造业│\r\n│北京新碳和能源有限公司  │联营企业│         ---│   1482.45万│           ---│\r\n│苏州永鼎源臻股权投资管理│子公司  │      100.00│   1138.21万│          投资│\r\n│有限公司                │        │            │            │              │\r\n│北京永鼎致远信息技术有限│孙公司  │      100.00│   1000.00万│        软件业│\r\n│公司                    │        │            │            │              │\r\n│上海巍尼电气工程有限公司│子公司  │      100.00│    990.00万│          贸易│\r\n│苏州鼎诚汽车零部件有限公│子公司  │       60.00│    619.27万│        制造业│\r\n│司                      │        │            │            │              │\r\n│北京永鼎科技发展有限公司│子公司  │      100.00│    362.86万│          贸易│\r\n│永鼎寰宇（国际）有限公司│子公司  │      100.00│    335.41万│          商业│\r\n│武汉永鼎光通科技有限公司│联营企业│         ---│    156.20万│           ---│\r\n│永鼎海缆（南通）有限公司│子公司  │      100.00│     50.00万│          贸易│\r\n│金亭汽车线束（武汉）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│金亭汽车线束（苏州）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│华东超导检测（江苏）有限│孙公司  │      100.00│         ---│        制造业│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎国际贸易有限公司│孙公司  │      100.00│         ---│          贸易│\r\n│苏州永鼎一园物业管理有限│子公司  │      100.00│         ---│      物业管理│\r\n│公司                    │        │            │            │              │\r\n│苏州永鼎智在云科技有限公│孙公司  │      100.00│         ---│        制造业│\r\n│司                      │        │            │            │              │\r\n└────────────┴────┴──────┴──────┴───────┘\r\n\r\n〖免责条款〗\r\n 1、本公司力求但不保证提供的任何信息的真实性、准确性、完整性及原创性等，投资者使\r\n 用前请自行予以核实，如有错漏请以中国证监会指定上市公司信息披露媒体为准，本公司\r\n 不对因上述信息全部或部分内容而引致的盈亏承担任何责任。\r\n 2、本公司无法保证该项服务能满足用户的要求，也不担保服务不会受中断，对服务的及时\r\n 性、安全性以及出错发生都不作担保。\r\n 3、本公司提供的任何信息仅供投资者参考，不作为投资决策的依据，本公司不对投资者依\r\n 据上述信息进行投资决策所产生的收益和损失承担任何责任。投资有风险，应谨慎至上。`
	c := content
	//c = strings.ReplaceAll(c, "\\r\\n", "\n")
	c = strings.ReplaceAll(c, "-\\u003e", "->")
	arr := strings.Split(c, "\\r\\n\\r\\n")
	for i, block := range arr {
		block = strings.TrimSpace(block)
		//v = strings.ReplaceAll(v, " ", "")
		//v = strings.ReplaceAll(v, "│\\r\\n││", "")
		//v = strings.Trim(v, "│")
		//fmt.Println(i, block)
		if i > 0 && strings.Index(block, "基本资料") >= 0 {
			arr := strings.Split(block, "\\r\\n")
			block = ""
			for _, v := range arr {
				if strings.Index(v, "基本资料") >= 0 {
					continue
				}
				if strings.Index(v, "┌") >= 0 && strings.Index(v, "┐") >= 0 {
					continue
				}
				if strings.Index(v, "└") >= 0 && strings.Index(v, "┘") >= 0 {
					continue
				}
				if strings.Index(v, "├") >= 0 && strings.Index(v, "┤") >= 0 {
					continue
				}
				v = strings.TrimLeft(v, "│")
				v = strings.TrimRight(v, "│")
				v = strings.TrimSpace(v)
				v = strings.ReplaceAll(v, "│", "|")
				//v = strings.TrimLeft(v, "|")
				if v[0] == '|' {
					block += v[1:]
				} else {
					block += "|" + v
				}
			}
			list := strings.Split(block[1:], "|")
			mapInfo := linkedhashmap.New()
			for k := 0; k < len(list); k += 2 {
				key := strings.TrimSpace(list[k])
				value := strings.TrimSpace(list[k+1])
				mapInfo.Put(key, value)
			}
			mapInfo.Each(func(key interface{}, value interface{}) {
				fmt.Println(key, value)
			})

			break
		}
	}
}
