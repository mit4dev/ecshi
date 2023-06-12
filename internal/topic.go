package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mit4dev/ecshi/domain"
)

func (api *api) getTopics(c *gin.Context) {
	topics, err := api.topicRepo.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (api *api) getTopic(c *gin.Context) {
	user, err := api.topicRepo.FindById(c.Request.Context(), c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (api *api) postTopic(c *gin.Context) {
	var topic domain.Topic

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inserted, err := api.topicRepo.Store(c.Request.Context(), &topic)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inserted)
}
