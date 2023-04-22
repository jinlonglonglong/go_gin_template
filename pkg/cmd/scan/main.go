package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"net/http"
	"scan/pkg/api"
	"scan/pkg/frame"
	"scan/pkg/middleware"
	"scan/pkg/ws"
)

type Option struct {
}

func main() {
	defer glog.Flush()

	frame.InitFramework()
	defer frame.UnInitFramework()

	service := NewService(Option{})
	service.Init()
	defer service.UnInit()

	debug := viper.GetBool("application.debug")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	//r.Use(static.Serve("/", static.LocalFile("./bin/dist", true)))

	r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())

	api.Setup(r)

	r.StaticFile("/", "./static/index.html")

	wsh, err := ws.NewWebSocketHandler()
	if err != nil {
		panic(err.Error())
	}
	ws.Setup(r, wsh)

	go ws.BroadcastChannel.ProcessTask()

	addr := viper.GetString("server.addr")

	if len(addr) != 0 {
		err = r.Run(addr)
	} else {
		err = r.Run()
	}

	if err != nil {
		panic(err.Error())
	}
}

func recoverHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, err)
}
