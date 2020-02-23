package dl

// ProgressPrinter interface
type ProgressPrinter interface {
	Before()
	Progress(progress uint64, total uint64)
	After()
}
