package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	restCmd = &cobra.Command{
		Use:   "rest",
		Short: "dummy-server",
		Run:   restServer,
	}
)

func init() {
	rootCmd.AddCommand(restCmd)
}
func restServer(cmd *cobra.Command, args []string) {
	var e = echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowMethods: []string{echo.GET, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//websocket.HandlerWsInit(e, webSocketHandler)

	err := e.Start(rootConfig.Server.HostServer + ":" + rootConfig.Server.PortServer)
	if err != nil {
		log.Errorf("Cannot Start the application !!, Err > ", err)
		os.Exit(1)
	}

}
