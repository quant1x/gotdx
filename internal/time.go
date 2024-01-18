package internal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gox/api"
	"time"
)

const (
	__tm_h_width = 1000000
	__tm_m_width = 10000
	__tm_t_width = 1000
)

func _format_time0(time_stamp string) string {
	// format time from reversed_bytes0
	// by using method from https://github.com/rainx/pytdx/issues/187
	length := len(time_stamp)
	t1 := api.ParseInt(time_stamp[:length-6])
	tm := fmt.Sprintf("%02d:", t1)
	tmp := time_stamp[length-6 : length-4]
	n := api.ParseInt(tmp)
	if n < 60 {
		tm += fmt.Sprintf("%02s:", tmp)
		tmp = time_stamp[length-4:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%06.3f", (f*60.0)/10000.00)
	} else {
		tmp = time_stamp[length-6:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%02d:", int64(f*60.0)/1000000)
		n = int64(f)
		tm += fmt.Sprintf("%06.3f", float64((n*60)%1000000)*60/1000000.0)
	}
	return tm
}

func timeFromStr(time_stamp string) string {
	// format time from reversed_bytes0
	// by using method from https://github.com/rainx/pytdx/issues/187
	length := len(time_stamp)
	t1 := api.ParseInt(time_stamp[:length-6])
	tm := fmt.Sprintf("%02d:", t1)
	tmp := time_stamp[length-6 : length-4]
	n := api.ParseInt(tmp)
	if n < 60 {
		tm += fmt.Sprintf("%02s:", tmp)
		tmp = time_stamp[length-4:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%06.3f", (f*60.0)/10000.00)
	} else {
		tmp = time_stamp[length-6:]
		f := api.ParseFloat(tmp)
		tm += fmt.Sprintf("%02d:", int64(f*60.0)/1000000)
		n = int64(f)
		tm += fmt.Sprintf("%06.3f", float64((n*60)%1000000)*60/1000000.0)
	}
	return tm
}

func TimeFromInt(stamp int) string {
	//123456789
	h := stamp / __tm_h_width
	tmp1 := stamp % __tm_h_width
	m1 := tmp1 / __tm_m_width
	tmp2 := tmp1 % __tm_m_width
	m := 0
	s := 0
	t := 0
	st := float64(0.00)
	if m1 < 60 {
		m = m1
		tmp3 := tmp2 * 60
		s = tmp3 / __tm_m_width
		t = tmp3 % __tm_m_width
		t /= 10
		st = float64(tmp3) / __tm_m_width
	} else {
		h += 1
		tmp3 := tmp1
		m = tmp3 / __tm_h_width
		tmp4 := (tmp3 % __tm_h_width) * 60
		s = tmp4 / __tm_h_width
		t = tmp4 % __tm_h_width
		t /= __tm_t_width
		st = float64(tmp4) / __tm_h_width
	}
	_ = s
	_ = t
	return fmt.Sprintf("%02d:%02d:%06.3f", h, m, st)
}

func GetDatetimeFromUint32(category int, zipday uint32, tminutes uint16) (year int, month int, day int, hour int, minute int) {
	hour = 15
	if category < 4 || category == 7 || category == 8 {
		year = int((zipday >> 11) + 2004)
		month = int((zipday % 2048) / 100)
		day = int((zipday % 2048) % 100)
		hour = int(tminutes / 60)
		minute = int(tminutes % 60)
	} else {
		year = int(zipday / 10000)
		month = int((zipday % 10000) / 100)
		day = int(zipday % 100)
	}
	return
}

func getDatetimeNow(category int, lasttime string) (year int, month int, day int, hour int, minute int) {
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

func GetTime(b []byte, pos *int) (h uint16, m uint16) {
	var sec uint16
	_ = binary.Read(bytes.NewBuffer(b[*pos:*pos+2]), binary.LittleEndian, &sec)
	h = sec / 60
	m = sec % 60
	*pos += 2
	return
}

func GetDatetime(category int, b []byte, pos *int) (year int, month int, day int, hour int, minute int) {
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
