package schedule

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule",
		Short: "schedule api commands",
	}
	cmd.AddCommand()

	return cmd
}
