package openurl

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var unsupportedPlatformError = errors.New("Unsupported platform")
var invalidUrlError = errors.New("Invalid url")

func Open(url string) error {
	if len(url) == 0 {
		return invalidUrlError
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	switch runtime.GOOS {
	case "windows":
		if _, err := os.Stat("\"C:/ProgramData/Microsoft/Windows/Start Menu/Programs/Google Chrome\""); !os.IsNotExist(err) { //Chrome  installed on this computer
			file, _ := os.Create("./openurl.bat")
			file.WriteString("\"C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Google Chrome.lnk\" --app=" + url)
			file.Close()
			cmd := exec.Command(".\\openurl.bat")
			return cmd.Start()
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
