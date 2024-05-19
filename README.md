# gotdx
golang实现的一个通达信数据协议库

## 1. 概要
- 整合了[gotdx](https://github.com/bensema/gotdx.git)和[TdxPy](https://github.com/rainx/pytdx)
- 增加了连接池的功能
- 自动探测主机网络速度
- 调用简单

## 2. 第一次使用, 获取日K线
第一次运行时, 连接池会探测服务器网络速度会慢一些, 网络测速后会缓存到本地。

#### 2.1. 引入gotdx
```go
import (
"fmt"
"gitee.com/quant1x/gotdx"
"gitee.com/quant1x/gotdx/proto"
)
```

#### 2.2. 获取1根K线数据
```go
    api := gotdx.GetTdxApi()
klines, err := api.GetKLine("sh600600", proto.KLINE_TYPE_RI_K, 0, 1)
fmt.Println(err)
fmt.Println(klines)
```
