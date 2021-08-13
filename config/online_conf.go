package config

import (
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

type OnlineConf struct {
	Tag      string
	Filename string
}

func (onlineConf *OnlineConf) getOnlineConf(userConf *UserConf) {

	apiUrl := GetApiUrl(userConf)

	resp, err := http.Get(apiUrl)
	if err != nil {
		color.Red("Can't get response from: %s", apiUrl)
		os.Exit(0)
	}

	body, err := ioutil.ReadAll(resp.Body)
	respStr := string(body)
	// Get the newest Tag
	reg, _ := regexp.Compile(`"tag_name":\s*"(\w+)"`)
	onlineConf.Tag = reg.FindStringSubmatch(respStr)[1]

	// Get the newest zip
	reg, _ = regexp.Compile(`"name":\s*"([^"]+)"`)
	onlineConf.Filename = reg.FindStringSubmatch(respStr)[1] + ".zip"

	// Print the last version package time
	packageTimestamp, _ := time.Parse("200601021504", onlineConf.Tag)
	color.Green("The last version packaged at: %s", packageTimestamp.Format("2006-01-02 15:04"))
}

func NewOnlineConf(userConf *UserConf) *OnlineConf {
	onlineConf := new(OnlineConf)
	onlineConf.getOnlineConf(userConf)
	return onlineConf
}
