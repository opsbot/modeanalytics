package space

import (
	"fmt"

	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CreateCommand returns a cobra command
func CreateCommand() *cobra.Command {
	var spaceName string
	var spaceDesc string
	var spaceType string
	var restricted bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a space",
		Run: func(cmd *cobra.Command, args []string) {
			space := fmt.Sprintf(`{
        "space":{
          "name": "%v",
          "description":"%v",
          "space_type":"%v",
          "restricted":%v
        }
      }`, spaceName, spaceDesc, spaceType, restricted)

			utils.PrettyPrint(space)
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			mode.SpaceCreate([]byte(space))
		},
	}

	cmd.Flags().StringVarP(&spaceName, "name", "n", "", "the space name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&spaceDesc, "description", "d", "", "the space description")
	cmd.Flags().StringVarP(&spaceType, "type", "t", "community", "private | community | custom")
	cmd.Flags().BoolVarP(&restricted, "restricted", "r", false, "is space restricted")

	return cmd
}
