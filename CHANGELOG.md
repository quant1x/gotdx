# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.6.8] - 2023-05-01
### Changed
- 调整分笔成交数据的部分代码.
- 分笔成交数据增加类型3, 暂时不清楚其含义.
- 调整服务器IP池.
- 增加判断个股的函数.

## [1.6.7] - 2023-05-01
### Changed
- 规范市场类型.
- 调整代码.

## [1.6.6] - 2023-04-30
### Changed
- 优化部分代码.

## [1.6.5] - 2023-04-30
### Changed
- 优化部分代码.

## [1.6.4] - 2023-04-30
### Changed
- 优化IP池的检测.
- 调整测试代码.
- 修订分笔成交价格的计算方法.

## [1.6.3] - 2023-04-29
### Changed
- 对齐退市状态.
- 增加即时行情测试代码.
- 增加财务数据测试代码.

## [1.6.2] - 2023-04-28
### Changed
- 分笔成交数据增加常量.

## [1.6.1] - 2023-04-27
### Changed
- 调整关闭心跳协程的方式.

## [1.6.0] - 2023-04-26
### Changed
- 调整服务器列表的轮询问题.

## [1.5.21] - 2023-04-26
### Changed
- 网络处理完成即更新时间戳.
- 调整定时检测的锁方式.

## [1.5.20] - 2023-04-26
### Changed
- 调整定时检测的锁方式.

## [1.5.19] - 2023-04-26
### Changed
- 定时任务退出时输出日志.
- 调整定时器的用法.

## [1.5.18] - 2023-04-26
### Changed
- 心跳时间戳加锁.

## [1.5.17] - 2023-04-26
### Changed
- 优化best ip数据处理过程.
- 更新gox版本.

## [1.5.16] - 2023-04-26
### Changed
- 调整time.Duration计算方法.

## [1.5.15] - 2023-04-26
### Changed
- 调整心跳处理方式.

## [1.5.14] - 2023-04-26
### Changed
- 调整心跳处理方式.

## [1.5.13] - 2023-04-25
### Changed
- 去掉业务握手阶段的关闭连接池的操作.
- 去掉业务握手阶段的关闭连接池的操作.

## [1.5.12] - 2023-04-25
### Changed
- !4 #I6YKA4 调整快照数据 * 优化即时行情快照数据字段 * 优化即时行情快照数据字段 * 优化心跳处理机制 * 优化心跳处理机制.

## [1.5.11] - 2023-04-24
### Changed
- Update gox.

## [1.5.10] - 2023-04-23
### Changed
- 更新gox工具版本.

## [1.5.9] - 2023-04-23
### Changed
- 可用的服务器数量作为连接池最大数.

## [1.5.8] - 2023-04-23
### Changed
- 优化连接池IP地址循环使用.

## [1.5.7] - 2023-04-20
### Changed
- 统一指令入口.

## [1.5.6] - 2023-04-20
### Changed
- 增加心跳机制.
- 优化从ip池获取一个连接, 增加锁机制.
- 增加超时机制的测试代码.

## [1.5.5] - 2023-04-12
### Changed
- 优化代码.

## [1.5.4] - 2023-04-12
### Changed
- Add CHANGELOG.md.
- 调整测试代码.
- 去掉无用的代码.
- 修正注释.
- 增加注解.

## [1.5.3] - 2023-03-24
### Changed
- 更新版本.
- 忽略保留项.
- 取消todo项.

## [1.5.2] - 2023-03-18
### Changed
- 增加日志处理方式.
- 删除部分注释.

## [1.5.1] - 2023-03-18
### Changed
- 增加debug日志.
- 更新gox版本.
- 更新版本.
- 测试新的行情数据结构, 不得要领，看不出未解密字段的含义.

## [1.5.0] - 2023-03-17
### Changed
- 增加心跳包.
- 优化常量的处理方式.

## [1.3.16] - 2023-03-17
### Changed
- 更改响应消息头字段名.
- 更改请求消息头字段名.

## [1.3.15] - 2023-03-16
### Changed
- 调整部分函数名为驼峰格式.

## [1.3.14] - 2023-03-16
### Changed
- 拆分数字型转换float64的功能函数.

## [1.3.13] - 2023-03-16
### Changed
- 修复zlib io.reader没有关闭.
- 调整部分函数名.

## [1.3.12] - 2023-03-15
### Changed
- 去掉部分输出控制台的代码.

## [1.3.11] - 2023-03-15
### Changed
- 修正0x054c命令字结构, 暂时划归即时行情, 从新旧两种结构来看, 0x054c缺少2-5档数据, 增加了12个其它数据.

## [1.3.10] - 2023-03-15
### Changed
- 增加新行情命令字.

