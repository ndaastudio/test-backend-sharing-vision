package main

import (
	db_config "test-backend-sharing-vision/config"
	"test-backend-sharing-vision/handler"
	"test-backend-sharing-vision/repository"
	"test-backend-sharing-vision/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db_config.InitDB()
	db := db_config.GetDB()

	repo := repository.NewArticleRepository(db)
	svc := service.NewArticleService(repo)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	handler.NewArticleHandler(e, svc)

	e.Logger.Fatal(e.Start(":4000"))
}
