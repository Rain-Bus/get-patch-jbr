package common

import (
	"archive/zip"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/miunangel/get-patch-jbr/config"
	"github.com/miunangel/get-patch-jbr/util"
	"io"
	"os"
	"time"
)

func UnzipFile(userConf *config.UserConf, onlineConf *config.OnlineConf) {

	time.Sleep(1 * time.Second)

	unzipDir := userConf.CandidateDir + "/" + onlineConf.Tag
	zipFile := userConf.DownloadDir + "/" + onlineConf.Filename

	// If unzip dir existed, recreate it
	if util.Exist(unzipDir) {
		err := os.RemoveAll(unzipDir)
		if err != nil {
			color.Red("No Permission: %s", unzipDir)
			os.Exit(0)
		}
		err = os.MkdirAll(unzipDir, 0755)
		if err != nil {
			color.Red("No Permission: %s", unzipDir)
			os.Exit(0)
		}
	}

	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		color.Red("Can't Open ZIP: %s", zipFile)
		os.Exit(0)
	}

	color.Green("Unzipping: %s", zipFile)
	zipFileCount := len(zipReader.Reader.File)
	bar := pb.New(zipFileCount)
	bar.SetRefreshRate(100 * time.Millisecond)
	bar.Start()

	for _, k := range zipReader.Reader.File {

		bar.Increment()

		if k.FileInfo().IsDir() {
			err := os.MkdirAll(unzipDir+"/"+k.Name, 0755)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.Create(unzipDir + "/" + k.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = io.Copy(file, r)
		if err != nil {
			fmt.Println(err)
		}
	}
	bar.Finish()
}
