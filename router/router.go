package router

import (
	"apiserver/handler/sd"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 你懂的~")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)

	}

	return g
}
