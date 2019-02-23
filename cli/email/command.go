package email

import (
	"github.com/spf13/cobra"
)

// Command returns a cobra command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "email",
		Short: "email api commands",
	}
	cmd.AddCommand()

	return cmd
}
