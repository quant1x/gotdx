package quotes

type StdApi struct {
	connPool *ConnPool
}

func NewStdApi(srv Server) *StdApi {
	size := 1
	opt := Opt{
		Host: srv.Host,
		Port: srv.Port,
	}
	_factory := func(string) (interface{}, error) {
		client := NewClient(&opt)
		err := client.Connect()
		return client, err
	}
	_close := func(v interface{}) error {
		client := v.(*TcpClient)
		return client.Close()
	}
	_ping := func(v interface{}) error {
		return nil
	}
	cp := NewConnPool2(srv.Host, srv.Port, size, _factory, _close, _ping)
	api := StdApi{
		connPool: cp,
	}
	return &api
}

// Close 关闭
func (this *StdApi) Close() {
	this.connPool.Close()
}

func (this *StdApi) command(msg Message) (interface{}, error) {

	// 2.1 获取TCP连接
	_conn := this.connPool.GetConn()
	cli := _conn.(*TcpClient)
	opt := cli.GetOpt()
	conn := cli.GetConn()
	err := process(conn, msg, opt)
	if err != nil {
		_ = cli.Close()
		return nil, err
	}

	this.connPool.ReturnConn(_conn)
	return msg.Reply(), nil
}

func (this *StdApi) Hello1() (*Hello1Reply, error) {
	// 创建一个hello1消息
	hello1 := NewHello1()
	reply, err := this.command(hello1)
	if err != nil {
		return nil, err
	}
	return reply.(*Hello1Reply), nil
}

func (this *StdApi) Hello2() (*Hello2Reply, error) {
	// 创建一个hello1消息
	hello2 := NewHello2()

	reply, err := this.command(hello2)
	if err != nil {
		return nil, err
	}
	return reply.(*Hello2Reply), nil
}

// GetFinanceInfo 基本面
func (this *StdApi) GetFinanceInfo(market int, code string) (*FinanceInfo, error) {
	msg := NewFinanceInfoPackage()
	_code := [6]byte{}
	_market := uint8(market)
	copy(_code[:], code)
	msg.SetParams(&FinanceInfoRequest{
		Market: _market,
		Code:   _code,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*FinanceInfo), nil
}

// GetKLine K线
func (this *StdApi) GetKLine(market int, code string, category uint16, start uint16, count uint16) (*SecurityBarsReply, error) {
	msg := NewSecurityBarsPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	msg.SetParams(&SecurityBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*SecurityBarsReply), nil
}

// GetSecurityList 股票列表
func (this *StdApi) GetSecurityList(market int, start uint16) (*SecurityListReply, error) {
	msg := NewSecurityListPackage()
	_market := uint16(market)
	msg.SetParams(&SecurityListRequest{Market: _market, Start: start})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}

	return reply.(*SecurityListReply), nil
}

// GetIndexBars 指数K线
func (this *StdApi) GetIndexBars(market int, code string, category uint16, start uint16, count uint16) (*IndexBarsReply, error) {
	msg := NewIndexBarsPackage()
	_code := [6]byte{}
	_market := uint16(market)
	copy(_code[:], code)
	msg.SetParams(&IndexBarsRequest{
		Market:   _market,
		Code:     _code,
		Category: category,
		Start:    start,
		Count:    count,
	})
	reply, err := this.command(msg)
	if err != nil {
		return nil, err
	}
	return reply.(*IndexBarsReply), err
}
