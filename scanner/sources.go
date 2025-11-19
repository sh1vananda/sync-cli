package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

// PathScanner scans the PATH environment variable for executables
type PathScanner struct{}

func (s *PathScanner) Scan() ([]ScoredPackage, error) {
	var packages []ScoredPackage
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	seen := make(map[string]bool)

	for _, dir := range paths {
		// Skip system directories to reduce noise
		if isSystemDir(dir) {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := entry.Name()
			ext := strings.ToLower(filepath.Ext(name))
			if ext != ".exe" && ext != ".bat" && ext != ".cmd" {
				continue
			}

			baseName := strings.TrimSuffix(name, ext)
			if seen[baseName] {
				continue
			}

			// Score the found executable
			score := ScorePackage(baseName, "", dir, "")
			if score >= 25 { // Threshold for PATH items
				packages = append(packages, ScoredPackage{
					Name:       baseName,
					Source:     "path",
					DevScore:   score,
					Confidence: 0.8,
					Metadata:   map[string]string{"path": filepath.Join(dir, name)},
				})
				seen[baseName] = true
			}
		}
	}
	return packages, nil
}

func isSystemDir(path string) bool {
	lower := strings.ToLower(path)
	return strings.Contains(lower, "windows\\system32") ||
		strings.Contains(lower, "windows\\syswow64")
}

// RegistryScanner scans the Windows Registry (stub for now, can be expanded)
type RegistryScanner struct{}

func (s *RegistryScanner) Scan() ([]ScoredPackage, error) {
	// In a real implementation, we would use "golang.org/x/sys/windows/registry"
	// For this prototype, we'll return an empty list or mock data if needed.
	// Since we can't easily add external dependencies without go get,
	// we will rely primarily on PATH and Winget for now, or implement a basic
	// command-based registry query if strictly necessary.

	// Returning empty for now to avoid adding complex dependencies in this step.
	return []ScoredPackage{}, nil
}
