package slack

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slack",
		Short: "slack api commands",
	}
	cmd.AddCommand()

	return cmd
}
