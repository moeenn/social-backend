package controller

import (
	"net/http"
	"sandbox/config"
	"sandbox/db/models"
	"sandbox/db/service"
	"sandbox/lib/server"

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
