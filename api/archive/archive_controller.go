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
			c.getArchive(),
		)

		v1.GET(
			"archive/image/:id",
			c.getArchiveImage(),
		)
	}
}

// getArchive godoc
// @Tags archive
// @Schemes http
// @Router /v1/archive [get]
// @Param page query int false "page index starting from 0"
// @Param size query int false "number of items per page"
// @Param sortBy query string false "key to sort by" Enums(date,cam,animal,temperature,moon_phase) default(date)
// @Param animal query 			[]string false "list of animals" collectionFormat(multi)
// @Param moon_phase query 		[]string false "list of moon phases" collectionFormat(multi)
// @Param temperature query 	[]int false "list of temperatures" collectionFormat(multi)
// @Param date query 			[]string false "list of dates" collectionFormat(multi)
// @Summary Returns a paginated response with the list of archive images
// @Description Return the list of all the archive images with their metadata. The response is paginated.
// @Accept json
// @Produce json
// @Success 200 {object} models.PaginatedResponseModel[models.AculeiImage] "The list of archive images with pagination metadata"
// @Failure 400 {object} models.ErrorResponseModel "Bad request"
// @Failure 500 {object} models.ErrorResponseModel "An error occurred"
func (c *ArchiveController) getArchive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var archive *[]models.AculeiImage
		var archiveCount int
		var err error

		page := ctx.Query("page")
		size := ctx.Query("size")
		sortBy := ctx.Query("sortBy")

		fg, err := models.BuildFilterGroup(ctx)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error building filters")
			if _, ok := err.(*models.ErrorFilter); ok {
				ctx.JSON(400, models.NewBadRequest(err.Error()))
			}

			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		archiveCount, err = c.archiveService.GetArchiveCount(ctx, *fg)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive list count")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		paginator := models.NewPaginator(page, size, archiveCount, sortBy)

		archive, err = c.archiveService.GetArchive(ctx, *paginator, *fg)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive list")
			if _, ok := err.(*models.ErrorFilter); ok {
				ctx.JSON(400, models.NewBadRequest(err.Error()))
				return
			}

			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		var next *int
		if (paginator.Page+1)*paginator.Size < archiveCount {
			nextVal := paginator.Page + 1
			next = &nextVal
		}

		response := models.PaginatedResponseModel[models.AculeiImage]{
			SortBy: paginator.SortBy.String(),
			Page:   paginator.Page,
			Size:   paginator.Size,
			Next:   next,
			Data:   *archive,
			Total:  archiveCount,
			Count:  len(*archive),
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
		var img *models.AculeiImage
		var err error

		id := ctx.Param("id")

		img, err = c.archiveService.GetArchiveImage(ctx, id)
		if err != nil {
			c.logger.Error().Err(err).Msg("Error getting archive image")
			ctx.JSON(500, models.ErrorInternalServerErrorResponseModel)
			return
		}

		ctx.JSON(200, img)
	}
}
