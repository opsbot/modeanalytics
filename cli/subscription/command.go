package subscription

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "subscription api commands",
	}
	cmd.AddCommand()

	return cmd
}
