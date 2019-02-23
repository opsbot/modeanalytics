package run

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run api commands",
	}
	cmd.AddCommand()

	return cmd
}
