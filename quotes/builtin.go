package quotes

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"math"
	"strings"
	"sync/atomic"
	"time"
)

// 局部变量
var (
	// 序列号
	_seqId uint32
)

// 生成序列号
func seqID() uint32 {
	atomic.AddUint32(&_seqId, 1)
	return _seqId
}

func Utf8ToGbk(text []byte) string {
	r := bytes.NewReader(text)
	decoder := transform.NewReader(r, simplifiedchinese.GBK.NewDecoder()) //GB18030
	content, _ := io.ReadAll(decoder)
	return strings.ReplaceAll(string(content), string([]byte{0x00}), "")
}

// pytdx : 类似utf-8的编码方式保存有符号数字
func getprice(b []byte, pos *int) int {

	//0x7f与常量做与运算实质是保留常量（转换为二进制形式）的后7位数，既取值区间为[0,127]
	//0x3f与常量做与运算实质是保留常量（转换为二进制形式）的后6位数，既取值区间为[0,63]
	//
	//0x80 1000 0000
	//0x7f 0111 1111
	//0x40  100 0000
	//0x3f  011 1111

	posByte := 6
	bData := b[*pos]
	data := int(bData & 0x3f)
	bSign := false
	if (bData & 0x40) > 0 {
		bSign = true
	}

	if (bData & 0x80) > 0 {
		for {
			*pos += 1
			bData = b[*pos]
			data += (int(bData&0x7f) << posByte)

			posByte += 7

			if (bData & 0x80) <= 0 {
				break
			}
		}
	}
	*pos++

	if bSign {
		data = -data
	}
	return data
}

func gettime(b []byte, pos *int) (h uint16, m uint16) {
	var sec uint16
	_ = binary.Read(bytes.NewBuffer(b[*pos:*pos+2]), binary.LittleEndian, &sec)
	h = sec / 60
	m = sec % 60
	*pos += 2
	return
}

func getdatetime(category int, b []byte, pos *int) (year int, month int, day int, hour int, minute int) {
	hour = 15
	if category < 4 || category == 7 || category == 8 {
		var zipday, tminutes uint16
		_ = binary.Read(bytes.NewBuffer(b[*pos:*pos+2]), binary.LittleEndian, &zipday)
		*pos += 2
		_ = binary.Read(bytes.NewBuffer(b[*pos:*pos+2]), binary.LittleEndian, &tminutes)
		*pos += 2

		year = int((zipday >> 11) + 2004)
		month = int((zipday % 2048) / 100)
		day = int((zipday % 2048) % 100)
		hour = int(tminutes / 60)
		minute = int(tminutes % 60)
	} else {
		var zipday uint32
		_ = binary.Read(bytes.NewBuffer(b[*pos:*pos+4]), binary.LittleEndian, &zipday)
		*pos += 4
		year = int(zipday / 10000)
		month = int((zipday % 10000) / 100)
		day = int(zipday % 100)
	}
	return
}

func getdatetimenow(category int, lasttime string) (year int, month int, day int, hour int, minute int) {
	utime, _ := time.Parse("2006-01-02 15:04:05", lasttime)
	switch category {
	case proto.KLINE_TYPE_5MIN:
		utime = utime.Add(time.Minute * 5)
	case proto.KLINE_TYPE_15MIN:
		utime = utime.Add(time.Minute * 15)
	case proto.KLINE_TYPE_30MIN:
		utime = utime.Add(time.Minute * 30)
	case proto.KLINE_TYPE_1HOUR:
		utime = utime.Add(time.Hour)
	case proto.KLINE_TYPE_DAILY:
		utime = utime.AddDate(0, 0, 1)
	case proto.KLINE_TYPE_WEEKLY:
		utime = utime.Add(time.Hour * 24 * 7)
	case proto.KLINE_TYPE_MONTHLY:
		utime = utime.AddDate(0, 1, 0)
	case proto.KLINE_TYPE_EXHQ_1MIN:
		utime = utime.Add(time.Minute)
	case proto.KLINE_TYPE_1MIN:
		utime = utime.Add(time.Minute)
	case proto.KLINE_TYPE_RI_K:
		utime = utime.AddDate(0, 0, 1)
	case proto.KLINE_TYPE_3MONTH:
		utime = utime.AddDate(0, 3, 0)
	case proto.KLINE_TYPE_YEARLY:
		utime = utime.AddDate(1, 0, 0)
	}

	if category < 4 || category == 7 || category == 8 {
		if (utime.Hour() >= 15 && utime.Minute() > 0) || (utime.Hour() > 15) {
			utime = utime.AddDate(0, 0, 1)
			utime = utime.Add(time.Minute * 30)
			hour = (utime.Hour() + 18) % 24
		} else {
			hour = utime.Hour()
		}
		minute = utime.Minute()
	} else {
		if utime.Unix() > time.Now().Unix() {
			utime = time.Now()
		}
		hour = utime.Hour()
		minute = utime.Minute()
		if utime.Hour() > 15 {
			hour = 15
			minute = 0
		}
	}
	year = utime.Year()
	month = int(utime.Month())
	day = utime.Day()
	return
}

