package cmds

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/dmeta/internal/api"
	"github.com/kooksee/dmeta/internal/config"
	"github.com/spf13/cobra"
)

func ApiServerCmd() *cobra.Command {
	var addr = ":8080"

	var handleArgs = func(cmd *cobra.Command) *cobra.Command {
		cmd.PersistentFlags().StringVar(&addr, "addr", addr, "task proxy addr")
		return cmd
	}

	return handleArgs(&cobra.Command{
		Use:   "s",
		Short: "dmeta proxy server",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.DefaultConfig()
			cfg.Init()

			gin.SetMode(gin.ReleaseMode)
			if cfg.IsDebug() {
				gin.SetMode(gin.DebugMode)
			}

			r := gin.Default()
			r.POST("/object", api.PutObject)
			r.GET("/object/:id", api.GetObject)
			r.POST("/metadata", api.PutMetadata)
			r.GET("/metadata/:id", api.GetMetadata)
			r.GET("/range", api.Ranger)

			return r.Run(addr)
		},
	})
}
