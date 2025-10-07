package manifest

import (
	"universal-manifest-sync-cli/scanner"
)

// AptGenerator creates apt manifest
type AptGenerator struct{}

func (g *AptGenerator) Generate(result *scanner.ScanResult, outputFile string) error {
	return nil
}
