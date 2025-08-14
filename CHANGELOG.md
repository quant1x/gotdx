# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.25.0] - 2025-08-14
### Changed
- 更新最新的行业板块数据据

## [1.24.0] - 2025-08-14
### Changed
- go版本最低要求1.25
- update changelog

## [1.23.15] - 2025-08-11
### Changed
- 更新最新的行业板块数据据
- update changelog

## [1.23.14] - 2025-08-11
### Changed
- 更新exchange版本到0.6.8
- update changelog

## [1.23.13] - 2025-08-08
### Changed
- sort imports
- 修正行业版本的代码匹配, 补充细分行业的代码
- update changelog

## [1.23.12] - 2025-08-07
### Changed
- sort imports
- 股价计算适配全部上海和深圳的证券代码
- update changelog

## [1.23.11] - 2025-08-06
### Changed
- 更新exchange版本到0.6.7
- update changelog

## [1.23.10] - 2025-08-05
### Changed
- 修正补全ETF价格基数, 调整上海从510xxx改为51xxxx, 深圳159xxx
- update changelog

## [1.23.9] - 2025-07-31
### Changed
- 修复行权价字段名的错误拼写
- update changelog

## [1.23.8] - 2025-07-08
### Changed
- 更新依赖库exchange版本到0.6.6
- update changelog

## [1.23.7] - 2025-07-02
### Changed
- !12 Merge branch 'master' of gitee.com:quant1x/gotdx into optm/timeout
* Merge branch 'master' of gitee.com:quant1x/gotdx into optm/timeout
* 优化bestip 的超时从50ms到1s
- update changelog

## [1.23.6] - 2025-06-25
### Changed
- 新增部分测试用例
- 150006364 会解析成150:00:38.184, 看情况150需要除以10, 是整体除以10还是就小时除以10
- 调整部分字段中文拼音拼写的错误
- 调整部分字段中文拼音拼写的错误
- K线的时间DateTime字段调整为年月日时分秒毫秒
- update changelog

## [1.23.5] - 2025-03-18
### Changed
- 删除废弃的代码段
- update changelog

## [1.23.4] - 2025-03-16
### Changed
- 修复两处异常日志输出参数错误的问题
- update changelog

## [1.23.3] - 2025-03-13
### Changed
- 心跳包改用ticker即时
- update changelog

## [1.23.2] - 2025-03-11
### Changed
- 更新依赖库gox版本到1.22.11
- update changelog

## [1.23.1] - 2025-03-09
### Changed
- 更新依赖库gox版本到1.22.6
- update changelog

## [1.23.0] - 2025-02-15
### Changed
- 更新依赖库exchange版本到0.6.0
- update changelog

## [1.22.23] - 2024-12-27
### Changed
- 更新依赖库exchange版本

## [1.22.22] - 2024-12-27
### Changed
- 更新依赖库exchange版本
- update changelog

## [1.22.21] - 2024-08-06
### Changed
- 更新依赖库版本
- update changelog
- update changelog

## [1.22.20] - 2024-06-27
### Changed
- 除权除息因子去掉四舍五入的计算

## [1.22.19] - 2024-06-27
### Changed
- 恢复BestIP强制更新本地服务器列表缓存的处理方式
- update changelog

## [1.22.18] - 2024-06-26
### Changed
- 调整BestIP工具,定义为强制更新服务器列表
- update changelog

## [1.22.17] - 2024-06-25
### Changed
- 调整BestIP工具
- update changelog

## [1.22.16] - 2024-06-24
### Changed
- 修复服务器列表更新逻辑,跨日期和时间戳都必须符合
- update changelog

## [1.22.15] - 2024-06-21
### Changed
- 修复时间戳判断逻辑，现在时间和观察点时间判断是必须的
- update changelog

## [1.22.14] - 2024-06-20
### Changed
- 更新依赖库exchange版本号到0.5.8
- update changelog

## [1.22.13] - 2024-06-20
### Changed
- 新增服务器列表文件当日缓存时间和约定时间戳的比对逻辑
- update changelog

## [1.22.12] - 2024-06-20
### Changed
- 调整BestIp的默认时间为当前时间戳
- update changelog

## [1.22.11] - 2024-06-19
### Changed
- 优化服务器列表更新逻辑, 如果在当日未过预定时间点, 服务器列表从缓存加载
- update changelog

## [1.22.10] - 2024-06-14
### Changed
- 更新依赖库版本
- update changelog

## [1.22.9] - 2024-05-20
### Changed
- 更新exchange版本到0.5.5
- update changelog
- update changelog

## [1.22.8] - 2024-05-19
### Changed
- 调整测试代码, 新增获取日K线的测试
- 修复过时的演示代码
- 调整演示代码
- update changelog

## [1.22.7] - 2024-05-16
### Changed
- 更新依赖库版本
- update changelog

## [1.22.6] - 2024-05-11
### Changed
- 更新依赖库版本
- update changelog

## [1.22.5] - 2024-05-11
### Changed
- 更新依赖库num版本
- update changelog

## [1.22.4] - 2024-05-11
### Changed
- 修复日志方法的错误用法
- 修复连接失败的时候存在服务器地址指针为空的bug
- update changelog

## [1.22.3] - 2024-04-16
### Changed
- 更新exchange版本到0.5.2
- update changelog

## [1.22.2] - 2024-04-12
### Changed
- 更新exchange版本到0.5.0
- git仓库忽略csv文件
- 快照新增检测竞价阶段的方法
- update changelog

## [1.22.1] - 2024-04-10
### Changed
- 新增两融标的判断函数
- update changelog

## [1.22.0] - 2024-04-10
### Changed
- 更新依赖库版本
- update changelog

## [1.21.9] - 2024-03-30
### Changed
- 更新依赖库版本
- update changelog

## [1.21.8] - 2024-03-21
### Changed
- 更新板块配置信息
- update changelog

## [1.21.7] - 2024-03-19
### Changed
- 调整测试代码, 使用指定的服务器地址
- 更新板块配置信息
- update changelog

## [1.21.6] - 2024-03-18
### Changed
- 更新依赖库exchange版本
- update changelog

## [1.21.5] - 2024-03-17
### Changed
- 更新依赖库版本
- update changelog

## [1.21.4] - 2024-03-12
### Changed
- 更新依赖库版本及go版本
- update changelog

## [1.21.3] - 2024-03-12
### Changed
- 更新依赖库版本
- update changelog

