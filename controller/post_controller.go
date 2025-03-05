package controller

import (
	"net/http"
	"sandbox/config"
	"sandbox/db/models"
	"sandbox/db/service"
	"sandbox/lib/server"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostController struct {
	authConfig  *config.AuthConfig
	postService *service.PostService
}

func NewPostController(authConfig *config.AuthConfig, postService *service.PostService) *PostController {
	return &PostController{
		authConfig:  authConfig,
		postService: postService,
	}
}

func (c *PostController) CreatePost(ctx echo.Context) error {
	user, err := server.CurrentUser(ctx, c.authConfig.AuthUserContextKey)
	if err != nil {
		return err
	}

	var postCreateRequestBody CreateUpdatePostRequestBody
	if err := ctx.Bind(&postCreateRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := postCreateRequestBody.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	postCreateParams := models.PostCreateParams{
		ID:          uuid.New(),
		Title:       postCreateRequestBody.Title,
		Content:     postCreateRequestBody.Body,
		CreatedByID: user.Id,
	}

	newPost, err := c.postService.CreatePost(ctx.Request().Context(), postCreateParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := PostResponseFromModel(newPost)
	return ctx.JSON(http.StatusOK, resp)
}

func (c *PostController) UpdatePost(ctx echo.Context) error {
	postId := ctx.Param("id")
	if postId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	// TODO: extract parseUUID to function.
	parsedPostId, err := uuid.Parse(postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	var postCreateRequestBody CreateUpdatePostRequestBody
	if err := ctx.Bind(&postCreateRequestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := postCreateRequestBody.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	postUpdateParams := models.PostUpdateParams{
		ID:      parsedPostId,
		Title:   postCreateRequestBody.Title,
		Content: postCreateRequestBody.Body,
	}

	updatedPost, err := c.postService.UpdatePost(ctx.Request().Context(), postUpdateParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := PostResponseFromModel(updatedPost)
	return ctx.JSON(http.StatusOK, resp)
}

func (c *PostController) DeletePost(ctx echo.Context) error {
	postId := ctx.Param("id")
	if postId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	parsedPostId, err := uuid.Parse(postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	err = c.postService.DeletePost(ctx.Request().Context(), parsedPostId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, server.MessageResponse{
		Message: "post deleted successfully",
	})
}

func (c *PostController) ListPosts(ctx echo.Context) error {
	// parse limit query param
	var limit int32 = 10
	limitString := ctx.QueryParam("limit")
	if limitString != "" {
		parsedLimit, err := strconv.Atoi(limitString)
		if err == nil {
			//nolint:gosec
			limit = int32(parsedLimit)
		}
	}

	// TODO: extract to function.
	// parse offset query param
	var offset int32 = 0
	offsetString := ctx.QueryParam("offset")
	if offsetString != "" {
		parsedOffset, err := strconv.Atoi(offsetString)
		if err == nil {
			//nolint:gosec
			offset = int32(parsedOffset)
		}
	}

	postsResult, err := c.postService.ListPosts(ctx.Request().Context(), models.PostsListParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := ListPostsResponseFromListPostsResult(postsResult)
	return ctx.JSON(http.StatusOK, resp)
}

func (c *PostController) GetPostById(ctx echo.Context) error {
	postId := ctx.Param("id")
	if postId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	parsedPostId, err := uuid.Parse(postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
	}

	post, err := c.postService.GetPostById(ctx.Request().Context(), parsedPostId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := PostWithCreatedByResponseFromPostById(post)
	return ctx.JSON(http.StatusOK, resp)
}
