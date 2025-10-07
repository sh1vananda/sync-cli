package manifest

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/scanner"
)

// BrewGenerator creates Homebrew manifest
type BrewGenerator struct{}

func (g *BrewGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	content := "# Homebrew manifest placeholder"
	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}
