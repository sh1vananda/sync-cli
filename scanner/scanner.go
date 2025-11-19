package scanner

// ScanResult holds the result of a scan
// ScoredPackage represents a discovered package with a confidence score
type ScoredPackage struct {
	Name       string
	Version    string
	Source     string // "path", "registry", "winget"
	DevScore   int    // 0-100
	Confidence float64
	Metadata   map[string]string
}

// ScanResult holds the result of a scan
type ScanResult struct {
	OS       string
	Packages []ScoredPackage
	Tools    []ScoredPackage
	Metadata map[string]string
}

// Scanner is the interface for scanning the environment
type Scanner interface {
	Scan() (*ScanResult, error)
}
