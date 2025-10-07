package scanner

// MacScanner scans macOS environment
type MacScanner struct{}

func (s *MacScanner) Scan() (*ScanResult, error) {
    return &ScanResult{
        OS:       "darwin",
        Packages: make(map[string]struct{}),
        Tools:    make(map[string]struct{}),
    }, nil
}
