package user

import (
	mode "github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AccountCommand returns a cobra command
func AccountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "get authorizing user",
		Run: func(cmd *cobra.Command, args []string) {
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			mode.AuthGet()
		},
	}
	return cmd
}
