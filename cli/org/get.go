package org

import (
	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/modeanalytics/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCommand returns a cobra command
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get org info",
		Run: func(cmd *cobra.Command, args []string) {
			mode.SetEnvironment(viper.GetStringMap(cmd.Flag("environment").Value.String()))
			org := mode.OrgGet()
			utils.PrettyPrint(org)
		},
	}
	return cmd
}
