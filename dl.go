package dl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Download downloads a file
func Download(url string, dst string, printProgress PrintProgressFunc) error {

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

	pw := &progressWriter{Total: uint64(size)}
	pw.PrintProgress = printProgress
	if _, err = io.Copy(out, io.TeeReader(resp.Body, pw)); err != nil {
		return err
	}

	if printProgress != nil {
		// The progress use the same line so print a new
		// line once it's finished downloading
		fmt.Print("\n")
	}

	return nil
}