## [1.3.9] - 2023-03-15
### Changed
- 旧版本的行情数据.
- 旧版本的行情数据.

## [1.3.8] - 2023-03-15
### Changed
- 修订即时行情的命令字.

## [1.3.7] - 2023-03-15
### Changed
- 恢复05.

## [1.3.6] - 2023-03-15
### Changed
- ContentHex第一个字节如果是0x05, 获取的数据可能不及时, 会延迟几分钟.
- 修订5档行情数据.
- 增加recv动作的超时时间.
- 增加recv动作的超时时间.

## [1.3.5] - 2023-03-13
### Changed
- 调整部分通达信系统批量数量限制的最大数类型.

## [1.3.4] - 2023-03-13
### Changed
- 增加实时数据最大请求数据.

## [1.3.3] - 2023-03-11
### Changed
- 恢复ping操作.

## [1.3.2] - 2023-03-11
### Changed
- 修正部分告警信息.

## [1.3.1] - 2023-03-11
### Changed
- 屏蔽ping代码, 直接返回.

## [1.3.0] - 2023-03-11
### Changed
- 精简代码.
- 精简代码.
- 调整命令字.
- 调整旧版本的包路径.
- 删除废弃的测试代码.
- 增加延时的测试代码.
- 修正注释.
- 增加读取超时的判断.
- 调整超时时间为10秒.
- 调整旧版本的包路径.
- 调整旧版本的包路径.
- 调整旧版本的包路径.

## [1.2.8] - 2023-03-10
### Changed
- 88开头的代码为通达信板块指数, 从上海市场获取数据.

## [1.2.7] - 2023-03-10
### Changed
- !3 #I6LKKR 新增板块接口 * 增加板块信息的测试代码 * 增加指数增加上涨和下跌家数 * 增加分笔成交的常量 * 增加K线的常量 * 增加股票列表的常量 * 增加block info数据接口 * 增加block meta数据接口 * 修订分时命令字 * 修订依赖库的版本号 * 修改文件名 * 增加注释 * 标准行情请求和响应header增加struc 表达式 * 计划接入板块数据.

## [1.2.6] - 2023-03-03
### Changed
- !2 #I6J879 统一当日分笔成交和历史分笔成交的数据结构 * 统一分笔成交的接口.

## [1.2.5] - 2023-02-27
### Changed
- 整理文档, 删除无用的代码.

## [1.2.4] - 2023-02-27
### Changed
- !1 #I6I2J1 实现除权除息接口 * #I6I2J1 新增除权除息接口.

## [1.2.3] - 2023-02-23
### Changed
- 升级gox版本.

## [1.2.2] - 2023-02-21
### Changed
- 指数和个股的K线数据统一返回结构.

## [1.2.1] - 2023-02-21
### Changed
- 屏蔽通过字符串解析服务时间的bug.

## [1.2.0] - 2023-02-21
### Changed
- 更新gox版本.
- 调整仓库同步脚本.

## [1.1.9] - 2023-02-20
### Changed
- 优化即时行情时间戳的整型处理方式.
- 增加退市提示信息.

## [1.1.8] - 2023-02-20
### Changed
- 即时行情数据修订服务器时间.

## [1.1.7] - 2023-02-20
### Changed
- 即时行情数据修订服务器时间.

## [1.1.6] - 2023-02-20
### Changed
- 调整部分代码.

## [1.1.5] - 2023-02-19
### Changed
- 修正字段名.
- 关闭debug信息的输出.

## [1.1.4] - 2023-02-18
### Changed
- 修正go.mod.
- 修正注释.

## [1.1.3] - 2023-02-18
### Changed
- 修正注释.
- 增加市场代码.
- 修正注释.
- 测试个股基本面信息, 可以确定的是可以取多个数据, 但是数据不完整, 具体问题还在分析.
- 修订v1版本的demo.
- 修订v1版本的demo.

## [1.1.2] - 2023-01-29
### Changed
- 修订README.

## [1.1.1] - 2023-01-29
### Changed
- 调整通信接口入口函数名.

## [1.1.0] - 2023-01-29
### Changed
- 增加多个服务器寻轮检测.
- 将前面实现的所有标准协议的接口定义v1.
- 修复类库名称错误.
- 修订gox版本, 增加gitee和github两个git代码仓库的同步脚本.

## [1.0.8] - 2023-01-27
### Changed
- 修订README.

## [1.0.7] - 2023-01-27
### Changed
- 升级gox版本.

## [1.0.6] - 2023-01-16
### Changed
- 通达信tcp协议连接调用之前再Hello2一次, 试验证明hello1就可以了.

## [1.0.5] - 2023-01-16
### Changed
- 通达信tcp协议连接调用之前必须先Hello1一次.
- 通达信tcp协议连接调用之前必须先Hello1一次.

