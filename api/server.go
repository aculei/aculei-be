package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/api/archive"
	"github.com/micheledinelli/aculei-be/api/experience"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine            *gin.Engine
	configuration     models.Configuration
	archiveService    *archive.Service
	experienceService *experience.Service
}

func NewServer(
	configuration models.Configuration,
	archiveService *archive.Service,
	experienceService *experience.Service,
) *Server {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowHeaders:     configuration.CORS.AllowHeaders,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowOrigins:     configuration.CORS.AllowOrigins,
		AllowCredentials: true,
	}))

	server := &Server{
		engine:            engine,
		configuration:     configuration,
		archiveService:    archiveService,
		experienceService: experienceService,
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

	archive.InjectRoutes(routes, s.configuration, s.archiveService)
	experience.InjectRoutes(routes, s.configuration, s.experienceService)

	if s.configuration.Environment == models.Development {
		log.Info().Msgf("Enabled swagger on http://%s:%d/swagger/index.html", s.configuration.HTTPHost, s.configuration.HTTPPort)
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