func getvolume(ivol int) (volume float64) {
	logpoint := ivol >> (8 * 3)
	//hheax := ivol >> (8 * 3)          // [3]
	hleax := (ivol >> (8 * 2)) & 0xff // [2]
	lheax := (ivol >> 8) & 0xff       //[1]
	lleax := ivol & 0xff              //[0]

	//dbl_1 := 1.0
	//dbl_2 := 2.0
	//dbl_128 := 128.0

	dwEcx := logpoint*2 - 0x7f
	dwEdx := logpoint*2 - 0x86
	dwEsi := logpoint*2 - 0x8e
	dwEax := logpoint*2 - 0x96
	tmpEax := dwEcx
	if dwEcx < 0 {
		tmpEax = -dwEcx
	} else {
		tmpEax = dwEcx
	}

	dbl_xmm6 := 0.0
	dbl_xmm6 = math.Pow(2.0, float64(tmpEax))
	if dwEcx < 0 {
		dbl_xmm6 = 1.0 / dbl_xmm6
	}

	dbl_xmm4 := 0.0
	dbl_xmm0 := 0.0

	if hleax > 0x80 {
		tmpdbl_xmm3 := 0.0
		//tmpdbl_xmm1 := 0.0
		dwtmpeax := dwEdx + 1
		tmpdbl_xmm3 = math.Pow(2.0, float64(dwtmpeax))
		dbl_xmm0 = math.Pow(2.0, float64(dwEdx)) * 128.0
		dbl_xmm0 += float64(hleax&0x7f) * tmpdbl_xmm3
		dbl_xmm4 = dbl_xmm0
	} else {
		if dwEdx >= 0 {
			dbl_xmm0 = math.Pow(2.0, float64(dwEdx)) * float64(hleax)
		} else {
			dbl_xmm0 = (1 / math.Pow(2.0, float64(dwEdx))) * float64(hleax)
		}
		dbl_xmm4 = dbl_xmm0
	}

	dbl_xmm3 := math.Pow(2.0, float64(dwEsi)) * float64(lheax)
	dbl_xmm1 := math.Pow(2.0, float64(dwEax)) * float64(lleax)
	if (hleax & 0x80) > 0 {
		dbl_xmm3 *= 2.0
		dbl_xmm1 *= 2.0
	}
	volume = dbl_xmm6 + dbl_xmm4 + dbl_xmm3 + dbl_xmm1
	return
}

func baseUnit(code string) float64 {
	c := code[:2]
	switch c {
	case "60", "30", "68", "00":
		return 100.0

	}
	return 1000.0
}

func _format_time(time_stamp string) string {
	// format time from reversed_bytes0
	// by using method from https://github.com/rainx/pytdx/issues/187
	length := len(time_stamp)
	tm := time_stamp[:length-6] + ":"
	tmp := time_stamp[length-6 : length-4]
	n := api.ParseInt(tmp)
	if n < 60 {
		tm += fmt.Sprintf("%s:", tmp)
		tmp = time_stamp[length-4:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%06.3f", (f*60.0)/10000.00)
	} else {
		tmp = time_stamp[length-6:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%02f.:", (f*60.0)/1000000)
		tm += fmt.Sprintf("%06.3f", float64(int64(f)*60%1000000)*60/1000000.0)
	}
	return tm
}
