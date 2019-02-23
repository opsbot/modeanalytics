package report

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "report api commands",
	}
	cmd.AddCommand()

	return cmd
}
