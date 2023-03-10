package quotes

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	log "github.com/mymmsc/gox/logger"
	"io"
	"net"
)

// StdRequestHeader 标准行情-请求-消息头
type StdRequestHeader struct {
	Zip        uint8  // ZipFlag
	SeqID      uint32 // 请求编号
	PacketType uint8
	PkgLen1    uint16
	PkgLen2    uint16
	Method     uint16 // method 请求方法
}

// StdResponseHeader 标准行情-响应-消息头
type StdResponseHeader struct {
	I1        uint32
	I2        uint8
	SeqID     uint32 // 请求编号
	I3        uint8
	Method    uint16 // method
	ZipSize   uint16 // 长度
	UnZipSize uint16 // 未压缩长度
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
func process(conn net.Conn, msg Message, opt Opt) error {
	// 1. 序列化
	sendData, err := msg.Serialize()
	if err != nil {
		return err
	}

	// 2. 发送指令
	retryTimes := 0
	//fmt.Println(util.Bytes2HexString(sendData))
	for {
		n, err := conn.Write(sendData)
		if n < len(sendData) {
			retryTimes++
			if retryTimes <= opt.MaxRetryTimes {
				log.Debugf("第%d次重试\n", retryTimes)
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
	_, err = io.ReadFull(conn, headerBytes)
	if err != nil {
		return err
	}

	// 3.2 响应的消息头, 反序列化
	headerBuf := bytes.NewReader(headerBytes)
	var header StdResponseHeader
	if err := binary.Read(headerBuf, binary.LittleEndian, &header); err != nil {
		return err
	}
	// 3.3 处理超长信息的异常
	if header.ZipSize > MessageMaxBytes {
		log.Debugf("msgData has bytes(%d) beyond max %d\n", header.ZipSize, MessageMaxBytes)
		return ErrBadData
	}
	// 3.4 读取响应的消息体
	msgData := make([]byte, header.ZipSize)
	_, err = io.ReadFull(conn, msgData)
	if err != nil {
		return err
	}
	// 3.5 反序列化响应的消息体
	var out bytes.Buffer
	if header.ZipSize != header.UnZipSize {
		b := bytes.NewReader(msgData)
		r, _ := zlib.NewReader(b)
		// TODO: 这里可能存在bug
		_, _ = io.Copy(&out, r)
		err = msg.UnSerialize(&header, out.Bytes())
	} else {
		err = msg.UnSerialize(&header, msgData)
	}
	// 4. 返回
	return err
}
