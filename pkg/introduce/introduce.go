package introduce

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewIntroduceCmd() *cobra.Command{
	io:=NewIntroduceOptions()

	cmd:=&cobra.Command{
		Use:   "introduce",
		Short: "Print the information of yourself",
		Long:  `Introduce yourself about your name , sex, age and hometown.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("My name is %s,sex is %s,age is %d,hometown is %s,father is %s.",
				io.name,io.sex,io.age,io.hometown,io.father)
		},
	}

	io.AddFlags(cmd.Flags())
	return cmd
}