## [1.21.2] - 2024-03-11
### Changed
- 更新依赖库版本
- 更新依赖库版本
- update changelog

## [1.21.1] - 2024-02-28
### Changed
- 更新依赖库版本
- update changelog

## [1.21.0] - 2024-02-12
### Changed
- 更新依赖库版本
- update changelog

## [1.20.9] - 2024-02-05
### Changed
- 修复板块名称存在乱码的现象。问题的原因是数据接口的证券名称长度为8个字节，而部分板块名称则超过了8个字节。修改的方案是用板块数据中的名称代替证券列表中的名称。
- 更新依赖库pandas版本
- update changelog

## [1.20.8] - 2024-01-27
### Changed
- 修订通达信业务2次握手的告警类日志内容
- 优化连接池及心跳协程的退出机制, 统一由客户端close方法处理资源释放
- 更新依赖库版本
- update changelog

## [1.20.7] - 2024-01-25
### Changed
- 更新依赖库版本
- update changelog

## [1.20.6] - 2024-01-25
### Changed
- 更新gox,优化RollingOnce功能
- 更新gox,优化RollingOnce功能
- update changelog

## [1.20.5] - 2024-01-24
### Changed
- 恢复连接池只初始化一次
- update changelog

## [1.20.4] - 2024-01-24
### Changed
- 设置每天初始化服务器列表的时间常量
- 连接池打开以后, 下一次才检查服务器列表是否需要更新
- update changelog

## [1.20.3] - 2024-01-24
### Changed
- 调整生成序列号函数
- 调整释放内存操作的位置
- 修订竞态数据的问题
- 调整生成序列号函数
- 更新依赖库版本
- update changelog

## [1.20.2] - 2024-01-23
### Changed
- 更新依赖库版本
- update changelog

## [1.20.1] - 2024-01-23
### Changed
- 更新依赖库版本
- update changelog

## [1.20.0] - 2024-01-23
### Changed
- 增加无法修复的异常日志
- 调整异常日志的输出
- update changelog

## [1.19.9] - 2024-01-22
### Changed
- 更新exchange版本号
- update changelog

## [1.19.8] - 2024-01-22
### Changed
- 优化连接池
- 调整缓存文件更新机制, 只是在重新bestip后才刷新本地缓存文件
- update changelog

## [1.19.7] - 2024-01-22
### Changed
- 更新exchange版本号
- update changelog

## [1.19.6] - 2024-01-22
### Changed
- 设置maxidle
- 调整BestIp的缓存更新方式
- 拆分init函数, 恢复BestIP函数
- 优化服务器地址同一时间只能调用一次的处理逻辑
- 更新gox版本
- !11 调整BestIp的缓存更新方式
Merge pull request !11 from heathen666/set_max_idle
- 更新依赖库版本
- 服务器列表先于默认9点整切换数据进行初始化服务器列表
- update changelog

## [1.19.5] - 2024-01-19
### Changed
- 统一计算价格小数点后的精度
- 优化快照部分代码
- 更新go.mod, 去掉本地调试配置
- update changelog

## [1.19.4] - 2024-01-19
### Changed
- 替换部分代码的立即数为常量
- 删除废弃的同步脚本
- 备注基本的交易单位函数
- 更新依赖库版本
- update changelog

## [1.19.3] - 2024-01-16
### Changed
- 优化行业板块的数据处理
- 删除调试中预留的控制台输出代码
- 更新依赖库版本
- update changelog

## [1.19.2] - 2024-01-13
### Changed
- 删除废弃的代码
- 更新依赖库版本
- update changelog

## [1.19.1] - 2024-01-11
### Changed
- 调整周期once组件
- 引入exchange组件, 移除交易日历和时间的处理功能到exchange
- 移除二级市场代码,改从exchange获取
- update changelog

## [1.19.0] - 2024-01-09
### Changed
- 更新依赖库, 优化服务器配置列表
- 调整滑动窗口式的once, 测试
- update changelog
- 更新依赖库gox版本

## [1.18.9] - 2024-01-03
### Changed
- 更新依赖库
- 精简部分代码
- 修订error信息秒数
- update changelog

## [1.18.8] - 2024-01-01
### Changed
- 修订服务列表模版的bug,标准行情和扩展行情都使用Std标签
- update changelog

## [1.18.7] - 2023-12-31
### Changed
- 修复CorrectSecurityCode没有正确处理空字符串参数
- !10 修复CorrectSecurityCode没有正确处理空字符串参数
Merge pull request !10 from heathen666/master
- update changelog

## [1.18.6] - 2023-12-31
### Changed
- 迁移fastjson从gox到pkg
- 调整服务器列表,新增通过模版直接生成地址列表的go源文件
- 调整地址模板, 去掉不符合go fmt格式的空白行
- update changelog

## [1.18.5] - 2023-12-30
### Changed
- 修复日历数组到尾部的bug
- update changelog

## [1.18.4] - 2023-12-30
### Changed
- 修复日历数组到尾部的bug
- update changelog

## [1.18.3] - 2023-12-30
### Changed
- 应广大股民要求内地三大交易所2024年除夕2月9日, 休市
- 测试金融节日数据接口
- 更新依赖库版本
- 更新依赖库版本
- 修订部分注释
- 调整两融列表逻辑, 如果从东方财富获取失败, 则从内置资源导出后加载
- update changelog

## [1.18.2] - 2023-12-25
### Changed
- 添加国泰君安服务器配置
- 添加华泰证券服务器配置
- 添加通达信服务器配置
- 添加中信服务器配置
- 打开读取配置文件的功能
- 调整服务器列表
- update changelog

## [1.18.1] - 2023-12-23
### Changed
- 新增返回服务器地址的数量
- 调整服务器列表
- update changelog

## [1.18.0] - 2023-12-23
### Changed
- 增加通达信最新的配置文件,备用.
- 新增东方财富两融数据页面地址, 备用.
- 更新依赖库版本
- 新增深证市场的交易日历链接地址
- 尝试调整日历的创建时间
- 从内置交易日历数据更新日历文件
- 本地日历文件损坏从内置日历数据加载
- update changelog

## [1.17.9] - 2023-12-17
### Changed
- 优化cache部分代码
- 更新依赖库gox版本
- 快照结构体增加本地时间戳字段
- update changelog

## [1.17.8] - 2023-12-14
### Changed
- 修订去重函数
- update changelog

