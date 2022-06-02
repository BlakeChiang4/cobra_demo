package main

import (
	"github.com/spf13/cobra_demo/cmd/option"
	"os"
)

func main(){
	cmd:=option.NewCobraDemoCommand()
	if err:=cmd.Execute();err!=nil{
		os.Exit(1)
	}
}
