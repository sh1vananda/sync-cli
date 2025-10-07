package manifest

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/scanner"
)

// WingetGenerator creates winget manifest
type WingetGenerator struct{}

func (g *WingetGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	content := "# Winget manifest placeholder"
	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}
