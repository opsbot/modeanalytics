package space

import (
	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCommand returns a cobra command
func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list spaces",
		Run: func(cmd *cobra.Command, args []string) {
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			spaces := mode.SpaceList()
			utils.PrettyPrint(spaces)
		},
	}
	return cmd
}
