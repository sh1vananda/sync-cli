package manifest

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/scanner"
)

// PipGenerator creates pip requirements
type PipGenerator struct{}

func (g *PipGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	content := "# requirements.txt placeholder"
	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}
