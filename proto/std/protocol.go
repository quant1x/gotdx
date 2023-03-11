package std

import (
	"bytes"
	"github.com/mymmsc/gox/encoding/binary/struc"
	"log"
)

type Factory func() (Marshaler, Unmarshaler, error)

type Marshaler interface {
	Marshal() ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

// PacketHeader 行情服务器发送第一条指令的返回数据
type PacketHeader struct {
	raw      []byte
	Unknown1 uint `struc:"uint32,little" json:"unknown_1"`

	// B will be encoded/decoded as a 16-bit int (a "short")
	// but is stored as a native int in the struct
	Unknown2 uint `struc:"uint32,little" json:"unknown_2"`

	// the sizeof key links a buffer's size to any int field
	Unknown3 uint `struc:"uint32,little" json:"unknown_3"`
	ZipSize  int  `struc:"uint16,little" json:"zip_size"`

	// you can get freaky if you want
	UnzipSize int `struc:"uint16,little" json:"unzip_size"`
}

func (h *PacketHeader) Bytes() []byte {
	return h.raw
}

func (h *PacketHeader) Compressed() bool {
	return h.ZipSize != h.UnzipSize
}

func (h *PacketHeader) Size() int {
	return h.UnzipSize
}
func (h *PacketHeader) Unmarshal(data []byte) error {
	h.raw = data
	return DefaultUnmarshal(data, h)
}

// DefaultUnmarshal 基于struc包的反序列化方案
func DefaultUnmarshal(data []byte, v interface{}) error {
	// 构造流
	buf := bytes.NewBuffer(data)
	// 使用struc解析到struct
	err := struc.Unpack(buf, v)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DefaultMarshal 基于struc包的序列化方案
func DefaultMarshal(v interface{}) ([]byte, error) {
	// 构造流
	buf := bytes.Buffer{}
	// 使用struc解析到struct
	err := struc.Pack(&buf, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
