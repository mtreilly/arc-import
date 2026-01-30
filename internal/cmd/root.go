// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/yourorg/arc-sdk/output"
)

// NewRootCmd creates the root command for arc-import.
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arc-import",
		Short: "Import existing filesystem content into the database",
		Example: strings.TrimSpace(`
  arc-import articles        # future Phase 2 ingest
  arc-import papers          # placeholder for academic PDFs
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(articlesCmd())
	cmd.AddCommand(papersCmd())

	return cmd
}

func articlesCmd() *cobra.Command {
	var opts output.OutputOptions

	c := &cobra.Command{
		Use:   "articles",
		Short: "Import articles from docs/research-external/articles (stub)",
		Example: strings.TrimSpace(`
  arc-import articles
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Resolve(); err != nil {
				return err
			}
			return emitImportStub(cmd, opts, "articles")
		},
	}
	opts.AddOutputFlags(c, output.OutputTable)
	return c
}

func papersCmd() *cobra.Command {
	var opts output.OutputOptions

	c := &cobra.Command{
		Use:   "papers",
		Short: "Import papers from docs/research-external/papers (stub)",
		Example: strings.TrimSpace(`
  arc-import papers
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Resolve(); err != nil {
				return err
			}
			return emitImportStub(cmd, opts, "papers")
		},
	}
	opts.AddOutputFlags(c, output.OutputTable)
	return c
}

func emitImportStub(cmd *cobra.Command, opts output.OutputOptions, collection string) error {
	result := map[string]any{
		"collection": collection,
		"status":     "stub",
		"next_step":  "Phase 2 ingestion",
	}
	switch {
	case opts.Is(output.OutputJSON):
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode(result)
	case opts.Is(output.OutputYAML):
		enc := yaml.NewEncoder(cmd.OutOrStdout())
		defer enc.Close()
		return enc.Encode(result)
	case opts.Is(output.OutputQuiet):
		return nil
	default:
		fmt.Fprintf(cmd.OutOrStdout(), "import %s: stub - to be implemented in Phase 2\n", collection)
		return nil
	}
}
