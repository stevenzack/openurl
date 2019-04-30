package openurl

import "os/exec"

func handleLinux(url string) error {
	if err := exec.Command("google-chrome", "--app="+url).Start(); err != nil {
		if err := exec.Command("google-chrome-stable", "--app="+url).Start(); err != nil {
			if err := exec.Command("chromium", "--app="+url).Start(); err != nil {
				return exec.Command("xdg-open", url).Start()
			}
		}
	}
	return nil
}
