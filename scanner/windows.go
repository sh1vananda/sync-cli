package scanner

import (
	"fmt"
	"sort"
	"sync"
)

// WindowsScanner scans for packages and tools on Windows
type WindowsScanner struct{}

func (s *WindowsScanner) Scan() (*ScanResult, error) {
	result := &ScanResult{
		OS:       "windows",
		Packages: []ScoredPackage{},
		Tools:    []ScoredPackage{},
		Metadata: make(map[string]string),
	}

	// Channels for collecting results
	pkgChan := make(chan ScoredPackage, 100)
	var wg sync.WaitGroup

	// 1. Path Scanner
	wg.Add(1)
	go func() {
		defer wg.Done()
		ps := &PathScanner{}
		pkgs, err := ps.Scan()
		if err == nil {
			for _, p := range pkgs {
				pkgChan <- p
			}
		}
	}()

	// 2. Registry Scanner (Stub for now)
	wg.Add(1)
	go func() {
		defer wg.Done()
		rs := &RegistryScanner{}
		pkgs, err := rs.Scan()
		if err == nil {
			for _, p := range pkgs {
				pkgChan <- p
			}
		}
	}()

	// Close channel when all scanners are done
	go func() {
		wg.Wait()
		close(pkgChan)
	}()

	// Collect results
	uniquePkgs := make(map[string]ScoredPackage)
	for pkg := range pkgChan {
		// Deduplicate: keep the one with higher score
		if existing, ok := uniquePkgs[pkg.Name]; ok {
			if pkg.DevScore > existing.DevScore {
				uniquePkgs[pkg.Name] = pkg
			}
		} else {
			uniquePkgs[pkg.Name] = pkg
		}
	}

<<<<<<< HEAD
	// Convert to slice and sort by score
	for _, pkg := range uniquePkgs {
		result.Packages = append(result.Packages, pkg)
=======
// mapPackageName maps Windows package IDs to cross-platform names
func mapPackageName(packageID string) string {
	// Handle special cases with prefixes first
	if strings.HasPrefix(packageID, "Python.Python.") {
		return "python"
	}

	mappings := map[string]string{
		"Git.Git":       "git",
		"Node.js.Node.js":             "node",
		"Docker.Docker":               "docker",
		"Kubernetes.kubectl":          "kubectl",
		"Helm.Helm":                   "helm",
		"OpenJDK.OpenJDK.17":          "openjdk",
		"Microsoft.VisualStudio.Code": "vscode",
>>>>>>> e917df1954192a4da5b79a10d869ad842df073a8
	}

	sort.Slice(result.Packages, func(i, j int) bool {
		return result.Packages[i].DevScore > result.Packages[j].DevScore
	})

	// Separate high-confidence tools
	for _, pkg := range result.Packages {
		if pkg.DevScore >= 50 {
			result.Tools = append(result.Tools, pkg)
		}
	}

	result.Metadata["scan_strategy"] = "heuristic_v1"
	result.Metadata["total_found"] = fmt.Sprintf("%d", len(result.Packages))

	return result, nil
}
