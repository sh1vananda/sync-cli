package scanner

import (
    "fmt"
    "os/exec"
    "strings"
)

// WindowsScanner scans for packages and tools on Windows
type WindowsScanner struct{}

func (s *WindowsScanner) Scan() (*ScanResult, error) {
    result := &ScanResult{
        OS:       "windows",
        Packages: make(map[string]struct{}),
        Tools:    make(map[string]struct{}),
    }

    // Get installed packages using winget
    cmd := exec.Command("winget", "list", "--accept-source-agreements")
    output, err := cmd.Output()
    if err != nil {
        return nil, fmt.Errorf("winget list failed: %v", err)
    }

    lines := strings.Split(string(output), "\n")
    // Skip the header lines (first 3 lines)
    for i, line := range lines {
        if i < 3 {
            continue
        }
        if line == "" {
            continue
        }
        fields := strings.Fields(line)
        if len(fields) < 2 {
            continue
        }
        packageID := fields[1]

        // Skip Windows system packages
        if isWindowsSystemPackage(packageID) {
            continue
        }

        // Map to cross-platform name
        mappedName := mapPackageName(packageID)
        if mappedName == "" {
            continue
        }

        result.Packages[mappedName] = struct{}{}
    }

    // Add common tools
    commonTools := []string{"git", "node", "python", "docker", "kubectl", "helm"}
    for _, tool := range commonTools {
        if isToolInstalled(tool) {
            result.Tools[tool] = struct{}{}
        }
    }

    return result, nil
}

// isWindowsSystemPackage returns true if the package is a Windows system component
func isWindowsSystemPackage(packageID string) bool {
	denylist := []string{
		"Microsoft.",
		"MSIX",
		"Xaml",
		"Windows Runtime",
		"Windows SDK",
		"Windows Driver Kit",
		"Windows Subsystem for Linux",
	}

	for _, prefix := range denylist {
		if strings.HasPrefix(packageID, prefix) {
			return true
		}
	}
	return false
}

// mapPackageName maps Windows package IDs to cross-platform names
func mapPackageName(packageID string) string {
	mappings := map[string]string{
		"Git.Git":                     "git",
		"Python.Python.3.12":          "python",
		"Node.js.Node.js":             "node",
		"Docker.Docker":               "docker",
		"Kubernetes.kubectl":          "kubectl",
		"Helm.Helm":                   "helm",
		"OpenJDK.OpenJDK.17":          "openjdk",
		"Microsoft.VisualStudio.Code": "vscode",
	}

	if mapped, ok := mappings[packageID]; ok {
		return mapped
	}

	// Fallback: use the package ID if no mapping is found
	return packageID
}

// isToolInstalled checks if a tool is installed by trying to run it
func isToolInstalled(tool string) bool {
	_, err := exec.LookPath(tool)
	return err == nil
}
