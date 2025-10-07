package manifest

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/scanner"
)

// VcpkgGenerator creates vcpkg manifest
type VcpkgGenerator struct{}

func (g *VcpkgGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	content := "# Vcpkg manifest placeholder"
	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}
