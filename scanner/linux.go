package scanner

// LinuxScanner scans for packages and tools on Linux
type LinuxScanner struct{}

func (s *LinuxScanner) Scan() (*ScanResult, error) {
	return &ScanResult{
		OS:       "linux",
		Packages: []ScoredPackage{},
		Tools:    []ScoredPackage{},
		Metadata: make(map[string]string),
	}, nil
}
