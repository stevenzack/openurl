package openurl

func handleDarwin(url string) error {
	e := darwinChromeApp(url)
	if e == nil {
		return nil
	}
	return darwinOpen(url)
}

func darwinChromeApp(url string) error {
	e := runAttachedCmd("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--app="+url)
	return e
}

func darwinOpen(url string) error {
	return runAttachedCmd("open", url)
}
