package securities

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

var (
	// ResourcesPath 资源路径
	ResourcesPath = "resources"
)

//go:embed resources/*
var resources embed.FS

// OpenEmbed 打开嵌入式文件
func OpenEmbed(name string) (fs.File, error) {
	filename := fmt.Sprintf("%s/%s", ResourcesPath, name)
	reader, err := resources.Open(filename)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

// 导出内嵌资源文件
func export(dest, source string) error {
	src, err := OpenEmbed(source)
	if err != nil {
		return err
	}
	output, err := os.Create(dest)
	if err != nil {
		return err
	}
	//const (
	//	BUFFERSIZE = 8192
	//)
	//buf := make([]byte, BUFFERSIZE)
	//for {
	//	n, err := src.Read(buf)
	//	if err != nil && err != io.EOF {
	//		return err
	//	}
	//	if n == 0 {
	//		break
	//	}
	//
	//	if _, err := output.Write(buf[:n]); err != nil {
	//		return err
	//	}
	//}
	_, err = io.Copy(output, src)
	if err != nil {
		return err
	}
	mtime := time.Now()
	err = os.Chtimes(dest, mtime, mtime)
	return err
}
