package scanner

// LinuxScanner scans Linux environment
type LinuxScanner struct{}

func (s *LinuxScanner) Scan() (*ScanResult, error) {
    return &ScanResult{
        OS:       "linux",
        Packages: make(map[string]struct{}),
        Tools:    make(map[string]struct{}),
    }, nil
}
