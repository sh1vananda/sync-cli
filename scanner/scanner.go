package scanner

// ScanResult holds the result of a scan
type ScanResult struct {
 OS       string
 Packages map[string]struct{}
 Tools    map[string]struct{}
}

// Scanner is the interface for scanning the environment
type Scanner interface {
 Scan() (*ScanResult, error)
}