## [1.17.7] - 2023-12-14
### Changed
- 更新依赖库gox版本
- update changelog

## [1.17.6] - 2023-12-12
### Changed
- 更新依赖库版本
- update changelog

## [1.17.5] - 2023-12-12
### Changed
- 北交所涨跌停板限制为30%
- update changelog

## [1.17.4] - 2023-12-11
### Changed
- 增加两融标的(2023-12-11)
- 增加获取两融标的列表的函数
- update changelog

## [1.17.3] - 2023-12-10
### Changed
- 调整交易日历的日期检测, 将文件的修改时间提前1年
- 更新依赖库版本号
- update changelog

## [1.17.2] - 2023-12-05
### Changed
- 更新gox依赖库版本
- 更新pkg, 将js vm路由到收录了goja的pkg(0.1.3)
- update changelog

## [1.17.1] - 2023-12-04
### Changed
- 更新http request的调用方法参数
- update changelog

## [1.17.0] - 2023-11-27
### Changed
- 修复A股尾盘竞价交易状态错误的bug
- update changelog

## [1.16.9] - 2023-11-13
### Changed
- 更新通达信行业板块配置文件
- update changelog

## [1.16.8] - 2023-11-07
### Changed
- 除权除息增加是否股本变化的判断方法
- update changelog

## [1.16.7] - 2023-10-31
### Changed
- 修复每日初始化一次日历的死锁的bug
- 更新依赖库版本
- update changelog

## [1.16.6] - 2023-10-29
### Changed
- 增加通过证券代码判断是否ETF
- 更新gox版本
- 修订TickTransaction的num字段, 在历史数据中该字段无效
- update changelog

## [1.16.5] - 2023-10-27
### Changed
- 快照测试增加退市的证券代码000666
- 调整滑动锁
- 更新gox版本
- update changelog

## [1.16.4] - 2023-10-26
### Changed
- 修复sz000666 经纬纺机主动退市引发的历史分时的bug, 数据校验不严谨
- update changelog

## [1.16.3] - 2023-10-26
### Changed
- 更新gox版本
- update changelog

## [1.16.2] - 2023-10-25
### Changed
- 调整尾盘交易结束到15:00:59.999
- update changelog

## [1.16.1] - 2023-10-23
### Changed
- 增加异常日志
- update changelog

## [1.16.0] - 2023-10-22
### Changed
- 更新gox版本
- update changelog

## [1.15.9] - 2023-10-19
### Changed
- 更新gox版本
- update changelog

## [1.15.8] - 2023-10-16
### Changed
- 更新gox版本
- 更新gox版本
- 调整周期性初始化组件
- update changelog

## [1.15.7] - 2023-10-11
### Changed
- 成交记录增加dataframe的tag
- 更新gox版本
- update changelog

## [1.15.6] - 2023-10-08
### Changed
- 更新gox版本
- update changelog

## [1.15.5] - 2023-10-07
### Changed
- 更新gox版本
- update changelog

## [1.15.4] - 2023-10-05
### Changed
- 调整缓存目录的初始化方法, 支持默认和定制两种方式
- 优化交易日历的初始化方式, 从init方法改为懒加载
- 修复分时历史数据股价不准确的bug
- update changelog

## [1.15.3] - 2023-10-02
### Changed
- 修复快照中ETF价格没有按照规定的3位小数点处理的bug
- 更新gox版本
- 调整util.MultiOne为coroutine.RollingMutex
- update changelog

## [1.15.2] - 2023-09-19
### Changed
- 增加尾盘时段

## [1.15.1] - 2023-09-15
### Changed
- 优化缓存路径的处理方式
- update changelog

## [1.15.0] - 2023-09-15
### Changed
- 增加时间段的分钟数计算函数

## [1.14.9] - 2023-09-14
### Changed
- 按照不同的操作系统拆分默认的缓存路径

## [1.14.8] - 2023-09-14
### Changed
- 删除废弃的代码

## [1.14.7] - 2023-09-14
### Changed
- 修订tcp连接关闭的保护方式

## [1.14.6] - 2023-09-14
### Changed
- 修复网络修复后的连接池hang的bug
- update changelog

## [1.14.5] - 2023-09-13
### Changed
- 修复网络修复后的连接池hang的bug
- update changelog

## [1.14.4] - 2023-09-13
### Changed
- 修复网络修复后的连接池hang的bug
- update changelog

## [1.14.3] - 2023-09-12
### Changed
- 更换golang.org/x/exp/slices为系统标准库
- update changelog

## [1.14.2] - 2023-09-12
### Changed
- 调整服务器IP池的锁机制
- update changelog

## [1.14.1] - 2023-09-10
### Changed
- 升级依赖库版本
- update changelog

## [1.14.0] - 2023-09-10
### Changed
- 调整单连接锁的用法
- update changelog

## [1.13.9] - 2023-09-07
### Changed
- 修订接口异常时终止心跳操作
- 心跳异常关闭连接
- update changelog

## [1.13.8] - 2023-08-28
### Changed
- update quotes/bestip_cache.go.

Signed-off-by: xubojam <xubojam@163.com>
- !9 update quotes/bestip_cache.go.
Merge pull request !9 from xubojam/N/A
- 更新依赖库
- update changelog

## [1.13.7] - 2023-08-14
### Changed
- !8 update README.md.
* update README.md.
- 快照增加计算平均竞价委托量
- update changelog

## [1.13.6] - 2023-08-13
### Changed
- 升级go版本到1.21.0
- update changelog

## [1.13.5] - 2023-08-01
### Changed
- update changelog
- 调整默认目录
- update changelog

## [1.13.4] - 2023-08-01
### Changed
- windows目录默认值c:/.quant1x

## [1.13.3] - 2023-07-25
### Changed
- 总成交量转换成股
- update changelog

## [1.13.2] - 2023-07-22
### Changed
- 修订集合竞价时段可刷新缓存文件的时间判断
- update changelog

## [1.13.1] - 2023-07-21
### Changed
- 更新gox库
- update changelog

## [1.13.0] - 2023-07-21
### Changed
- 增加竞价数据结束时间及判断函数
- update changelog

## [1.12.9] - 2023-07-08
### Changed
- 调整全部板块的缓存文件名
- update changelog

## [1.12.8] - 2023-07-08
### Changed
- 更新依赖库版本
- update changelog

## [1.12.7] - 2023-07-07
### Changed
- 更新依赖库
- update changelog

