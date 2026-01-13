package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.GET("/home", GetHome)
		api.GET("/ongoing", GetOngoing)
		api.GET("/popular", GetPopular)
		api.GET("/search", Search)
		api.GET("/anime/:id", GetDetail)
		api.GET("/episode/:id", GetEpisode)
		api.GET("/genres", GetGenres)
		api.GET("/server/:id", GetServerLink)
	}

	return r
}
