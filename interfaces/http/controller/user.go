package controllers

import (
	"fmt"
	appsvc "go-ddd/application/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	app *appsvc.Service
}

func NewUserController(app *appsvc.Service) *UserController {
	return &UserController{
		app: app,
	}
}

func (c *UserController) Index(ctx echo.Context) error {
	res, err := c.app.GetAll(ctx.Request().Context())
	fmt.Println(res)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	users := make([]appsvc.UserDTO, 0, len(res))

	users = append(users, res...)

	return ctx.JSON(http.StatusOK, users)
}

// func (c *UserController) Get(ctx echo.Context) error {

// }

// func (c *UserController) Post(ctx echo.Context) error {

// }

// func (c *UserController) Put(ctx echo.Context) error {

// }

// func (c *UserController) Delete(ctx echo.Context) error {

// }