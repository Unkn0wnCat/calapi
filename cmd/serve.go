package cmd

import (
	"github.com/Unkn0wnCat/calapi/internal/database"
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/Unkn0wnCat/calapi/internal/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the API server",
	//Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		err := database.Initialize()
		if err != nil {
			return err
		}
		defer database.Shutdown()

		go server.Serve()

		<-c // Wait for quit signal
		logger.Logger.Info("goodbye, shutting down...")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
