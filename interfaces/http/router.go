package http

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, uc *UserHandler){
	v1 := e.Group("/v1")
	v1Users := v1.Group("/users")

	v1Users.GET("", uc.Index)
	v1Users.GET("/:id", uc.Get)
	v1Users.POST("", uc.Post)
	v1Users.PUT("/:id", uc.Put)
	v1Users.DELETE("/:id", uc.Delete)
}