## [1.12.6] - 2023-07-07
### Changed
- 修复返回值存在超出日期范围的bug
- update changelog

## [1.12.5] - 2023-07-06
### Changed
- update changelog
- 删除独立的元数据包
- update changelog

## [1.12.4] - 2023-07-06
### Changed
- 调整证券信息的同步工具为MultiOne

## [1.12.3] - 2023-07-06
### Changed
- 调整证券类数据的package
- update changelog

## [1.12.2] - 2023-07-06
### Changed
- 板块数据归于元数据
- 交易日历归于元数据
- 上海代码新增510开头的ETF视为个股
- 增加510开头的ETF股价基数, 除以1000
- 增加基础的证券信息结构体
- 收敛日历文件的文件名函数
- 新增证券列表
- 收敛日历文件的文件名函数
- 测试周线前复权, 很遗憾, 接口上并没有发现这个参数, 只能自己实现了
- update changelog

## [1.12.1] - 2023-07-04
### Changed
- 修正板块数据不更新的bug
- update changelog

## [1.12.0] - 2023-07-02
### Changed
- 更新依赖库版本
- update changelog

## [1.11.9] - 2023-06-30
### Changed
- 清理部分服务器节点
- update changelog

## [1.11.8] - 2023-06-30
### Changed
- 修复死锁的bug
- update changelog

## [1.11.7] - 2023-06-30
### Changed
- 调整切换日期的数据重置逻辑
- update changelog

## [1.11.6] - 2023-06-29
### Changed
- 清理废弃的服务器节点
- 调整服务器节点
- update changelog

## [1.11.5] - 2023-06-29
### Changed
- 删除废弃的主站点
- update changelog

## [1.11.4] - 2023-06-29
### Changed
- 删除北京双线主站7节点
- update changelog

## [1.11.3] - 2023-06-27
### Changed
- 修复连接池重启的bug
- update changelog

## [1.11.2] - 2023-06-23
### Changed
- 修复除权除息信息流通股本字段的bug
- update changelog

## [1.11.1] - 2023-06-22
### Changed
- 调整交易时段状态
- update changelog

## [1.11.0] - 2023-06-22
### Changed
- 优化代码
- update changelog

## [1.10.9] - 2023-06-22
### Changed
- 增加板块类型
- 增加板块类型
- update changelog

## [1.10.8] - 2023-06-22
### Changed
- 调整package
- update changelog

## [1.10.7] - 2023-06-22
### Changed
- 优化export函数, 零拷贝
- update changelog

## [1.10.6] - 2023-06-22
### Changed
- 增加板块数据缓存
- update changelog

## [1.10.5] - 2023-06-20
### Changed
- 修订A股常量前缀为CN
- 修订A股常量前缀为CN
- update changelog

## [1.10.4] - 2023-06-20
### Changed
- 修订收盘量计算错误, 少乘了100
- update changelog

## [1.10.3] - 2023-06-20
### Changed
- 修订收盘量中指数板块与个股的不同
- update changelog

## [1.10.2] - 2023-06-19
### Changed
- 拆分开盘和收盘集合竞价函数
- update changelog

## [1.10.1] - 2023-06-19
### Changed
- 新增检查集合竞价时段的函数
- update changelog

## [1.10.0] - 2023-06-19
### Changed
- 快照数据增加交易状态、证券代码和收盘量计算
- update changelog
- 升级主版本到2.x.x
- 调整上证指数代码和sz000001区分
- update changelog
- 修订package路径
- update changelog

## [1.9.9] - 2023-06-19
### Changed
- 增加交易日午间休市的判断
- update changelog

## [1.9.8] - 2023-06-19
### Changed
- 调整CanUpdate这类函数支持默认参数, 即当前时间检查
- update changelog

## [1.9.7] - 2023-06-17
### Changed
- 显式暴露日历文件
- 更新依赖库版本
- update changelog

## [1.9.6] - 2023-06-17
### Changed
- update changelog
- 添加一个简易的客户端调用
- update changelog

## [1.9.5] - 2023-06-16
### Changed
- 删除废弃的字段

## [1.9.4] - 2023-06-16
### Changed
- 调整除权除息接口返回的浮点精度
- update changelog

## [1.9.3] - 2023-06-16
### Changed
- 更新依赖库
- update changelog

## [1.9.2] - 2023-06-16
### Changed
- 增加东方财富获取K线的函数
- 修订东方财富接口中时间戳的用法
- 快照数据增加日期
- update changelog

## [1.9.1] - 2023-06-16
### Changed
- 修订交易日、交易时间的约束
- update changelog

## [1.9.0] - 2023-06-16
### Changed
- 除权除息对象增加除权因子方法
- 快照数据补充当前交易日期
- 修订交易日、交易时间的约束
- update changelog

## [1.8.9] - 2023-06-16
### Changed
- 调整GetFinanceInfo函数的入参, 去掉num参数；FixTradeDate支持传入其它格式
- update changelog

## [1.8.8] - 2023-06-14
### Changed
- 修复交易日历功能里面的证券代码函数
- update changelog

## [1.8.7] - 2023-06-14
### Changed
- 新增修正证券代码函数
- update changelog

## [1.8.6] - 2023-06-14
### Changed
- 简化接口入参证券代码
- update changelog

## [1.8.5] - 2023-06-14
### Changed
- 梳理上市公司资料的方法代码
- update changelog

## [1.8.4] - 2023-06-14
### Changed
- update changelog
- 新增日历工具包
- 调整日历组件的依赖
- update changelog

## [1.8.3] - 2023-06-14
### Changed
- 新增涨停板函数

## [1.8.2] - 2023-06-13
### Changed
- 更新依赖库
- update changelog

## [1.8.1] - 2023-06-13
### Changed
- 快照增加日期字段
- update changelog

## [1.8.0] - 2023-06-13
### Changed
- 增加一个独立的快照方法
- update changelog

## [1.7.2] - 2023-06-08
### Changed
- 修订创业板的代码段为68xxxx
- update changelog

## [1.7.1] - 2023-06-06
### Changed
- 更新依赖库版本号
- update changelog

## [1.7.0] - 2023-05-15
### Changed
- 优化部分企业信息的代码
- 删除部分废弃的代码
- 调整测试代码
- 增加常量, 市场启动日期
- update changelog

## [1.6.32] - 2023-05-13
### Changed
- 更新依赖库版本号
- update changelog

