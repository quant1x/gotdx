package gotdx

import (
	"gitee.com/quant1x/gotdx/quotes"
	"sync"
)

var (
	stdApi   *quotes.StdApi = nil
	tdxMutex sync.Mutex
)

func initTdxApi() {
	if stdApi == nil {
		api_, err := quotes.NewStdApi()
		if err != nil {
			return
		}
		stdApi = api_
	}
}

func GetTdxApi() *quotes.StdApi {
	tdxMutex.Lock()
	defer tdxMutex.Unlock()
	initTdxApi()
	return stdApi
}

func ReOpen() {
	tdxMutex.Lock()
	defer tdxMutex.Unlock()
	if stdApi != nil {
		stdApi.Close()
		stdApi = nil
	}
}
