package scanner

// MacScanner scans for packages and tools on macOS
type MacScanner struct{}

func (s *MacScanner) Scan() (*ScanResult, error) {
	return &ScanResult{
		OS:       "darwin",
		Packages: []ScoredPackage{},
		Tools:    []ScoredPackage{},
		Metadata: make(map[string]string),
	}, nil
}
