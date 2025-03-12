package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/micheledinelli/aculei-be/utils"
	"github.com/rs/zerolog"
)

type FiltersController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	filtersService       *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	filtersService *Service,
) {
	controller := newFiltersController(
		routes,
		configuration,
		filtersService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newFiltersController(routes models.DefaultRoutes,
	configuration models.Configuration,
	filtersService *Service,
) *FiltersController {
	controllerLogger := utils.InitServiceAdvancedLogger("FiltersController")

	return &FiltersController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		filtersService:       filtersService,
	}
}

func (c *FiltersController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.GET(
			"filters",
			c.getFilters(),
		)
	}
}

// getFilters godoc
// @Tags filters
// @Schemes http
// @Router /v1/filters [get]
// @Summary Returns the list of available filters
// @Description Returns the list of available filters
// @Accept json
// @Produce json
// @Success 200 {array} models.Filter "The list of available filters"
// @Failure 500 {object} models.ErrorResponseModel "An error occurred"
func (c *FiltersController) getFilters() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filters *[]models.Filter
		var err error

		filters, err = c.filtersService.GetFilters(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting filters")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, filters)
	}
}
