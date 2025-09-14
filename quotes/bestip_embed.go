package quotes

import (
	"embed"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"text/template"

	"gitee.com/quant1x/gotdx/internal"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/logger"
	"gopkg.in/ini.v1"
)

var (
	// ResourcesPath 资源路径
	ResourcesPath = "resources"
)

//go:embed resources/*
var resources embed.FS

const (
	sectionStandardServer  = "HQHOST"
	defaultStandardPort    = 7709
	keyHostNum             = "HostNum"
	sectionExtensionServer = "DSHOST"
	defaultExtensionPort   = 7727
)

var (
	ignoreStandardPortList  = []int{80} // 标准行情需要忽略的端口号
	ignoreExtensionPortList = []int{}   // 扩展行情需要忽略的端口号
)

type tdxConfig struct {
	source   string
	filename string
}

var (
	tdxServerList = []tdxConfig{
		tdxConfig{source: "通达信", filename: "tdx.cfg"},
		tdxConfig{source: "中信证券", filename: "zhongxin.cfg"},
		tdxConfig{source: "华泰证券", filename: "huatai.cfg"},
		tdxConfig{source: "国泰君安", filename: "guotaijunan.cfg"},
	}
)

var (
	templateAddress = `package quotes

var (
	// StandardServerList 标准行情服务器列表
	StandardServerList = []Server{
		{{- range .Std}}
		{Source: "{{.Source}}", Name: "{{.Name}}", Host: "{{.Host}}", Port: {{.Port}}, CrossTime: 0},
		{{- end}}
	}
	// ExtensionServerList 扩展行情服务器列表
	ExtensionServerList = []Server{
		{{- range .Ext}}
		{Source: "{{.Source}}", Name: "{{.Name}}", Host: "{{.Host}}", Port: {{.Port}}, CrossTime: 0},
		{{- end}}
	}
)

`
)

func loadAllConfig() {
	var standardServers, extensionServers []Server
	for _, config := range tdxServerList {
		std, ext := loadTdxConfig(config)
		if len(std) > 0 {
			standardServers = append(standardServers, std...)
		}
		if len(ext) > 0 {
			extensionServers = append(extensionServers, ext...)
		}
	}
	fmt.Println("----------<" + sectionStandardServer + ">----------")
	for _, v := range standardServers {
		fmt.Printf(`{Source: "%s", Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", v.Source, v.Name, v.Host, v.Port)
	}
	fmt.Println("----------<" + sectionExtensionServer + ">----------")
	for _, v := range extensionServers {
		fmt.Printf(`{Source: "%s", Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", v.Source, v.Name, v.Host, v.Port)
	}
	tmpl, err := template.New("address").Parse(templateAddress)
	if err != nil {
		logger.Fatalf("解析服务器地址模版失败, error=%+v", err)
	}
	data := struct {
		Std []Server
		Ext []Server
	}{
		Std: standardServers,
		Ext: extensionServers,
	}
	writer, err := os.Create("bestip_address.go")
	if err != nil {
		logger.Fatalf("创建bestip_address.go源文件失败, error=%+v", err)
	}
	err = tmpl.Execute(writer, data)
	if err != nil {
		logger.Fatalf("执行服务器地址模版失败, error=%+v", err)
	}
}

func loadTdxConfig(config tdxConfig) (std, ext []Server) {
	name := config.filename
	source := config.source
	fs, err := api.OpenEmbed(resources, ResourcesPath+"/"+name)
	if err != nil {
		logger.Fatalf("%+v", err)
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		logger.Fatalf("%+v", err)
	}
	cfg, err := ini.Load(data)
	if err != nil {
		logger.Fatalf("%+v", err)
	}
	//fmt.Println("----------<" + sectionStandardServer + ">----------")
	section := cfg.Section(sectionStandardServer)
	if section == nil {
		return
	}
	v := section.Key(keyHostNum).Value()
	hostNum := int(api.ParseInt(v))
	for i := 0; i < hostNum; i++ {
		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
		hostName = internal.Utf8ToGbk(api.String2Bytes(hostName))
		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
		if isIPV6(ipAddress) {
			continue
		}
		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
		port := int(api.ParseInt(tmpPort))
		if slices.Contains(ignoreStandardPortList, port) {
			continue
		}
		srv := Server{Source: source, Name: hostName, Host: ipAddress, Port: port}
		std = append(std, srv)
	}
	//fmt.Println("----------<" + sectionExtensionServer + ">----------")
	section = cfg.Section(sectionExtensionServer)
	if section == nil {
		return
	}
	v = section.Key(keyHostNum).Value()
	hostNum = int(api.ParseInt(v))
	for i := 0; i < hostNum; i++ {
		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
		hostName = internal.Utf8ToGbk(api.String2Bytes(hostName))
		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
		if isIPV6(ipAddress) {
			continue
		}
		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
		port := int(api.ParseInt(tmpPort))
		if slices.Contains(ignoreExtensionPortList, port) {
			continue
		}
		srv := Server{Source: source, Name: hostName, Host: ipAddress, Port: port}
		ext = append(ext, srv)
	}
	return
}

func isIPV6(address string) bool {
	arr := strings.Split(address, ":")
	return len(arr) > 2
}
