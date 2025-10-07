package main

import (
	"fmt"
	"os"

	"universal-manifest-sync-cli/manifest"
	"universal-manifest-sync-cli/scanner"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ums",
		Short: "Universal Manifest Sync CLI",
		Long:  "Scans dev environment and generates polyglot manifests for various package managers",
	}

	rootCmd.AddCommand(createScanCommand())
	rootCmd.AddCommand(createGenerateCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createScanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "scan",
		Short: "Scan development environment",
		Run: func(cmd *cobra.Command, args []string) {
			s := scanner.NewScanner()
			result, err := s.Scan()
			if err != nil {
				fmt.Printf("Scan failed: %v\n", err)
				return
			}
			fmt.Printf("Scanned %d packages and %d tools\n", len(result.Packages), len(result.Tools))
		},
	}
}

func createGenerateCommand() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:   "generate [format]",
		Short: "Generate manifest for specified format",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			format := args[0]
			g := manifest.NewGenerator(format)
			if g == nil {
				fmt.Printf("Unsupported format: %s\n", format)
				return
			}

			s := scanner.NewScanner()
			result, err := s.Scan()
			if err != nil {
				fmt.Printf("Scan failed: %v\n", err)
				return
			}

			if err := g.Generate(result, outputFile); err != nil {
				fmt.Printf("generation failed: %v\n", err)
			}
		},
	}
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
	return cmd
}