## [1.6.31] - 2023-05-13
### Changed
- 迁移gox工具库
- update changelog

## [1.6.30] - 2023-05-13
### Changed
- !7 #I72I4X fixed: 屏蔽关闭连接时可能出现的panic
* 关闭TcpClient对象时忽略异常
- update changelog

## [1.6.29] - 2023-05-12
### Changed
- 更新依赖库版本号
- update changelog

## [1.6.28] - 2023-05-12
### Changed
- 更新依赖库版本号
- update changelog

## [1.6.27] - 2023-05-11
### Changed
- 修复连接池计算的bug
- update changelog

## [1.6.26] - 2023-05-11
### Changed
- 修改心跳操作为获取市场品种个数
- update changelog

## [1.6.25] - 2023-05-11
### Changed
- 增加读写超时时间到5秒
- update changelog

## [1.6.24] - 2023-05-11
### Changed
- 加强异常检测
- update changelog

## [1.6.23] - 2023-05-11
### Changed
- 加强异常检测
- update changelog

## [1.6.22] - 2023-05-11
### Changed
- 获取tcp连接异常输出日志
- update changelog

## [1.6.21] - 2023-05-11
### Changed
- 更新依赖库版本号
- update changelog

## [1.6.20] - 2023-05-11
### Changed
- 增加异常日志
- update changelog

## [1.6.19] - 2023-05-10
### Changed
- 更新依赖库版本号
- update changelog

## [1.6.18] - 2023-05-07
### Changed
- 调整部分代码
- update changelog

## [1.6.17] - 2023-05-07
### Changed
- 调整目录结构
- update changelog

## [1.6.16] - 2023-05-05
### Changed
- 增加测试代码
- 修复int转float64的bug
- update changelog

## [1.6.15] - 2023-05-05
### Changed
- 增加板块的涨跌情况统计数据
- update changelog

## [1.6.14] - 2023-05-05
### Changed
- 调整部分代码
- update changelog

## [1.6.13] - 2023-05-04
### Changed
- 收敛功能性函数
- update changelog

## [1.6.12] - 2023-05-04
### Changed
- !6 #I6ZWPX 新增F10企业基础信息接口
* 新增F10函数
* 调整财务数据的结构注释
- update changelog

## [1.6.11] - 2023-05-02
### Changed
- 财务数据判断是否退市
- update changelog

## [1.6.10] - 2023-05-01
### Changed
- update changelog
- 分笔成交数据的vol单位统一调整为股
- update changelog

## [1.6.9] - 2023-05-01
### Changed

## [1.6.8] - 2023-05-01
### Changed
- 增加判断个股的函数
- 调整服务器IP池
- 分笔成交数据增加类型3, 暂时不清楚其含义
- 调整分笔成交数据的部分代码
- update changelog
- !5 #I6ZU9B 统一分笔成交数据中的vol的单位, 调整为股
* 分笔成交数据的vol单位统一调整为股
- update changelog

## [1.6.7] - 2023-05-01
### Changed
- 调整代码
- 规范市场类型
- update changelog

## [1.6.6] - 2023-04-30
### Changed
- 优化部分代码
- update changelog

## [1.6.5] - 2023-04-30
### Changed
- 优化部分代码
- update changelog

## [1.6.4] - 2023-04-30
### Changed
- 修订分笔成交价格的计算方法
- 调整测试代码
- 优化IP池的检测
- update changelog

## [1.6.3] - 2023-04-29
### Changed
- 增加财务数据测试代码
- 增加即时行情测试代码
- 对齐退市状态
- update changelog

## [1.6.2] - 2023-04-28
### Changed
- 分笔成交数据增加常量
- update changelog

## [1.6.1] - 2023-04-27
### Changed
- 调整关闭心跳协程的方式
- update changelog

## [1.6.0] - 2023-04-26
### Changed
- 调整服务器列表的轮询问题
- update changelog

## [1.5.21] - 2023-04-26
### Changed
- 调整定时检测的锁方式
- 网络处理完成即更新时间戳
- update changelog

## [1.5.20] - 2023-04-26
### Changed
- 调整定时检测的锁方式
- update changelog

## [1.5.19] - 2023-04-26
### Changed
- 调整定时器的用法
- 定时任务退出时输出日志
- update changelog

## [1.5.18] - 2023-04-26
### Changed
- 心跳时间戳加锁
- update changelog

## [1.5.17] - 2023-04-26
### Changed
- 更新gox版本
- 优化best ip数据处理过程
- update changelog

## [1.5.16] - 2023-04-26
### Changed
- 调整time.Duration计算方法
- update changelog

## [1.5.15] - 2023-04-26
### Changed
- 调整心跳处理方式
- update changelog

## [1.5.14] - 2023-04-26
### Changed
- 调整心跳处理方式
- update changelog

## [1.5.13] - 2023-04-25
### Changed
- 去掉业务握手阶段的关闭连接池的操作
- 去掉业务握手阶段的关闭连接池的操作
- update changelog

## [1.5.12] - 2023-04-25
### Changed
- !4 #I6YKA4 调整快照数据
* 优化即时行情快照数据字段
* 优化即时行情快照数据字段
* 优化心跳处理机制
* 优化心跳处理机制
- update changelog

## [1.5.11] - 2023-04-24
### Changed
- update gox
- update changelog

## [1.5.10] - 2023-04-23
### Changed
- 更新gox工具版本
- update changelog

## [1.5.9] - 2023-04-23
### Changed
- 可用的服务器数量作为连接池最大数
- update changelog

## [1.5.8] - 2023-04-23
### Changed
- 优化连接池IP地址循环使用
- update changelog

## [1.5.7] - 2023-04-20
### Changed
- 统一指令入口
- update changelog

## [1.5.6] - 2023-04-20
### Changed
- 增加超时机制的测试代码
- 优化从ip池获取一个连接, 增加锁机制
- 增加心跳机制
- update changelog

## [1.5.5] - 2023-04-12
### Changed
- 优化代码
- 优化代码

## [1.5.4] - 2023-04-12
### Changed
- 增加注解
- 修正注释
- 去掉无用的代码
- 调整测试代码
- add CHANGELOG.md

## [1.5.3] - 2023-03-24
### Changed
- 取消todo项
- 忽略保留项
- 更新版本

## [1.5.2] - 2023-03-18
### Changed
- 删除部分注释
- 增加日志处理方式

