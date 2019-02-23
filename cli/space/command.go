package space

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "space",
		Short: "space api commands",
	}
	cmd.AddCommand(
		ListCommand(),
		CreateCommand(),
		GetCommand(),
		DeleteCommand(),
	)

	return cmd
}
