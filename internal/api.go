package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mit4dev/ecshi/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	gin       *gin.Engine
	userRepo  domain.UserRepository
	topicRepo domain.TopicRepository
	jwtRepo   domain.JwtRepository
}

func NewApi(db *mongo.Database, gin *gin.Engine) *api {
	userRepo := domain.NewUserRepository(db.Collection("users"))
	topicRepo := domain.NewTopicRepository(db.Collection("topics"))
	jwtRepo := domain.NewJwtRepository(db.Collection("jwts"))

	return &api{userRepo: userRepo, topicRepo: topicRepo, jwtRepo: jwtRepo, gin: gin}
}

func (api *api) RegisterRoutes() {
	api.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := api.gin.Group("/auth")
	auth.POST("/login", api.login)
	auth.POST("/refresh", api.refresh)
	auth.POST("/logout", api.logout)

	users := api.gin.Group("/users")
	users.GET("/:id", api.getUser)
	users.GET("/", api.getUsers)
	users.POST("/", api.postUser)
	users.POST("/register", api.postUser)

	topics := api.gin.Group("/topics")
	topics.Use(api.authGuard)
	topics.GET("/:id", api.getTopic)
	topics.GET("/", api.getTopics)
	topics.POST("/", api.postTopic)

}
