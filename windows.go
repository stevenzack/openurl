package openurl

import (
	"os"
	"os/exec"
)

func handleWindows(url string) error {
	if _, err := os.Stat("\"C:/ProgramData/Microsoft/Windows/Start Menu/Programs/Google Chrome\""); !os.IsNotExist(err) { //Chrome  installed on this computer
		exec.Command("cmd", "/C", "copy", "C:\\ProgramData\\Microsoft\\Windows\\StartM~1\\Programs\\Google Chrome.lnk", ".\\GC.lnk").Run()
		return exec.Command("cmd", "/C", "GC.lnk", "--app="+url).Start()
	}
	return exec.Command("explorer", url).Start()
}
