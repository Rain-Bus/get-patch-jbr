package common

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/miunangel/get-patch-jbr/config"
	"github.com/miunangel/get-patch-jbr/util"
	"io"
	"net/http"
	"os"
)

// DownloadFile
//  @Description: Download last version from GitHub
//  @param userConf
//  @param onlineConf
//
func DownloadFile(userConf *config.UserConf, onlineConf *config.OnlineConf) {
	downloadUrl := config.GetDownloadUrl(userConf, onlineConf)

	color.Green("Downloading from %s", downloadUrl)
	resp, err := http.Get(downloadUrl)
	if err != nil {
		color.Red("Fail to download from: %s", downloadUrl)
	}

	fileSize := resp.ContentLength
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// If download dir not exist, create it
	if !util.Exist(userConf.DownloadDir) {
		err := os.Mkdir(userConf.DownloadDir, 0755)
		if err != nil {
			color.Red("No Permission: %s", userConf.DownloadDir)
			os.Exit(0)
		}
	}

	// If the file already exist, delete it, and then download the file
	downloadPath := userConf.DownloadDir + "/" + onlineConf.Filename
	if util.Exist(downloadPath) {
		err := os.Remove(downloadPath)
		if err != nil {
			color.Red("No Permission: %s", downloadPath)
		}
	}
	out, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			color.Red("Download error: %s", downloadPath)
		}
	}(out)

	// Show download process bar
	bar := pb.Full.Start64(fileSize)
	barWriter := bar.NewProxyWriter(out)
	_, _ = io.Copy(barWriter, resp.Body)
	bar.Finish()
}
