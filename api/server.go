package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/api/dataset"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine         *gin.Engine
	configuration  models.Configuration
	datasetService *dataset.Service
}

func NewServer(
	configuration models.Configuration,
	datasetService *dataset.Service,
) *Server {
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowHeaders:     configuration.CORS.AllowHeaders,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowOrigins:     configuration.CORS.AllowOrigins,
		AllowCredentials: true,
	}))

	server := &Server{
		engine:         gin.Default(),
		configuration:  configuration,
		datasetService: datasetService,
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	unauthenticatedRoute := s.engine.Group("/")
	// internalRoute := s.engine.Group("/")

	unauthenticatedRoute.GET("/health", s.createHealthRoute())

	routes := models.DefaultRoutes{
		UnauthenticatedRoute: unauthenticatedRoute,
		// AuthenticatedRoute:   authenticatedRoute,
		// InternalRoute:        internalRoute,
	}

	dataset.InjectRoutes(routes, s.configuration, s.datasetService)

	if s.configuration.Environment == models.Development {
		log.Info().Msgf("Enable swagger on http://%s:%d/swagger/index.html", s.configuration.HTTPHost, s.configuration.HTTPPort)
		s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (s *Server) createHealthRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}

func (s *Server) Listen() error {
	address := fmt.Sprintf("%s:%d", s.configuration.HTTPHost, s.configuration.HTTPPort)

	log.Info().Msgf("Listening on %s", address)
	return s.engine.Run(address)
}
