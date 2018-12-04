package cmds

import (
	"fmt"
	"github.com/kooksee/dmeta/version"
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show Version Info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dmeta version:", version.Version)
			fmt.Println("dmeta commit version:", version.CommitVersion)
			fmt.Println("dmeta build version:", version.BuildVersion)
		},
	}
}
