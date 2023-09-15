package functions

import (
	"context"
	"time"

	"gitee.com/quant1x/gotdx/trading"
	"gitee.com/quant1x/gox/logger"
)

const (
	TOTALFZNUM = 240 // 每天总交易分钟数
)

var (
	//__global_context context.Context = nil
	//__global_context_cancel context.CancelFunc = nil
	CURRBARSCOUNT = 1                 // 到最后K线的周期数
	FROMOPEN      = trading.Minutes() // 已开盘分钟数
)

// func init() {
// 	fmt.Println("init ...")
// 	//__global_context, __global_context_cancel = context.WithCancel(context.Background())
// 	//go updateAllConstants()
// }

func updateAllConstants(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 取出值, 结束协程
			logger.Infof("signle, parent contex exiting, time=%s", time.Now().Format(time.DateTime))
			return
		default:
			FROMOPEN = trading.Minutes()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Init(ctx context.Context) {
	go updateAllConstants(ctx)
}
