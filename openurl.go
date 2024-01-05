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
		return handleWindows(url)
	case "darwin":
		return handleDarwin(url)
	case "linux":
		return handleLinux(url)
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

func runAttachedCmd(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	return c.Run()
}
