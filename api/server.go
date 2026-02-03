package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/api/archive"
	"github.com/micheledinelli/aculei-be/api/experience"
	"github.com/micheledinelli/aculei-be/api/filters"
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
	filtersService    *filters.Service
}

func NewServer(
	configuration models.Configuration,
	archiveService *archive.Service,
	experienceService *experience.Service,
	filtersService *filters.Service,
) *Server {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowHeaders:     configuration.CORS.AllowHeaders,
		AllowOrigins:     configuration.CORS.AllowOrigins,
		AllowMethods:     []string{ "GET", "OPTIONS", "HEAD"},
		AllowCredentials: true,
	}))
	
	engine.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	server := &Server{
		engine:            engine,
		configuration:     configuration,
		archiveService:    archiveService,
		experienceService: experienceService,
		filtersService:    filtersService,
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
	filters.InjectRoutes(routes, s.configuration, s.filtersService)

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
