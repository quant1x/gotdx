package quotes

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"gitee.com/quant1x/gotdx/util"
	"github.com/mymmsc/gox/api"
	log "github.com/mymmsc/gox/logger"
	"io"
	"time"
)

// StdRequestHeader 标准行情-请求-消息头
type StdRequestHeader struct {
	ZipFlag    uint8  `struc:"uint8,little"`  // ZipFlag
	SeqID      uint32 `struc:"uint32,little"` // 请求编号
	PacketType uint8  `struc:"uint8,little"`  // 包类型
	PkgLen1    uint16 `struc:"uint16,little"` // 消息体长度1
	PkgLen2    uint16 `struc:"uint16,little"` // 消息体长度2
	Method     uint16 `struc:"uint16,little"` // method 请求方法
}

// StdResponseHeader 标准行情-响应-消息头
type StdResponseHeader struct {
	I1        uint32 `struc:"uint32,little"`
	ZipFlag   uint8  `struc:"uint8,little"`  // ZipFlag
	SeqID     uint32 `struc:"uint32,little"` // 请求编号
	I3        uint8  `struc:"uint8,little"`
	Method    uint16 `struc:"uint16,little"` // method
	ZipSize   uint16 `struc:"uint16,little"` // 长度
	UnZipSize uint16 `struc:"uint16,little"` // 未压缩长度
}

// Message 消息接口
type Message interface {
	// Serialize 编码
	Serialize() ([]byte, error)
	// UnSerialize 解码
	UnSerialize(head interface{}, in []byte) error
	// Reply 获取返回值
	Reply() interface{}
}

// 消息处理
func process(client *TcpClient, msg Message) error {
	defer client.updateCompletedTimestamp()
	conn := client.conn
	opt := client.opt
	// 1. 序列化
	sendData, err := msg.Serialize()
	if err != nil {
		return err
	}

	// 2. 发送指令
	retryTimes := 0
	if log.IsDebug() {
		log.Debug(util.Bytes2HexString(sendData))
	}
	for {
		// 设置写timeout
		err = conn.SetWriteDeadline(time.Now().Add(opt.WriteTimeout))
		if err != nil {
			return err
		}
		n, err := conn.Write(sendData)
		if n < len(sendData) {
			retryTimes++
			if retryTimes <= opt.MaxRetryTimes {
				log.Warnf("第%d次重试\n", retryTimes)
			} else {
				return err
			}
		} else {
			if err != nil {
				return err
			}
			break
		}
	}

	// 3. 读取响应
	// 3.1 读取响应的消息头
	headerBytes := make([]byte, MessageHeaderBytes)
	// 设置读timeout
	err = conn.SetReadDeadline(time.Now().Add(opt.ReadTimeout))
	if err != nil {
		return err
	}
	_, err = io.ReadFull(conn, headerBytes)
	if err != nil {
		return err
	}
	if log.IsDebug() {
		log.Debug("response header: ", hex.EncodeToString(headerBytes))
	}

	// 3.2 响应的消息头, 反序列化
	headerBuf := bytes.NewReader(headerBytes)
	var header StdResponseHeader
	if err := binary.Read(headerBuf, binary.LittleEndian, &header); err != nil {
		return err
	}
	if log.IsDebug() {
		log.Debugf("response header: %+v", header)
	}
	// 3.3 处理超长信息的异常
	if header.ZipSize > MessageMaxBytes {
		log.Warnf("msgData has bytes(%d) beyond max %d\n", header.ZipSize, MessageMaxBytes)
		return ErrBadData
	}
	// 3.4 读取响应的消息体
	msgData := make([]byte, header.ZipSize)
	// 设置读timeout
	err = conn.SetReadDeadline(time.Now().Add(opt.ReadTimeout))
	if err != nil {
		return err
	}
	_, err = io.ReadFull(conn, msgData)
	if err != nil {
		return err
	}
	// 3.5 反序列化响应的消息体
	var out bytes.Buffer
	if log.IsDebug() {
		log.Debugf("response body: %+v", hex.EncodeToString(msgData))
	}
	var respBody []byte
	if header.ZipSize != header.UnZipSize {
		b := bytes.NewReader(msgData)
		r, _ := zlib.NewReader(b)
		defer api.CloseQuietly(r)
		_, _ = io.Copy(&out, r)
		respBody = out.Bytes()
	} else {
		respBody = msgData
	}
	if log.IsDebug() {
		log.Debugf("response body: %+v", hex.EncodeToString(respBody))
	}
	err = msg.UnSerialize(&header, respBody)
	// 4. 返回
	return err
}
