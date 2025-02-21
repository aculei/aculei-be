package archive

import (
	"github.com/gin-gonic/gin"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/micheledinelli/aculei-be/utils"
	"github.com/rs/zerolog"
)

type ArchiveController struct {
	authenticatedRoute   *gin.RouterGroup
	internalRoute        *gin.RouterGroup
	unauthenticatedRoute *gin.RouterGroup
	logger               *zerolog.Logger
	configuration        models.Configuration
	archiveService       *Service
}

func InjectRoutes(
	routes models.DefaultRoutes,
	configuration models.Configuration,
	archiveService *Service,
) {
	controller := newArchiveController(
		routes,
		configuration,
		archiveService,
	)

	controller.injectUnAuthenticatedRoutes()
}

func newArchiveController(routes models.DefaultRoutes,
	configuration models.Configuration,
	archiveService *Service,
) *ArchiveController {
	controllerLogger := utils.InitServiceAdvancedLogger("ArchiveController")

	return &ArchiveController{
		unauthenticatedRoute: routes.UnauthenticatedRoute,
		authenticatedRoute:   routes.AuthenticatedRoute,
		internalRoute:        routes.InternalRoute,
		configuration:        configuration,
		logger:               controllerLogger,
		archiveService:       archiveService,
	}
}

func (c *ArchiveController) injectUnAuthenticatedRoutes() {
	v1 := c.unauthenticatedRoute.Group("v1")

	{
		v1.GET(
			"archive",
			c.getArchiveList(),
		)

		v1.GET(
			"archive/image/:id",
			c.getArchiveImage(),
		)
	}
}

// getArchiveList godoc
// @Tags archive
// @Schemes http
// @Router /v1/archive [get]
// @Param page query int false "page index starting from 0"
// @Param size query int false "number of items per page"
// @Summary Returns a paginated response with the list of archive images
// @Description Return the list of all the archive images with their metadata. The response is paginated.
// @Accept json
// @Produce json
// @Success 200 {object} models.PaginatedResponseModel[models.AculeiImage] "The list of archive images"
// @Failure 500 {object} models.ErrorResponseModel "An error occurred"
func (c *ArchiveController) getArchiveList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var archiveList *[]models.AculeiImage
		var archiveCount int
		var err error

		page := ctx.Query("page")
		size := ctx.Query("size")

		archiveCount, err = c.archiveService.GetArchiveListCount(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive list count")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		paginator := models.NewPaginator(page, size, archiveCount)

		archiveList, err = c.archiveService.GetArchiveList(ctx, *paginator)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive list")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		var next *int
		if (paginator.Page+1)*paginator.Size < archiveCount {
			nextVal := paginator.Page + 1
			next = &nextVal
		}

		response := models.PaginatedResponseModel[models.AculeiImage]{
			Page:  paginator.Page,
			Size:  paginator.Size,
			Next:  next,
			Data:  *archiveList,
			Total: archiveCount,
			Count: len(*archiveList),
		}

		ctx.JSON(200, response)
	}
}

// getArchiveImage godoc
// @Tags archive
// @Schemes http
// @Router /v1/archive/image/{id} [get]
// @Param id path string true "the archive image id"
// @Summary Returns a single archive image
// @Description Returns a single archive with its metadata
// @Accept json
// @Produce json
// @Success 200 {object} models.AculeiImage "The archive image and its metadata"
// @Failure 500 {object} models.ErrorResponseModel "An error occurred"
func (c *ArchiveController) getArchiveImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var archiveImage *models.AculeiImage
		var err error

		id := ctx.Param("id")

		archiveImage, err = c.archiveService.GetArchiveImage(ctx, id)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive image")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, archiveImage)
	}
}
