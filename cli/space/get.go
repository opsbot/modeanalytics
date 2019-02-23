package space

import (
	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCommand returns a cobra command
func GetCommand() *cobra.Command {
	var space string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get a space",
		Run: func(cmd *cobra.Command, args []string) {
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			space := mode.SpaceGet(space)
			utils.PrettyPrint(space)
		},
	}

	cmd.Flags().StringVarP(&space, "space", "s", "", "the space id token")
	cmd.MarkFlagRequired("space")

	return cmd
}
