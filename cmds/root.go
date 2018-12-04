package cmds

import (
	"github.com/kooksee/dmeta/internal/config"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	var handle = func(cmd *cobra.Command) *cobra.Command {
		cfg := config.DefaultConfig()
		cmd.PersistentFlags().BoolVar(&cfg.Debug, "debug", cfg.Debug, "debug")
		return cmd
	}

	return handle(&cobra.Command{
		Use:   "dmeta",
		Short: "分布式任务处理系统",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			return nil
		},
	})
}
