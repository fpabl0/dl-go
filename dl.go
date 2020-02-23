package dl

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Download downloads a file
func Download(url string, dst string, progresser ProgressPrinter) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("resource not found")
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		size = 0
	}

	if progresser != nil {
		progresser.Before()
	}

	pw := &progressWriter{Total: uint64(size)}
	if progresser != nil {
		pw.PrintProgress = progresser.Progress
	}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, pw)); err != nil {
		return err
	}

	if progresser != nil {
		progresser.After()
	}

	return nil
}
