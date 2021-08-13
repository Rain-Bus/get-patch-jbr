package config

import (
	"github.com/miunangel/get-patch-jbr/util"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type IdeInfo struct {
	IdeName        string
	IdeVersion     []string
	IdeJdkPath     []string
	IdeJdkConfPath []string
}

type IdeInfos map[string]*IdeInfo

func (ideInfos IdeInfos) scanIde() {
	// Assembly and read JetBrains' IDE config dir
	homeDir, _ := os.UserHomeDir()
	jetConfDir := homeDir + "/.config/JetBrains"
	dirs, _ := os.ReadDir(jetConfDir)

	// Init regular exception to match dir name
	reg, _ := regexp.Compile("([A-Z][A-Za-z]+)([0-9.]+)")
	for _, dir := range dirs {
		if dir.IsDir() && reg.MatchString(dir.Name()) {

			// split dir name to ide name and ide version
			splitDirName := reg.FindStringSubmatch(dir.Name())
			splitIdeName := splitDirName[1]
			splitIdeVersion := splitDirName[2]

			// Read the jdk config of IDEs
			jetIdeConfPath := jetConfDir + "/" + dir.Name()
			if util.Exist(jetIdeConfPath) {
				// Some IDE will use the simple name create the {ide}.jdk
				if simpleIdeName := ideInfos.simplifyIdeName(splitIdeName); simpleIdeName != "" {
					splitIdeName = simpleIdeName
				}

				// Get the {ide}.jdk file, if it doesn't exist
				jetJdkConfPath := jetIdeConfPath + "/" + strings.ToLower(splitIdeName) + ".jdk"
				fileReader, _ := os.Open(jetJdkConfPath)
				jdkPath, err := ioutil.ReadAll(fileReader)
				jdkPathStr := strings.TrimSpace(string(jdkPath))
				if err != nil || jdkPathStr == "" {
					jdkPathStr = "IDE Insert"
				}
				if ideInfos[splitIdeName] == nil {
					ideInfos[splitIdeName] = &IdeInfo{
						splitIdeName,
						[]string{splitIdeVersion},
						[]string{jdkPathStr},
						[]string{jetJdkConfPath},
					}
				} else {
					ideInfos[splitIdeName].IdeVersion = append(ideInfos[splitIdeName].IdeVersion, splitIdeVersion)
					ideInfos[splitIdeName].IdeJdkPath = append(ideInfos[splitIdeName].IdeJdkPath, jdkPathStr)
					ideInfos[splitIdeName].IdeJdkConfPath = append(ideInfos[splitIdeName].IdeJdkConfPath, jetJdkConfPath)
				}
			}
		}
	}
}

func (ideInfos IdeInfos) simplifyIdeName(ideName string) string {
	simplifyMap := map[string]string{"IntelliJIdea": "Idea"}
	return simplifyMap[ideName]
}

func (ideInfos IdeInfos) restoreIdeName(simpleName string) string {
	restoreMap := map[string]string{"Idea": "IntelliJIdea"}
	return restoreMap[simpleName]
}

func GetIdeInfos() *IdeInfos {
	ideInfos := IdeInfos{}
	ideInfos.scanIde()
	return &ideInfos
}
