package cli

import (
	"github.com/spf13/cobra"

	"github.com/opsbot/modeanalytics/cli/org"
	"github.com/opsbot/modeanalytics/cli/space"
	"github.com/opsbot/modeanalytics/cli/user"
)

// AddCommands add the commands from subcommand groups to the root command
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		org.Command(),
		space.Command(),
		user.Command(),
		VersionCommand(),
	)
}
