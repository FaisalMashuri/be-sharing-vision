package app

import (
	"be-sharing-vision/app/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RouteParams struct {
	PostController *controller.PostController
}

func (p RouteParams) Routes(e *echo.Echo) {
	e.GET("/article", func(c echo.Context) error {
		return c.String(http.StatusOK, "hai")
	})
	e.POST("/article", p.PostController.CreatePost)
	e.GET("/article/:offset/:limit", p.PostController.GetAllPost)
	e.GET("/article/:id", p.PostController.GetByID)
	e.PUT("/article/:id", p.PostController.UpdatePost)
	e.DELETE("/article/:id", p.PostController.DeletePost)
}