## [1.5.1] - 2023-03-18
### Changed
- 测试新的行情数据结构, 不得要领，看不出未解密字段的含义
- 更新版本
- 更新gox版本
- 增加debug日志

## [1.5.0] - 2023-03-17
### Changed
- 优化常量的处理方式
- 增加心跳包

## [1.3.16] - 2023-03-17
### Changed
- 更改请求消息头字段名
- 更改响应消息头字段名

## [1.3.15] - 2023-03-16
### Changed
- 调整部分函数名为驼峰格式

## [1.3.14] - 2023-03-16
### Changed
- 拆分数字型转换float64的功能函数

## [1.3.13] - 2023-03-16
### Changed
- 调整部分函数名
- 修复zlib io.reader没有关闭

## [1.3.12] - 2023-03-15
### Changed
- 去掉部分输出控制台的代码

## [1.3.11] - 2023-03-15
### Changed
- 修正0x054c命令字结构, 暂时划归即时行情, 从新旧两种结构来看, 0x054c缺少2-5档数据, 增加了12个其它数据

## [1.3.10] - 2023-03-15
### Changed
- 增加新行情命令字

## [1.3.9] - 2023-03-15
### Changed
- 旧版本的行情数据
- 旧版本的行情数据

## [1.3.8] - 2023-03-15
### Changed
- 修订即时行情的命令字

## [1.3.7] - 2023-03-15
### Changed
- 恢复05

## [1.3.6] - 2023-03-15
### Changed
- 增加recv动作的超时时间
- 增加recv动作的超时时间
- 修订5档行情数据
- contentHex第一个字节如果是0x05, 获取的数据可能不及时, 会延迟几分钟

## [1.3.5] - 2023-03-13
### Changed
- 调整部分通达信系统批量数量限制的最大数类型

## [1.3.4] - 2023-03-13
### Changed
- 增加实时数据最大请求数据

## [1.3.3] - 2023-03-11
### Changed
- 恢复ping操作

## [1.3.2] - 2023-03-11
### Changed
- 修正部分告警信息

## [1.3.1] - 2023-03-11
### Changed
- 屏蔽ping代码, 直接返回

## [1.3.0] - 2023-03-11
### Changed
- 调整旧版本的包路径
- 调整旧版本的包路径
- 调整旧版本的包路径
- 调整超时时间为10秒
- 增加读取超时的判断
- 修正注释
- 增加延时的测试代码
- 删除废弃的测试代码
- 调整旧版本的包路径
- 调整命令字
- 精简代码
- 精简代码

## [1.2.8] - 2023-03-10
### Changed
- 88开头的代码为通达信板块指数, 从上海市场获取数据

## [1.2.7] - 2023-03-10
### Changed
- !3 #I6LKKR 新增板块接口
* 增加板块信息的测试代码
* 增加指数增加上涨和下跌家数
* 增加分笔成交的常量
* 增加K线的常量
* 增加股票列表的常量
* 增加block info数据接口
* 增加block meta数据接口
* 修订分时命令字
* 修订依赖库的版本号
* 修改文件名
* 增加注释
* 标准行情请求和响应header增加struc 表达式
* 计划接入板块数据

## [1.2.6] - 2023-03-03
### Changed
- !2 #I6J879 统一当日分笔成交和历史分笔成交的数据结构
* 统一分笔成交的接口

## [1.2.5] - 2023-02-27
### Changed
- 整理文档, 删除无用的代码

## [1.2.4] - 2023-02-27
### Changed
- !1 #I6I2J1 实现除权除息接口
* #I6I2J1 新增除权除息接口

## [1.2.3] - 2023-02-23
### Changed
- 升级gox版本

## [1.2.2] - 2023-02-21
### Changed
- 指数和个股的K线数据统一返回结构

## [1.2.1] - 2023-02-21
### Changed
- 屏蔽通过字符串解析服务时间的bug

## [1.2.0] - 2023-02-21
### Changed
- 调整仓库同步脚本
- 更新gox版本

## [1.1.9] - 2023-02-20
### Changed
- 增加退市提示信息
- 优化即时行情时间戳的整型处理方式

## [1.1.8] - 2023-02-20
### Changed
- 即时行情数据修订服务器时间

## [1.1.7] - 2023-02-20
### Changed
- 即时行情数据修订服务器时间

## [1.1.6] - 2023-02-20
### Changed
- 调整部分代码

## [1.1.5] - 2023-02-19
### Changed
- 关闭debug信息的输出
- 修正字段名

## [1.1.4] - 2023-02-18
### Changed
- 修正注释
- 修正go.mod

## [1.1.3] - 2023-02-18
### Changed
- 修订v1版本的demo
- 修订v1版本的demo
- 测试个股基本面信息, 可以确定的是可以取多个数据, 但是数据不完整, 具体问题还在分析
- 修正注释
- 增加市场代码
- 修正注释

## [1.1.2] - 2023-01-29
### Changed
- 修订README

## [1.1.1] - 2023-01-29
### Changed
- 调整通信接口入口函数名

## [1.1.0] - 2023-01-29
### Changed
- 修订gox版本, 增加gitee和github两个git代码仓库的同步脚本
- 修复类库名称错误
- 将前面实现的所有标准协议的接口定义v1
- 增加多个服务器寻轮检测

## [1.0.8] - 2023-01-27
### Changed
- 修订README

## [1.0.7] - 2023-01-27
### Changed
- 升级gox版本

## [1.0.6] - 2023-01-16
### Changed
- 通达信tcp协议连接调用之前再Hello2一次, 试验证明hello1就可以了

## [1.0.5] - 2023-01-16
### Changed
- 通达信tcp协议连接调用之前必须先Hello1一次
- 通达信tcp协议连接调用之前必须先Hello1一次

## [1.0.4] - 2023-01-16
### Changed
- add LICENSE.

Signed-off-by: 王布衣 <wangfengxy@sina.cn>

## [1.0.3] - 2023-01-16
### Changed
- 修订注释
- 修订注释
- 增加主机测试代码
- 更新gox到1.2.4, 利用lambda优化数组的处理
- 整合不同的协议处理方式的代码
- 增加2个新接口
- 增加2个新接口
- 增加2个新接口
- 增加4个接口
- 调整包路径
- 调整包路径
- 调整包路径
- 增加运行api初期测试主机速度

