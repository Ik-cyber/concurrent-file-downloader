package downloader

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

type Downloader struct {
	Progress *mpb.Progress
	MaxRetry int
}

func NewDownloader(p *mpb.Progress, maxRetry int) *Downloader {
	return &Downloader{Progress: p, MaxRetry: maxRetry}
}

func (d *Downloader) DownloadFile(ctx context.Context, url string) error {
	fileName := path.Base(url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	size := resp.ContentLength

	bar := d.Progress.AddBar(size,
		mpb.PrependDecorators(
			decor.Name(fileName+" ", decor.WC{W: len(fileName) + 1, C: decor.DSyncWidth}),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
			decor.OnComplete(decor.Name(" ✔️"), "done"),
		),
	)

	proxyReader := bar.ProxyReader(resp.Body)
	defer proxyReader.Close()

	buf := make([]byte, 32*1024) // 32 KB buffer
	for {
		select {
		case <-ctx.Done():
			return errors.New("download cancelled")
		default:
			n, err := proxyReader.Read(buf)
			if n > 0 {
				if _, writeErr := out.Write(buf[:n]); writeErr != nil {
					return writeErr
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}
