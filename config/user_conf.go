package config

import (
	"encoding/json"
	"errors"
	"github.com/fatih/color"
	"github.com/miunangel/get-patch-jbr/util"
	"io/ioutil"
	"os"
	"strings"
)

type UserConf struct {
	ProgramDir         string
	DownloadDir        string
	CandidateDir       string
	ApiMirrorSite      string `json:"api_mirror_site"`
	DownloadMirrorSite string `json:"download_mirror_site"`
}

func (userConf *UserConf) getUserConf() {

	var (
		localConf *UserConf
	)

	homePath, _ := os.UserHomeDir()
	defaultPaths := []string{homePath + "/.gpjbr/gpjbr.conf", "/etc/gpjbr/gpjbr.conf", "/etc/gpjbr.conf"}

	for _, defaultPath := range defaultPaths {
		file, err := os.Open(defaultPath)
		if err != nil {
			continue
		}
		localConf, err = userConf.getConfFromFile(file)
		if err != nil {
			continue
		}
		break
	}

	if localConf != nil {
		userConf.setFileConfig(localConf)
	} else {
		// If user don't want to use built-in config, exit process
		if !util.GetYNConfirm("Use built-in config?", "y") {
			os.Exit(0)
		}
	}
}

func (userConf *UserConf) getConfFromFile(file *os.File) (*UserConf, error) {
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		color.Red("No Permission: %s", file.Name())
		return nil, err
	}
	var tempConf *UserConf
	err = json.Unmarshal(fileByte, &tempConf)
	if err != nil {
		color.Red("Can't Parse: %s", file.Name())
		return nil, err
	}
	if !util.GetYNConfirm("Use conf file ("+file.Name()+")?", "y") {
		return nil, errors.New("User not use " + file.Name() + "\n")
	}
	return tempConf, nil
}

// setFileConfig	Set the config of the file into user config
func (userConf *UserConf) setFileConfig(fileConf *UserConf) {

	fileConf.ApiMirrorSite = strings.TrimSpace(fileConf.ApiMirrorSite)
	if fileConf.ApiMirrorSite != "" {
		userConf.ApiMirrorSite = fileConf.ApiMirrorSite
	}

	fileConf.DownloadMirrorSite = strings.TrimSpace(fileConf.DownloadMirrorSite)
	if fileConf.DownloadMirrorSite != "" {
		userConf.DownloadMirrorSite = fileConf.DownloadMirrorSite
	}

}

func GetFileConfig() *UserConf {
	homedir, _ := os.UserHomeDir()
	programDir := homedir + "/.gpjbr"
	defaultConf := &UserConf{
		programDir,
		programDir + "/download",
		programDir + "/candidate",
		"https://api.github.com",
		"https://github.com",
	}
	defaultConf.getUserConf()
	return defaultConf
}
