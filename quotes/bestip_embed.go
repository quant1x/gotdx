package quotes

//
//import (
//	"embed"
//	"fmt"
//	"gitee.com/quant1x/gotdx/internal"
//	"gitee.com/quant1x/gox/api"
//	"gopkg.in/ini.v1"
//	"io"
//	"strings"
//)
//
//var (
//	// ResourcesPath 资源路径
//	ResourcesPath = "resources"
//)
//
////go:embed resources/*
//var resources embed.FS
//
//const (
//	//tdxConfig = "tdx.cfg"
//	//tdxConfig = "zhongxin.cfg"
//	//tdxConfig = "huatai.cfg"
//	tdxConfig              = "guotaijunan.cfg"
//	sectionStandardServer  = "HQHOST"
//	keyHostNum             = "HostNum"
//	sectionExtensionServer = "DSHOST"
//)
//
//var (
//	tdxServerConfigList = []string{"tdx.cfg", "zhongxin.cfg", "huatai.cfg", "guotaijunan.cfg"}
//	tdxServerSourceList = []string{"通达信", "中信证券", "华泰证券", "国泰君安"}
//)
//
//func loadAllConfig() {
//	var standardServers, extensionServers []Server
//	for i, filename := range tdxServerConfigList {
//		std, ext := loadTdxConfig(filename, tdxServerSourceList[i])
//		if len(std) > 0 {
//			standardServers = append(standardServers, std...)
//		}
//		if len(ext) > 0 {
//			extensionServers = append(extensionServers, ext...)
//		}
//	}
//	fmt.Println("----------<" + sectionStandardServer + ">----------")
//	for _, v := range standardServers {
//		fmt.Printf(`{Source: "%s", Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", v.Source, v.Name, v.Host, v.Port)
//	}
//	fmt.Println("----------<" + sectionExtensionServer + ">----------")
//	for _, v := range extensionServers {
//		fmt.Printf(`{Source: "%s", Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", v.Source, v.Name, v.Host, v.Port)
//	}
//}
//
//func loadTdxConfig(name, source string) (std, ext []Server) {
//	fs, err := api.OpenEmbed(resources, ResourcesPath+"/"+name)
//	if err != nil {
//		panic(err)
//	}
//	data, err := io.ReadAll(fs)
//	if err != nil {
//		panic(err)
//	}
//	cfg, err := ini.Load(data)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Println("----------<" + sectionStandardServer + ">----------")
//	section := cfg.Section(sectionStandardServer)
//	if section == nil {
//		return
//	}
//	v := section.Key(keyHostNum).Value()
//	hostNum := int(api.ParseInt(v))
//	for i := 0; i < hostNum; i++ {
//		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
//		hostName = internal.Utf8ToGbk(api.String2Bytes(hostName))
//		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
//		if isIPV6(ipAddress) {
//			continue
//		}
//		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
//		port := int(api.ParseInt(tmpPort))
//		srv := Server{Source: source, Name: hostName, Host: ipAddress, Port: port}
//		std = append(std, srv)
//	}
//	//fmt.Println("----------<" + sectionExtensionServer + ">----------")
//	section = cfg.Section(sectionExtensionServer)
//	if section == nil {
//		return
//	}
//	v = section.Key(keyHostNum).Value()
//	hostNum = int(api.ParseInt(v))
//	for i := 0; i < hostNum; i++ {
//		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
//		hostName = internal.Utf8ToGbk(api.String2Bytes(hostName))
//		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
//		if isIPV6(ipAddress) {
//			continue
//		}
//		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
//		port := int(api.ParseInt(tmpPort))
//		srv := Server{Source: source, Name: hostName, Host: ipAddress, Port: port}
//		ext = append(ext, srv)
//	}
//	return
//}
//
//func isIPV6(address string) bool {
//	arr := strings.Split(address, ":")
//	return len(arr) > 2
//}
