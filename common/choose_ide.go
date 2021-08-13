package common

import (
	"github.com/miunangel/get-patch-jbr/config"
	"github.com/miunangel/get-patch-jbr/util"
)

func ChooseIde(ideInfos *config.IdeInfos) {
	const TbHeaderIde = "Scanned IDE"
	const TbHeaderVersion = "Version"
	const TbHeaderJDK = "Using Jdk"
	//selectedIdes := map[string]config.IdeInfo{}

	ideNames := []string{"All IDEs' Latest Version"}
	for key, _ := range *ideInfos {
		ideNames = append(ideNames, key)
	}
	ideNames = append(ideNames, "Start Change Jdk", "Exit")
	ideElements := map[string][]string{TbHeaderIde: ideNames}
	selectIdeTable := util.NewSerialTable(true)
	selectIdeTable.TableHeader = []string{TbHeaderIde}
	selectIdeTable.EleWidth = []int{40}
	selectIdeTable.Elements = ideElements
	selectIdeTable.ShowSerialTable()
	selectIdeIndex := selectIdeTable.GetNumConfirm("Please select the IDE")

	selectedIde := ideElements[TbHeaderIde][selectIdeIndex]

	infoMap := map[string][]string{}
	for index, version := range (*ideInfos)[selectedIde].IdeVersion {
		infoMap[TbHeaderVersion] = append(infoMap[TbHeaderVersion], version)
		infoMap[TbHeaderJDK] = append(infoMap[TbHeaderJDK], (*ideInfos)[selectedIde].IdeJdkPath[index])
	}
	selectVersionTable := util.NewSerialTable(false)
	selectVersionTable.TableHeader = []string{TbHeaderVersion, TbHeaderJDK}
	selectVersionTable.EleWidth = []int{20, 60}
	selectVersionTable.Elements = infoMap
	selectVersionTable.ShowSerialTable()
	selectVersionTable.GetNumConfirm("Please select the IDE version")

}
