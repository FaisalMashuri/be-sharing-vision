package controller

import (
	"be-sharing-vision/app/domain"
	"be-sharing-vision/app/request"
	"be-sharing-vision/app/response"
	"be-sharing-vision/app/usecase"
	"strconv"
	"strings"

	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PostController struct {
	UseCase usecase.PostUsecase
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewPostController(u usecase.PostUsecase) *PostController {
	return &PostController{
		UseCase: u,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (p *PostController) CreatePost(c echo.Context) error {
	ctx := c.Request().Context()
	createdPost := request.PostRequest{}
	c.Bind(&createdPost)
	if err := c.Validate(createdPost); err != nil {
		return err
	}

	status := strings.Title(createdPost.Status)
	if status == "Publish" || status == "Draft" || status == "Trash" {
		createdPost.Status = status
	} else {
		defer c.Logger().Error(fmt.Errorf("status tidak sesuai"))
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("status tidak sesuai harus pilih salah satu dari (Publish, Draft,Trash)"))
	}
	PostDomain := domain.PostDomain{

		Title:    createdPost.Title,
		Content:  createdPost.Content,
		Category: createdPost.Category,
		Status:   createdPost.Status,
	}

	post, err := p.UseCase.CreatePost(ctx, PostDomain)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, post)
}

func (p *PostController) GetAllPost(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	post, err := p.UseCase.GetAllPost(ctx, offset, limit)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(200, post)
}

func (p *PostController) GetByID(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	id := c.Param("id")
	product, err := p.UseCase.GetPostById(ctx, id)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, product)
}

func (p *PostController) UpdatePost(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	id := c.Param("id")
	createdPost := request.PostRequest{}
	c.Bind(&createdPost)
	if err := c.Validate(createdPost); err != nil {
		return err
	}

	status := strings.Title(createdPost.Status)
	if status == "Publish" || status == "Draft" || status == "Trash" {
		createdPost.Status = status
	} else {
		defer c.Logger().Error(fmt.Errorf("status tidak sesuai"))
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("status tidak sesuai harus pilih salah satu dari (Publish, Draft,Trash)"))
	}
	PostDomain := domain.PostDomain{

		Title:    createdPost.Title,
		Content:  createdPost.Content,
		Category: createdPost.Category,
		Status:   createdPost.Status,
	}

	post, err := p.UseCase.UpdatePost(ctx, id, PostDomain)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, post)
}

func (p *PostController) DeletePost(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	id := c.Param("id")
	if err := p.UseCase.DeletePost(ctx, id); err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(200, "success")
}
