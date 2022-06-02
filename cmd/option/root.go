/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package option

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra_demo/pkg/introduce"
	"github.com/spf13/cobra_demo/pkg/version"
)

func NewCobraDemoCommand() *cobra.Command{
	cmds:=&cobra.Command{
		Use:   "cobra_demo",
		Short: "cobra_demo is my first study cobra project",
		Long: `cobra is used in many famous opensource frame,such as kubernetes,hugo,openyurt.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cmds.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmds.AddCommand(version.NewVersionCmd())
	cmds.AddCommand(introduce.NewIntroduceCmd())
	return cmds
}



