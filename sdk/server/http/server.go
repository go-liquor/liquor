package http

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-liquor/liquor/sdk/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func instanceServer(config *config.Config) *gin.Engine {
	var svc *gin.Engine
	if config.IsDebug() {
		gin.SetMode(gin.DebugMode)
		svc = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		svc = gin.New()
	}
	crs := cors.Default()
	if !config.GetServerHttpCorsDefaultAllow() {
		crs = cors.New(cors.Config{
			AllowOrigins:     config.GetServerHttpCorsAllowOrigins(),
			AllowMethods:     config.GetServerHttpCorsAllowMethods(),
			AllowHeaders:     config.GetServerHttpCorsAllowHeaders(),
			AllowCredentials: config.GetServerHttpCorsAllowCredentials(),
		})
	}
	svc.Use(crs)
	return svc
}

func startServer(config *config.Config, server *gin.Engine, lg *zap.Logger, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lg.Info("starting HTTP server", zap.Int64("port", config.GetServerHttpPort()))
			go server.Run(fmt.Sprintf(":%d", config.GetServerHttpPort()))
			return nil
		},
		OnStop: func(context.Context) error {
			lg.Info("stopping HTTP server")
			return nil
		},
	})
}

func initialRoute(server *gin.Engine) {
	server.GET("/-/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
