package dl

// PrintProgressFunc type
type PrintProgressFunc func(progress uint64, total uint64)

type progressWriter struct {
	Progress      uint64
	Total         uint64
	PrintProgress PrintProgressFunc
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Progress += uint64(n)
	if pw.PrintProgress != nil {
		pw.PrintProgress(pw.Progress, pw.Total)
	}
	return n, nil
}
