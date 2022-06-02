package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command{
	cmds:=&cobra.Command{
		Use:   "version",
		Short: "Print the version number of cobra_demo",
		Long:  `All software has versions. This is cobra_demo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cobra_demo v0.1 -- BlakeChiang")
		},
	}
	return cmds
}
