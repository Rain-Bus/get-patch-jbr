package common

import (
	"archive/zip"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/miunangel/get-patch-jbr/config"
	"github.com/miunangel/get-patch-jbr/util"
	"io"
	"os"
	"time"
)

// UnzipFile
//  @Description: Unzip the download zip
//  @param userConf
//  @param onlineConf
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

	//
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		color.Red("Can't Open ZIP: %s", zipFile)
		os.Exit(0)
	}

	// Show unzip progress bar
	color.Green("Unzipping: %s", zipFile)
	zipFileCount := len(zipReader.Reader.File)
	bar := pb.New(zipFileCount)
	bar.SetRefreshRate(100 * time.Millisecond)
	bar.Start()

	for _, k := range zipReader.Reader.File {

		// Increase progress bar's progress
		bar.Increment()

		// If os can't create dir, maybe no permission
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(unzipDir+"/"+k.Name, 0755)
			if err != nil {
				color.Red("No Permission: ", unzipDir+"/"+k.Name)
			}
			continue
		}

		r, err := k.Open()
		if err != nil {
			color.Red("Can't open zip file: %s", k.Name)
		}

		file, err := os.Create(unzipDir + "/" + k.Name)
		if err != nil {
			color.Red("Can't unzip to: %s", unzipDir+"/"+k.Name)
			continue
		}

		_, _ = io.Copy(file, r)
	}
	bar.Finish()
}
