package dataset

import (
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/micheledinelli/aculei-be/utils"
	"github.com/rs/zerolog"
)

type DatasetController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	datasetService       *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	datasetService *Service,
) {
	controller := newDatasetController(
		routes,
		configuration,
		datasetService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newDatasetController(routes models.DefaultRoutes,
	configuration models.Configuration,
	datasetService *Service,
) *DatasetController {
	controllerLogger := utils.InitServiceAdvancedLogger("AlbumController")

	return &DatasetController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		datasetService:       datasetService,
	}
}

func (c *DatasetController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.GET(
			"dataset/",
			c.getDataset(),
		)
	}
}

// getDataset godoc
// @Tags dataset
// @Schemes https
// @Router /v1/dataset [get]
// @Summary return dataset info
// @Description Return info about dataset like number of records, number of columns, etc.
// @Accept json
// @Produce json
// @Success 200 {object} models.Dataset "The dataset info"
// @Failure 500 {object} models.ErrorInternalServerError "An error occurred"
func (c *DatasetController) getDataset() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataset *models.Dataset
		var err error

		dataset, err = c.datasetService.GetDataset()
		if err != nil {
			ctx.JSON(500, models.ErrorInternalServerError{Message: "An error occurred"})
			return
		}

		ctx.JSON(200, dataset)
	}
}
