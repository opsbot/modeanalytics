package org

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org",
		Short: "org api commands",
	}
	cmd.AddCommand(
		GetCommand(),
	)

	return cmd
}