## [1.0.2] - 2023-01-15
### Changed
- 新增struc包
- 更新gox库, 从1.2.0升级到1.2.1
- 修订package对项目的变动
- 更新ASIO库版本
- 规范注释性资料
- 增加协议处理方式v1版本的个股基本面
- 删除项目内的c-struct package

## [1.0.1] - 2023-01-12
### Changed
- 修正常量

## [1.0.0] - 2023-01-12
### Changed
- first commit
- init
- get security quotes
- get index bar
- api
- readme
- 修正ioutil包
- 调整package
- 修订结构名
- 修订结构名
- 修订结构名
- 修订结构名
- 修订结构名
- 修订结构名
- 修订结构名
- 修订结构名
- 调整package名
- 修订结构名
- 测试当日分时数据
- 调整分时测试参数


[Unreleased]: https://gitee.com/quant1x/gotdx.git/compare/v1.25.0...HEAD
[1.25.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.24.0...v1.25.0
[1.24.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.15...v1.24.0
[1.23.15]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.14...v1.23.15
[1.23.14]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.13...v1.23.14
[1.23.13]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.12...v1.23.13
[1.23.12]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.11...v1.23.12
[1.23.11]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.10...v1.23.11
[1.23.10]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.9...v1.23.10
[1.23.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.8...v1.23.9
[1.23.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.7...v1.23.8
[1.23.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.6...v1.23.7
[1.23.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.5...v1.23.6
[1.23.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.4...v1.23.5
[1.23.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.3...v1.23.4
[1.23.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.2...v1.23.3
[1.23.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.1...v1.23.2
[1.23.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.23.0...v1.23.1
[1.23.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.23...v1.23.0
[1.22.23]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.22...v1.22.23
[1.22.22]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.21...v1.22.22
[1.22.21]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.20...v1.22.21
[1.22.20]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.19...v1.22.20
[1.22.19]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.18...v1.22.19
[1.22.18]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.17...v1.22.18
[1.22.17]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.16...v1.22.17
[1.22.16]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.15...v1.22.16
[1.22.15]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.14...v1.22.15
[1.22.14]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.13...v1.22.14
[1.22.13]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.12...v1.22.13
[1.22.12]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.11...v1.22.12
[1.22.11]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.10...v1.22.11
[1.22.10]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.9...v1.22.10
[1.22.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.8...v1.22.9
[1.22.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.7...v1.22.8
[1.22.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.6...v1.22.7
[1.22.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.5...v1.22.6
[1.22.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.4...v1.22.5
[1.22.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.3...v1.22.4
[1.22.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.2...v1.22.3
[1.22.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.1...v1.22.2
[1.22.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.22.0...v1.22.1
[1.22.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.9...v1.22.0
[1.21.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.8...v1.21.9
[1.21.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.7...v1.21.8
[1.21.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.6...v1.21.7
[1.21.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.5...v1.21.6
[1.21.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.4...v1.21.5
[1.21.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.3...v1.21.4
[1.21.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.2...v1.21.3
[1.21.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.1...v1.21.2
[1.21.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.21.0...v1.21.1
[1.21.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.9...v1.21.0
[1.20.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.8...v1.20.9
[1.20.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.7...v1.20.8
[1.20.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.6...v1.20.7
[1.20.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.5...v1.20.6
[1.20.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.4...v1.20.5
[1.20.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.3...v1.20.4
[1.20.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.2...v1.20.3
[1.20.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.1...v1.20.2
[1.20.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.20.0...v1.20.1
[1.20.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.9...v1.20.0
[1.19.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.8...v1.19.9
[1.19.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.7...v1.19.8
[1.19.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.6...v1.19.7
[1.19.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.5...v1.19.6
[1.19.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.4...v1.19.5
[1.19.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.3...v1.19.4
[1.19.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.2...v1.19.3
[1.19.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.1...v1.19.2
[1.19.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.19.0...v1.19.1
[1.19.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.9...v1.19.0
[1.18.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.8...v1.18.9
[1.18.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.7...v1.18.8
[1.18.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.6...v1.18.7
[1.18.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.5...v1.18.6
[1.18.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.4...v1.18.5
[1.18.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.3...v1.18.4
[1.18.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.2...v1.18.3
[1.18.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.1...v1.18.2
[1.18.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.18.0...v1.18.1
[1.18.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.9...v1.18.0
[1.17.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.8...v1.17.9
[1.17.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.7...v1.17.8
[1.17.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.6...v1.17.7
[1.17.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.5...v1.17.6
[1.17.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.4...v1.17.5
[1.17.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.3...v1.17.4
[1.17.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.2...v1.17.3
[1.17.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.1...v1.17.2
[1.17.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.17.0...v1.17.1
[1.17.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.9...v1.17.0
[1.16.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.8...v1.16.9
[1.16.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.7...v1.16.8
[1.16.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.6...v1.16.7
[1.16.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.5...v1.16.6
[1.16.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.4...v1.16.5
[1.16.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.3...v1.16.4
[1.16.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.2...v1.16.3
[1.16.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.1...v1.16.2
[1.16.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.16.0...v1.16.1
[1.16.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.9...v1.16.0
[1.15.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.8...v1.15.9
[1.15.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.7...v1.15.8
[1.15.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.6...v1.15.7
[1.15.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.5...v1.15.6
[1.15.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.4...v1.15.5
[1.15.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.3...v1.15.4
[1.15.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.2...v1.15.3
[1.15.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.1...v1.15.2
[1.15.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.15.0...v1.15.1
[1.15.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.9...v1.15.0
[1.14.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.8...v1.14.9
[1.14.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.7...v1.14.8
[1.14.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.6...v1.14.7
[1.14.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.5...v1.14.6
[1.14.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.4...v1.14.5
[1.14.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.3...v1.14.4
[1.14.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.2...v1.14.3
[1.14.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.1...v1.14.2
[1.14.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.14.0...v1.14.1
[1.14.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.9...v1.14.0
[1.13.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.8...v1.13.9
[1.13.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.7...v1.13.8
[1.13.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.6...v1.13.7
[1.13.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.5...v1.13.6
[1.13.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.4...v1.13.5
[1.13.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.3...v1.13.4
[1.13.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.2...v1.13.3
[1.13.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.1...v1.13.2
[1.13.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.13.0...v1.13.1
[1.13.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.9...v1.13.0
[1.12.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.8...v1.12.9
[1.12.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.7...v1.12.8
[1.12.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.6...v1.12.7
[1.12.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.5...v1.12.6
[1.12.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.4...v1.12.5
[1.12.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.3...v1.12.4
[1.12.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.2...v1.12.3
[1.12.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.1...v1.12.2
[1.12.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.12.0...v1.12.1
[1.12.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.9...v1.12.0
[1.11.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.8...v1.11.9
[1.11.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.7...v1.11.8
[1.11.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.6...v1.11.7
[1.11.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.5...v1.11.6
[1.11.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.4...v1.11.5
[1.11.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.3...v1.11.4
[1.11.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.2...v1.11.3
[1.11.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.1...v1.11.2
[1.11.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.11.0...v1.11.1
[1.11.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.9...v1.11.0
[1.10.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.8...v1.10.9
[1.10.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.7...v1.10.8
[1.10.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.6...v1.10.7
[1.10.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.5...v1.10.6
[1.10.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.4...v1.10.5
[1.10.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.3...v1.10.4
[1.10.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.2...v1.10.3
[1.10.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.1...v1.10.2
[1.10.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.10.0...v1.10.1
[1.10.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.9...v1.10.0
[1.9.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.8...v1.9.9
[1.9.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.7...v1.9.8
[1.9.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.6...v1.9.7
[1.9.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.5...v1.9.6
[1.9.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.4...v1.9.5
[1.9.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.3...v1.9.4
[1.9.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.2...v1.9.3
[1.9.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.1...v1.9.2
[1.9.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.9.0...v1.9.1
[1.9.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.9...v1.9.0
[1.8.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.8...v1.8.9
[1.8.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.7...v1.8.8
[1.8.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.6...v1.8.7
[1.8.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.5...v1.8.6
[1.8.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.4...v1.8.5
[1.8.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.3...v1.8.4
[1.8.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.2...v1.8.3
[1.8.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.1...v1.8.2
[1.8.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.8.0...v1.8.1
[1.8.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.7.2...v1.8.0
[1.7.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.7.1...v1.7.2
[1.7.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.7.0...v1.7.1
[1.7.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.32...v1.7.0
[1.6.32]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.31...v1.6.32
[1.6.31]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.30...v1.6.31
[1.6.30]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.29...v1.6.30
[1.6.29]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.28...v1.6.29
[1.6.28]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.27...v1.6.28
[1.6.27]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.26...v1.6.27
[1.6.26]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.25...v1.6.26
[1.6.25]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.24...v1.6.25
[1.6.24]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.23...v1.6.24
[1.6.23]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.22...v1.6.23
[1.6.22]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.21...v1.6.22
[1.6.21]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.20...v1.6.21
[1.6.20]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.19...v1.6.20
[1.6.19]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.18...v1.6.19
[1.6.18]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.17...v1.6.18
[1.6.17]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.16...v1.6.17
[1.6.16]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.15...v1.6.16
[1.6.15]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.14...v1.6.15
[1.6.14]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.13...v1.6.14
[1.6.13]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.12...v1.6.13
[1.6.12]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.11...v1.6.12
[1.6.11]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.10...v1.6.11
[1.6.10]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.9...v1.6.10
[1.6.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.8...v1.6.9
[1.6.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.7...v1.6.8
[1.6.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.6...v1.6.7
[1.6.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.5...v1.6.6
[1.6.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.4...v1.6.5
[1.6.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.3...v1.6.4
[1.6.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.2...v1.6.3
[1.6.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.1...v1.6.2
[1.6.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.6.0...v1.6.1
[1.6.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.21...v1.6.0
[1.5.21]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.20...v1.5.21
[1.5.20]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.19...v1.5.20
[1.5.19]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.18...v1.5.19
[1.5.18]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.17...v1.5.18
[1.5.17]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.16...v1.5.17
[1.5.16]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.15...v1.5.16
[1.5.15]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.14...v1.5.15
[1.5.14]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.13...v1.5.14
[1.5.13]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.12...v1.5.13
[1.5.12]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.11...v1.5.12
[1.5.11]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.10...v1.5.11
[1.5.10]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.9...v1.5.10
[1.5.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.8...v1.5.9
[1.5.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.7...v1.5.8
[1.5.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.6...v1.5.7
[1.5.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.5...v1.5.6
[1.5.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.4...v1.5.5
[1.5.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.3...v1.5.4
[1.5.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.2...v1.5.3
[1.5.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.1...v1.5.2
[1.5.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.5.0...v1.5.1
[1.5.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.16...v1.5.0
[1.3.16]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.15...v1.3.16
[1.3.15]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.14...v1.3.15
[1.3.14]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.13...v1.3.14
[1.3.13]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.12...v1.3.13
[1.3.12]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.11...v1.3.12
[1.3.11]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.10...v1.3.11
[1.3.10]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.9...v1.3.10
[1.3.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.8...v1.3.9
[1.3.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.7...v1.3.8
[1.3.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.6...v1.3.7
[1.3.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.5...v1.3.6
[1.3.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.4...v1.3.5
[1.3.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.3...v1.3.4
[1.3.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.2...v1.3.3
[1.3.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.1...v1.3.2
[1.3.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.3.0...v1.3.1
[1.3.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.8...v1.3.0
[1.2.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.7...v1.2.8
[1.2.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.6...v1.2.7
[1.2.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.5...v1.2.6
[1.2.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.4...v1.2.5
[1.2.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.3...v1.2.4
[1.2.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.2...v1.2.3
[1.2.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.1...v1.2.2
[1.2.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.9...v1.2.0
[1.1.9]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.8...v1.1.9
[1.1.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.7...v1.1.8
[1.1.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.6...v1.1.7
[1.1.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.5...v1.1.6
[1.1.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.4...v1.1.5
[1.1.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.3...v1.1.4
[1.1.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.2...v1.1.3
[1.1.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.1...v1.1.2
[1.1.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.1.0...v1.1.1
[1.1.0]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.8...v1.1.0
[1.0.8]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.7...v1.0.8
[1.0.7]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.6...v1.0.7
[1.0.6]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.5...v1.0.6
[1.0.5]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.4...v1.0.5
[1.0.4]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.3...v1.0.4
[1.0.3]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.2...v1.0.3
[1.0.2]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.1...v1.0.2
[1.0.1]: https://gitee.com/quant1x/gotdx.git/compare/v1.0.0...v1.0.1

[1.0.0]: https://gitee.com/quant1x/gotdx.git/releases/tag/v1.0.0
