package downloader

// Did not actually use this file but used progress.go, easier to implement

import (
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadFile(url string) error {
	// Get the file name from the URL
	fileName := path.Base(url)

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create local file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write response body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
