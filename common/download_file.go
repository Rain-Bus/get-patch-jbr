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

	if !util.Exist(userConf.DownloadDir) {
		err := os.Mkdir(userConf.DownloadDir, 0755)
		if err != nil {
			color.Red("No Permission: %s", userConf.DownloadDir)
			os.Exit(0)
		}
	}
	downloadPath := userConf.DownloadDir + "/" + onlineConf.Filename
	if util.Exist(downloadPath) {
		err := os.Remove(downloadPath)
		if err != nil {

		}
	}
	out, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	bar := pb.Full.Start64(fileSize)
	barWriter := bar.NewProxyWriter(out)
	_, _ = io.Copy(barWriter, resp.Body)
	bar.Finish()
}
