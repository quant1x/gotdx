package quotes

import (
	"embed"
)

var (
	// ResourcesPath 资源路径
	ResourcesPath = "resources"
)

//go:embed resources/*
var resources embed.FS

const (
	tdxConfig              = "connect.cfg"
	sectionStandardServer  = "HQHOST"
	keyHostNum             = "HostNum"
	sectionExtensionServer = "DSHOST"
)

//func loadTdxConfig() {
//	fs, err := api.OpenEmbed(resources, ResourcesPath+"/"+tdxConfig)
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
//	fmt.Println("----------<" + sectionStandardServer + ">----------")
//	section := cfg.Section(sectionStandardServer)
//	if section == nil {
//		return
//	}
//	v := section.Key(keyHostNum).Value()
//	hostNum := int(api.ParseInt(v))
//	for i := 0; i < hostNum; i++ {
//		// HostName01=通达信深圳双线主站1
//		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
//		//IPAddress01=110.41.147.114
//		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
//		//Port01=7709
//		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
//		port := int(api.ParseInt(tmpPort))
//		//fmt.Println(hostName, ipAddress, port)
//		fmt.Printf(`{Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", hostName, ipAddress, port)
//	}
//	fmt.Println("----------<" + sectionExtensionServer + ">----------")
//	section = cfg.Section(sectionExtensionServer)
//	if section == nil {
//		return
//	}
//	v = section.Key(keyHostNum).Value()
//	hostNum = int(api.ParseInt(v))
//	for i := 0; i < hostNum; i++ {
//		// HostName01=通达信深圳双线主站1
//		hostName := section.Key(fmt.Sprintf("HostName%02d", i+1)).Value()
//		//IPAddress01=110.41.147.114
//		ipAddress := section.Key(fmt.Sprintf("IPAddress%02d", i+1)).Value()
//		//Port01=7709
//		tmpPort := section.Key(fmt.Sprintf("Port%02d", i+1)).Value()
//		port := int(api.ParseInt(tmpPort))
//		//fmt.Println(hostName, ipAddress, port)
//		fmt.Printf(`{Name: "%s", Host: "%s", Port: %d, CrossTime: 0},`+"\n", hostName, ipAddress, port)
//	}
//}
