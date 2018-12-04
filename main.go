package main

import (
	"github.com/kooksee/dmeta/cmds"
	"github.com/kooksee/dmeta/internal/utils"
)

func main() {

	rootCmd := cmds.RootCmd()
	rootCmd.AddCommand(
		cmds.VersionCmd(),
		cmds.ApiServerCmd(),
	)
	utils.MustNotError(rootCmd.Execute())
}
