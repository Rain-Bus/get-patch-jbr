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

//  getUserConf
//  @Description: Get User Config from file or use built-in config
//  @receiver userConf
func (userConf *UserConf) getUserConf() {

	var localConf *UserConf

	// Assemble config path
	homePath, _ := os.UserHomeDir()
	defaultPaths := []string{homePath + "/.gpjbr/gpjbr.conf", "/etc/gpjbr/gpjbr.conf", "/etc/gpjbr.conf"}

	// Cycle read config from the config path
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

	// If no config read from file, will use the default
	if localConf != nil {
		userConf.setFileConfig(localConf)
	} else {
		// If user don't want to use built-in config, exit process
		if !util.GetYNConfirm("Use built-in config?", "y") {
			os.Exit(0)
		}
	}
}

//  getConfFromFile
//  @Description: Get config from file
//  @receiver userConf
//  @param file: The file to be read
//  @return *UserConf
//  @return error
func (userConf *UserConf) getConfFromFile(file *os.File) (*UserConf, error) {
	// If file can't read, Maybe no permission
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		color.Red("No Permission: %s", file.Name())
		return nil, err
	}

	// When error occur, the string read from file is not json or format is error
	var tempConf *UserConf
	err = json.Unmarshal(fileByte, &tempConf)
	if err != nil {
		color.Red("Can't Parse: %s", file.Name())
		return nil, err
	}

	// Ask User use this config
	if !util.GetYNConfirm("Use conf file ("+file.Name()+")?", "y") {
		return nil, errors.New("User not use " + file.Name() + "\n")
	}

	// Return the config read form file
	return tempConf, nil
}

//  setFileConfig
//  @Description: Set the config of the file into user config if not empty
//  @receiver userConf
//  @param fileConf: The conf read from config file
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

// NewUserConfig
//  @Description: The construct of UserConf
//  @return *UserConf
func NewUserConfig() *UserConf {
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
