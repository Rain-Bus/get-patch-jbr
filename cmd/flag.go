package cmd

import (
	"flag"
	"github.com/miunangel/get-patch-jbr/config"
)

func Flag(cmdConf *config.CmdConf) {
	Banner()
	flag.StringVar(&cmdConf.DownloadMirrorSite, "d", "", "Appoint the download mirror site")
	flag.StringVar(&cmdConf.ApiMirrorSite, "a", "", "Appoint the api mirror site")
	flag.Parse()
}
