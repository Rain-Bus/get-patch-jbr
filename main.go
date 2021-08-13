package main

import (
	"github.com/miunangel/get-patch-jbr/common"
	"github.com/miunangel/get-patch-jbr/config"
)

func main() {
	//localFileConf := config.GetFileConfig()
	//online := config.NewOnlineConf(localFileConf)
	////common.DownloadFile(localFileConf, online)
	//common.UnzipFile(localFileConf, online)
	//var cmdConf config.CmdConf
	//cmd.Flag(&cmdConf)
	//
	ideInfos := config.GetIdeInfos()
	common.ChooseIde(ideInfos)
	//table := util.NewSerialTable()
	//table.TableHeader = []string{"header1", "header2"}
	//table.EleWidth = []int{20, 10}
	//table.Elements = map[string][]string{
	//	"header1": []string{"ele1.1", "ele1.2"},
	//	"header2": []string{"ele2.1", "ele2.2"},
	//}
	//table.ShowSerialTable()
}
