package space

import (
	"github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DeleteCommand returns a cobra command
func DeleteCommand() *cobra.Command {
	var space string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a space",
		Run: func(cmd *cobra.Command, args []string) {
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			mode.SpaceDelete(space)
		},
	}

	cmd.Flags().StringVarP(&space, "space", "s", "", "the space id token")
	cmd.MarkFlagRequired("space")

	return cmd
}
