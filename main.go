package main

import (
	"be-sharing-vision/app"
	"be-sharing-vision/app/controller"
	"be-sharing-vision/app/repository"
	"be-sharing-vision/app/usecase"
	"be-sharing-vision/config"
	"be-sharing-vision/driver"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "min":
					report.Message = fmt.Sprintf("%s value must be minimal %s character", err.Field(), err.Param())
				case "max":
					report.Message = fmt.Sprintf("%s value must be maximal %s character", err.Field(), err.Param())

				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	if err != nil {
		e.Logger.Fatal(err)
		log.Fatal("Error loading .env file")
	}
	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
		fmt.Println("error")
	}
	db := driver.InitDB(cfg.DBUSER, cfg.DBPASSWORD, cfg.HOSTDB, cfg.DBPORT, cfg.DBName)
	timeout := time.Duration(cfg.SERVER_TIMEOUT) * time.Second

	//init product
	postRepo := repository.NewPostRepositoryImpl(db)
	postUseCase := usecase.NewPostUsecase(postRepo, timeout)
	postControllerInit := controller.NewPostController(*postUseCase)
	routesInit := app.RouteParams{
		PostController: postControllerInit,
	}
	routesInit.Routes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.APPPORT)))
}
