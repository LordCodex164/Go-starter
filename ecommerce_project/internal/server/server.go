package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordcodex164/ecommerce_project/internal/config"
	"github.com/rs/zerolog"
)

type Server struct {
	logger *zerolog.Logger
	config *config.Config
}

func New(logger *zerolog.Logger, config *config.Config) *Server {
	return &Server{
		logger: logger,
		config: config,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	//add middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	router.GET("/health", s.healthCheck)

	s.logger.Info().Msg("routes successfully setup")

	return router

}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

}