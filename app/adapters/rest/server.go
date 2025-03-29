package rest

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-liquor/liquor/v2/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func instanceServer(cfg *config.Config) *gin.Engine {
	var svc *gin.Engine
	if cfg.GetBool(config.AppDebug) && !cfg.GetBool(config.RestDisabled) {
		gin.SetMode(gin.DebugMode)
		svc = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		svc = gin.New()
	}
	if !cfg.GetBool(config.RestDisabled) {
		crs := cors.Default()
		if !cfg.GetBool(config.CorsDefault) {
			corsConfig := cors.Config{
				AllowMethods:     cfg.GetStringSlice(config.CorsAllowMethods),
				AllowHeaders:     cfg.GetStringSlice(config.CorsAllowHeaders),
				AllowCredentials: cfg.GetBool(config.CorsAllowCredentials),
			}

			if len(cfg.GetStringSlice(config.CorsAllowOrigins)) == 1 && cfg.GetStringSlice(config.CorsAllowOrigins)[0] == "*" {
				corsConfig.AllowAllOrigins = true
			} else {
				corsConfig.AllowOrigins = cfg.GetStringSlice(config.CorsAllowOrigins)
			}

			crs = cors.New(corsConfig)
		}
		svc.Use(crs)
	}
	return svc
}

func startServer(cfg *config.Config, server *gin.Engine, lg *zap.Logger, lc fx.Lifecycle) {
	if cfg.GetBool(config.RestDisabled) {
		return
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lg.Info("starting HTTP server", zap.Int64("port", cfg.GetInt64(config.RestPort)))
			go server.Run(fmt.Sprintf(":%d", cfg.GetInt64(config.RestPort)))
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
