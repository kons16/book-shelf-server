package web

import (
	"github.com/kons16/book-shelf-server/usecase"
	"github.com/kons16/book-shelf-server/web/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userUC *usecase.UserUseCase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	v1 := e.Group("/api/v1")

	userHandler := handler.NewUserHandler(userUC)

	v1.POST("/user", userHandler.Create)

	return e
}
