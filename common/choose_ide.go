package common

import (
	"github.com/fatih/color"
	"github.com/miunangel/get-patch-jbr/config"
	"github.com/miunangel/get-patch-jbr/util"
	"os"
	"strconv"
	"strings"
)

type SelectedVersionInfo struct {
	IdeName        string
	IdeVersion     string
	IdeJdkConfPath string
}

func ChooseIde(ideInfos *config.IdeInfos) {

	const (
		TbHeaderIde      = "Scanned IDE"
		OptAllIde        = "Select All IDEs' Latest Version"
		OptChangeJdk     = "Start Change Jdk"
		OptExit          = "Exit"
		OptClear         = "Clear Selected"
		OptShowSelected  = "Show Selected Version"
		CfmSelectIde     = "Please select the IDE"
		CfmSelectVersion = "Please select the IDE version"
	)

	// Record selected version info
	selectedIdes := map[string]SelectedVersionInfo{}

	selectIdeTable := assemblySelectIdeTable(ideInfos)
	ideElements := selectIdeTable.Elements

	// Need cycle to select the IDEs user want to operate
	for true {
		selectIdeTable.ShowSerialTable()
		selectIdeIndex := selectIdeTable.GetNumConfirm(CfmSelectIde)
		selectedIde := ideElements[TbHeaderIde][selectIdeIndex]

		// There are some special option when select
		if selectedIde == OptExit {
			// This is to exit the application
			os.Exit(0)
		} else if selectedIde == OptShowSelected {
			// This is to show the selected infos
			selectedIdesTable := assemblySelectedTable(selectedIdes)
			color.Green("Selected IDE as fallow")
			selectedIdesTable.ShowSerialTable()
			util.PressEnterContinue()
			continue
		} else if selectedIde == OptAllIde {
			// This is to select all last IDE version, simplify user operation
			selectedIdes = selectAllLastVersion(ideInfos)
			color.Green("Selected IDE as fallow")
			selectedIdesTable := assemblySelectedTable(selectedIdes)
			selectedIdesTable.ShowSerialTable()
			util.PressEnterContinue()
			continue
		} else if selectedIde == OptClear {
			// This is to clear all selected version
			selectedIdes = map[string]SelectedVersionInfo{}
			color.Yellow("Selected IDE are cleared")
			util.PressEnterContinue()
			continue
		} else if selectedIde == OptChangeJdk {
			// Selected finish, need confirm to change IDEs' Jdk
			break
		}

		selectVersionTable := assemblySelectVersionTable((*ideInfos)[selectedIde])
		selectVersionTable.ShowSerialTable()
		selectVersionIndex := selectVersionTable.GetNumConfirm(CfmSelectVersion)
		selectedIdes[selectedIde] = SelectedVersionInfo{
			selectedIde,
			(*ideInfos)[selectedIde].IdeVersion[selectVersionIndex],
			(*ideInfos)[selectedIde].IdeJdkConfPath[selectVersionIndex],
		}
	}
}

//  selectAllLastVersion
//  @Description: Get the all last version of ide
//  @param ideInfos: all the IDE infos
//  @return map[string]SelectedVersionInfo: The map of IDE name and version info
func selectAllLastVersion(ideInfos *config.IdeInfos) map[string]SelectedVersionInfo {
	lastVersionMap := map[string]SelectedVersionInfo{}
	for name, info := range *ideInfos {
		newestIndex := getNewestIndex(info.IdeVersion)
		lastVersionMap[name] = SelectedVersionInfo{
			name,
			info.IdeVersion[newestIndex],
			info.IdeJdkConfPath[newestIndex],
		}
	}
	return lastVersionMap
}

//  getNewestIndex
//  @Description: According to JetBrains' name rule, get the last version's index
//  @param versions: The version string array
//  @return int: Last version's index
func getNewestIndex(versions []string) int {
	newestIndex := 0
	// The JetBrains' name rule is: Year.Quarter
	newestSplitVersion := strings.Split(versions[0], ".")
	for index, version := range versions {
		splitVersion := strings.Split(version, ".")
		currentYear, _ := strconv.Atoi(splitVersion[0])
		newestYear, _ := strconv.Atoi(newestSplitVersion[0])
		currentQuarter, _ := strconv.Atoi(splitVersion[1])
		newestQuarter, _ := strconv.Atoi(newestSplitVersion[1])
		if currentYear > newestYear || currentYear == newestYear && currentQuarter > newestQuarter {
			newestIndex = index
			newestSplitVersion = splitVersion
			continue
		}
	}
	return newestIndex
}

//  assemblySelectedTable
//  @Description: Generate the show selected info table
//  @param selectedVersion: The selected info
//  @return *util.SerialTable: The generated table
func assemblySelectedTable(selectedVersion map[string]SelectedVersionInfo) *util.SerialTable {
	const (
		TbHeaderIde     = "IDE"
		TbHeaderVersion = "Version"
	)

	// Assembly the selected info table elements
	infoMap := map[string][]string{}
	for ide, info := range selectedVersion {
		infoMap[TbHeaderVersion] = append(infoMap[TbHeaderVersion], info.IdeVersion)
		infoMap[TbHeaderIde] = append(infoMap[TbHeaderIde], ide)
	}

	selectedTable := util.NewSerialTable(false)
	selectedTable.TableHeader = []string{TbHeaderIde, TbHeaderVersion}
	selectedTable.EleWidth = []int{30, 30}
	selectedTable.Elements = infoMap
	return selectedTable
}

//  assemblySelectVersionTable
//  @Description: Generate the version table
//  @param ideInfo: The info of the selected IDE
//  @return *util.SerialTable: The table generated
func assemblySelectVersionTable(ideInfo *config.IdeInfo) *util.SerialTable {
	const (
		TbHeaderVersion = "Version"
		TbHeaderJDK     = "Using Jdk"
	)

	// Assembly the version table element
	infoMap := map[string][]string{}
	for index, version := range ideInfo.IdeVersion {
		infoMap[TbHeaderVersion] = append(infoMap[TbHeaderVersion], version)
		infoMap[TbHeaderJDK] = append(infoMap[TbHeaderJDK], ideInfo.IdeJdkPath[index])
	}

	selectVersionTable := util.NewSerialTable(false)
	selectVersionTable.TableHeader = []string{TbHeaderVersion, TbHeaderJDK}
	selectVersionTable.EleWidth = []int{20, 60}
	selectVersionTable.Elements = infoMap
	return selectVersionTable
}

//  assemblySelectIdeTable
//  @Description: Generate select IDE table
//  @param ideInfos: The ide infos
//  @return *util.SerialTable: Generated table
func assemblySelectIdeTable(ideInfos *config.IdeInfos) *util.SerialTable {
	const (
		TbHeaderIde     = "Scanned IDE"
		OptAllIde       = "Select All IDEs' Latest Version"
		OptChangeJdk    = "Start Change Jdk"
		OptExit         = "Exit"
		OptClear        = "Clear Selected"
		OptShowSelected = "Show Selected Version"
	)

	// Assembly the ide table element
	ideNames := []string{OptAllIde}
	for key, _ := range *ideInfos {
		ideNames = append(ideNames, key)
	}
	ideNames = append(ideNames, OptChangeJdk, OptShowSelected, OptClear, OptExit)
	ideElements := map[string][]string{TbHeaderIde: ideNames}

	selectIdeTable := util.NewSerialTable(true)
	selectIdeTable.TableHeader = []string{TbHeaderIde}
	selectIdeTable.EleWidth = []int{40}
	selectIdeTable.Elements = ideElements

	return selectIdeTable
}
