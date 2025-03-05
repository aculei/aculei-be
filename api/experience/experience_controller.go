package experience

import (
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/micheledinelli/aculei-be/utils"
	"github.com/rs/zerolog"
)

type ExperienceController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	experienceService    *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	experienceService *Service,
) {
	controller := newExperienceController(
		routes,
		configuration,
		experienceService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newExperienceController(routes models.DefaultRoutes,
	configuration models.Configuration,
	experienceService *Service,
) *ExperienceController {
	controllerLogger := utils.InitServiceAdvancedLogger("ExperienceController")

	return &ExperienceController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		experienceService:    experienceService,
	}
}

func (c *ExperienceController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.GET(
			"experience/random",
			c.getRandomExperienceImage(),
		)
	}
}

// getRandomExperienceImage godoc
// @Tags experience
// @Schemes http
// @Router /v1/experience/random [get]
// @Summary Returns a random image to be displayed in the experience page
// @Description Returns a random image. Randomness is achieved using sample aggregation in MongoDB.
// @Accept json
// @Produce json
// @Success 200 {object} models.AculeiImage "The random image"
// @Failure 500 {object} models.ErrorResponseModel "An error occurred"
func (c *ExperienceController) getRandomExperienceImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var randomImage *models.AculeiImage
		var err error

		randomImage, err = c.experienceService.GetRandomExperienceImage(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting random image")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, randomImage)
	}
}
