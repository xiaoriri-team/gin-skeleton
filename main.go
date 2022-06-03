package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/xiaoriri-team/gin-skeleton/global"
	"github.com/xiaoriri-team/gin-skeleton/internal/routers"
	"github.com/xiaoriri-team/gin-skeleton/pkg/util"
)

var (
	version, buildDate, commitID string
)

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           global.ServerSetting.HttpIp + ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	util.PrintHelloBanner(fmt.Sprintf("%s %s (build:%s %s)", global.AppSetting.Name, version, commitID, buildDate))
	fmt.Fprintf(color.Output, "%s service listen on %s\n", global.AppSetting.Name,
		color.GreenString(fmt.Sprintf("http://%s:%s", global.ServerSetting.HttpIp, global.ServerSetting.HttpPort)),
	)
	s.ListenAndServe()
}
