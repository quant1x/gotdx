package dfcf

import (
	"encoding/csv"
	"fmt"
	"reflect"
)

type DfcfHistory struct {
	// date        时间
	Date string `json:"date" array:"0"`
	// open       开盘价
	Open float64 `json:"open" array:"1"`
	// high       最高价
	High float64 `json:"high" array:"3"`
	// low        最低价
	Low float64 `json:"low" array:"4"`
	// close      收盘价
	Close float64 `json:"close" array:"2"`
	// volume     成交量, 单位股, 除以100为手
	Volume int64 `json:"volume" array:"5"`
}

type KLine struct {
	Date   string  `json:"date" array:"0" name:"日期" dataframe:"date,string"`
	Open   float64 `json:"open" array:"1" name:"开盘价" dataframe:"open,float64"`
	Close  float64 `json:"close" array:"2" name:"收盘价" dataframe:"close,float64"`
	High   float64 `json:"high" array:"3" name:"最高价" dataframe:"high,float64"`
	Low    float64 `json:"low" array:"4" name:"最低价" dataframe:"low,float64"`
	Volume int64   `json:"volume" array:"5" name:"成交量" dataframe:"volume,int64"`
	Amount float64 `json:"amount" array:"6" name:"成交金额" dataframe:"amount,float64"`
	//Amplitude    float64 `json:"amplitude" array:"7" name:"振幅" dataframe:"-"`
	//RiseAndFall  float64 `json:"rise_and_fall" array:"8" name:"涨跌幅" dataframe:"-"`
	//UpAndDown    float64 `json:"up_and_down" array:"9" name:"涨跌额" dataframe:"-"`
	//TurnoverRate float64 `json:"turnover_rate" array:"10" name:"换手率" dataframe:"-"`
}

func reflectType(i any) reflect.Type {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// Init 初始化
func (k KLine) Init(_writer *csv.Writer) error {
	t := reflectType(k)
	fieldNum := t.NumField()
	var line []string
	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		fn := field.Name
		tag := field.Tag
		if len(tag) > 0 {
			tv, ok := tag.Lookup("json")
			if ok {
				fn = tv
			}
		}
		line = append(line, fn)
	}
	return _writer.Write(line)
}

// WriteCSV 写入
func (k KLine) WriteCSV(_writer *csv.Writer) error {
	val := reflect.ValueOf(k)
	var line []string
	for i := 0; i < val.NumField(); i++ {
		fd := val.Field(i)
		vs := ""
		if fd.Kind() == reflect.Float32 || fd.Kind() == reflect.Float64 {
			vs = fmt.Sprintf("%.03f", fd.Float())
		} else {
			vs = fmt.Sprintf("%v", fd.Interface())
		}
		line = append(line, vs)
	}
	return _writer.Write(line)
}
