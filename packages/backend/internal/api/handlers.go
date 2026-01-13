package api

import (
	"backend/internal/scraper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	anime, err := scraper.GetLatestAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func GetOngoing(c *gin.Context) {
	anime, err := scraper.GetOngoingAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func GetPopular(c *gin.Context) {
	anime, err := scraper.GetPopularAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}
	anime, err := scraper.SearchAnime(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func GetDetail(c *gin.Context) {
	id := c.Param("id")
	detail, err := scraper.GetAnimeDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detail)
}

func GetEpisode(c *gin.Context) {
	id := c.Param("id")
	episode, err := scraper.GetEpisodeDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, episode)
}

func GetGenres(c *gin.Context) {
	genres, err := scraper.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genres)
}

func GetServerLink(c *gin.Context) {
	id := c.Param("id")
	url, err := scraper.GetServerURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url)
}