## [1.0.4] - 2023-01-16
### Changed
- Add LICENSE.

## [1.0.3] - 2023-01-16
### Changed
- 增加运行api初期测试主机速度.
- 调整包路径.
- 调整包路径.
- 调整包路径.
- 增加4个接口.
- 增加2个新接口.
- 增加2个新接口.
- 增加2个新接口.
- 整合不同的协议处理方式的代码.
- 更新gox到1.2.4, 利用lambda优化数组的处理.
- 增加主机测试代码.
- 修订注释.
- 修订注释.

## [1.0.2] - 2023-01-15
### Changed
- 删除项目内的c-struct package.
- 增加协议处理方式v1版本的个股基本面.
- 规范注释性资料.
- 更新ASIO库版本.
- 修订package对项目的变动.
- 更新gox库, 从1.2.0升级到1.2.1.
- 新增struc包.

## [1.0.1] - 2023-01-12
### Changed
- 修正常量.

## [1.0.0] - 2023-01-12
### Changed
- 调整分时测试参数.
- 测试当日分时数据.
- 修订结构名.
- 调整package名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 修订结构名.
- 调整package.
- 修正ioutil包.
- Readme.
- Api.
- Get index bar.
- Get security quotes.
- Init.
- First commit.

[Unreleased]: https://gitee.com/quant1x/gotdx/compare/v1.6.8...HEAD
[1.6.8]: https://gitee.com/quant1x/gotdx/compare/v1.6.7...v1.6.8
[1.6.7]: https://gitee.com/quant1x/gotdx/compare/v1.6.6...v1.6.7
[1.6.6]: https://gitee.com/quant1x/gotdx/compare/v1.6.5...v1.6.6
[1.6.5]: https://gitee.com/quant1x/gotdx/compare/v1.6.4...v1.6.5
[1.6.4]: https://gitee.com/quant1x/gotdx/compare/v1.6.3...v1.6.4
[1.6.3]: https://gitee.com/quant1x/gotdx/compare/v1.6.2...v1.6.3
[1.6.2]: https://gitee.com/quant1x/gotdx/compare/v1.6.1...v1.6.2
[1.6.1]: https://gitee.com/quant1x/gotdx/compare/v1.6.0...v1.6.1
[1.6.0]: https://gitee.com/quant1x/gotdx/compare/v1.5.21...v1.6.0
[1.5.21]: https://gitee.com/quant1x/gotdx/compare/v1.5.20...v1.5.21
[1.5.20]: https://gitee.com/quant1x/gotdx/compare/v1.5.19...v1.5.20
[1.5.19]: https://gitee.com/quant1x/gotdx/compare/v1.5.18...v1.5.19
[1.5.18]: https://gitee.com/quant1x/gotdx/compare/v1.5.17...v1.5.18
[1.5.17]: https://gitee.com/quant1x/gotdx/compare/v1.5.16...v1.5.17
[1.5.16]: https://gitee.com/quant1x/gotdx/compare/v1.5.15...v1.5.16
[1.5.15]: https://gitee.com/quant1x/gotdx/compare/v1.5.14...v1.5.15
[1.5.14]: https://gitee.com/quant1x/gotdx/compare/v1.5.13...v1.5.14
[1.5.13]: https://gitee.com/quant1x/gotdx/compare/v1.5.12...v1.5.13
[1.5.12]: https://gitee.com/quant1x/gotdx/compare/v1.5.11...v1.5.12
[1.5.11]: https://gitee.com/quant1x/gotdx/compare/v1.5.10...v1.5.11
[1.5.10]: https://gitee.com/quant1x/gotdx/compare/v1.5.9...v1.5.10
[1.5.9]: https://gitee.com/quant1x/gotdx/compare/v1.5.8...v1.5.9
[1.5.8]: https://gitee.com/quant1x/gotdx/compare/v1.5.7...v1.5.8
[1.5.7]: https://gitee.com/quant1x/gotdx/compare/v1.5.6...v1.5.7
[1.5.6]: https://gitee.com/quant1x/gotdx/compare/v1.5.5...v1.5.6
[1.5.5]: https://gitee.com/quant1x/gotdx/compare/v1.5.4...v1.5.5
[1.5.4]: https://gitee.com/quant1x/gotdx/compare/v1.5.3...v1.5.4
[1.5.3]: https://gitee.com/quant1x/gotdx/compare/v1.5.2...v1.5.3
[1.5.2]: https://gitee.com/quant1x/gotdx/compare/v1.5.1...v1.5.2
[1.5.1]: https://gitee.com/quant1x/gotdx/compare/v1.5.0...v1.5.1
[1.5.0]: https://gitee.com/quant1x/gotdx/compare/v1.3.16...v1.5.0
[1.3.16]: https://gitee.com/quant1x/gotdx/compare/v1.3.15...v1.3.16
[1.3.15]: https://gitee.com/quant1x/gotdx/compare/v1.3.14...v1.3.15
[1.3.14]: https://gitee.com/quant1x/gotdx/compare/v1.3.13...v1.3.14
[1.3.13]: https://gitee.com/quant1x/gotdx/compare/v1.3.12...v1.3.13
[1.3.12]: https://gitee.com/quant1x/gotdx/compare/v1.3.11...v1.3.12
[1.3.11]: https://gitee.com/quant1x/gotdx/compare/v1.3.10...v1.3.11
[1.3.10]: https://gitee.com/quant1x/gotdx/compare/v1.3.9...v1.3.10
[1.3.9]: https://gitee.com/quant1x/gotdx/compare/v1.3.8...v1.3.9
[1.3.8]: https://gitee.com/quant1x/gotdx/compare/v1.3.7...v1.3.8
[1.3.7]: https://gitee.com/quant1x/gotdx/compare/v1.3.6...v1.3.7
[1.3.6]: https://gitee.com/quant1x/gotdx/compare/v1.3.5...v1.3.6
[1.3.5]: https://gitee.com/quant1x/gotdx/compare/v1.3.4...v1.3.5
[1.3.4]: https://gitee.com/quant1x/gotdx/compare/v1.3.3...v1.3.4
[1.3.3]: https://gitee.com/quant1x/gotdx/compare/v1.3.2...v1.3.3
[1.3.2]: https://gitee.com/quant1x/gotdx/compare/v1.3.1...v1.3.2
[1.3.1]: https://gitee.com/quant1x/gotdx/compare/v1.3.0...v1.3.1
[1.3.0]: https://gitee.com/quant1x/gotdx/compare/v1.2.8...v1.3.0
[1.2.8]: https://gitee.com/quant1x/gotdx/compare/v1.2.7...v1.2.8
[1.2.7]: https://gitee.com/quant1x/gotdx/compare/v1.2.6...v1.2.7
[1.2.6]: https://gitee.com/quant1x/gotdx/compare/v1.2.5...v1.2.6
[1.2.5]: https://gitee.com/quant1x/gotdx/compare/v1.2.4...v1.2.5
[1.2.4]: https://gitee.com/quant1x/gotdx/compare/v1.2.3...v1.2.4
[1.2.3]: https://gitee.com/quant1x/gotdx/compare/v1.2.2...v1.2.3
[1.2.2]: https://gitee.com/quant1x/gotdx/compare/v1.2.1...v1.2.2
[1.2.1]: https://gitee.com/quant1x/gotdx/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitee.com/quant1x/gotdx/compare/v1.1.9...v1.2.0
[1.1.9]: https://gitee.com/quant1x/gotdx/compare/v1.1.8...v1.1.9
[1.1.8]: https://gitee.com/quant1x/gotdx/compare/v1.1.7...v1.1.8
[1.1.7]: https://gitee.com/quant1x/gotdx/compare/v1.1.6...v1.1.7
[1.1.6]: https://gitee.com/quant1x/gotdx/compare/v1.1.5...v1.1.6
[1.1.5]: https://gitee.com/quant1x/gotdx/compare/v1.1.4...v1.1.5
[1.1.4]: https://gitee.com/quant1x/gotdx/compare/v1.1.3...v1.1.4
[1.1.3]: https://gitee.com/quant1x/gotdx/compare/v1.1.2...v1.1.3
[1.1.2]: https://gitee.com/quant1x/gotdx/compare/v1.1.1...v1.1.2
[1.1.1]: https://gitee.com/quant1x/gotdx/compare/v1.1.0...v1.1.1
[1.1.0]: https://gitee.com/quant1x/gotdx/compare/v1.0.8...v1.1.0
[1.0.8]: https://gitee.com/quant1x/gotdx/compare/v1.0.7...v1.0.8
[1.0.7]: https://gitee.com/quant1x/gotdx/compare/v1.0.6...v1.0.7
[1.0.6]: https://gitee.com/quant1x/gotdx/compare/v1.0.5...v1.0.6
[1.0.5]: https://gitee.com/quant1x/gotdx/compare/v1.0.4...v1.0.5
[1.0.4]: https://gitee.com/quant1x/gotdx/compare/v1.0.3...v1.0.4
[1.0.3]: https://gitee.com/quant1x/gotdx/compare/v1.0.2...v1.0.3
[1.0.2]: https://gitee.com/quant1x/gotdx/compare/v1.0.1...v1.0.2
[1.0.1]: https://gitee.com/quant1x/gotdx/compare/v1.0.0...v1.0.1
[1.0.0]: https://gitee.com/quant1x/gotdx/releases/tag/v1.0.0
