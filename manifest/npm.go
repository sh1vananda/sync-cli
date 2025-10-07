package manifest

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/scanner"
)

// NPMGenerator creates package.json
type NPMGenerator struct{}

func (g *NPMGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	content := "{\n  \"name\": \"placeholder\"\n}"
	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(content), 0644)
	}
	fmt.Println(content)
	return nil
}
