package config

func GetDownloadUrl(userConf *UserConf, onlineConf *OnlineConf) string {
	//return userConf.DownloadMirrorSite + "/RikudouPatrickstar/JetBrainsRuntime-for-Linux-x64/releases/download/" + onlineConf.Tag + "/" + onlineConf.Filename
	return "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux_x64/909663/chrome-linux.zip"
}

func GetApiUrl(userConf *UserConf) string {
	return userConf.ApiMirrorSite + "/repos/RikudouPatrickstar/JetBrainsRuntime-for-Linux-x64/releases/latest"
}
