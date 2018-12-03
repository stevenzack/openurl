package openurl

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var unsupportedPlatformError = errors.New("Unsupported platform")
var invalidUrlError = errors.New("Invalid url")

func OpenApp(url string) error {
	if len(url) == 0 {
		return invalidUrlError
	}
	switch runtime.GOOS {
	case "windows":
		if _, err := os.Stat("\"C:/ProgramData/Microsoft/Windows/Start Menu/Programs/Google Chrome\""); !os.IsNotExist(err) { //Chrome  installed on this computer
			exec.Command("cmd", "/C", "copy", "C:\\ProgramData\\Microsoft\\Windows\\StartM~1\\Programs\\Google Chrome.lnk", ".\\GC.lnk").Run()
			return exec.Command("cmd", "/C", "GC.lnk", "--app="+url).Start()
		}
		return exec.Command("explorer", url).Start()
	case "darwin":
		if err := exec.Command("google-chrome", "--app="+url).Start(); err != nil {
			if err := exec.Command("google-chrome-stable", "--app="+url).Start(); err != nil {
				if err := exec.Command("chromium", "--app="+url).Start(); err != nil {
					return exec.Command("open", url).Start()
				}
			}
		}
		return nil
	case "linux":
		if err := exec.Command("google-chrome", "--app="+url).Start(); err != nil {
			if err := exec.Command("google-chrome-stable", "--app="+url).Start(); err != nil {
				if err := exec.Command("chromium", "--app="+url).Start(); err != nil {
					return exec.Command("xdg-open", url).Start()
				}
			}
		}
		return nil
	}
	return unsupportedPlatformError
}

func Open(url string) error {
	if len(url) == 0 {
		return invalidUrlError
	}
	switch runtime.GOOS {
	case "windows":
		return exec.Command("explorer", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	}
	return unsupportedPlatformError
}
