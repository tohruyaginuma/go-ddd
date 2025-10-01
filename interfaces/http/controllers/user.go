package userController

import (
	appsvc "go-ddd/application/user"
	interfaceHttp "go-ddd/interfaces/http"
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

func (uc *UserController) Index(c echo.Context) error {
	ctx := c.Request().Context()
	list, err := uc.app.GetAll(ctx)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get users"})
	}
	
	res := make([]interfaceHttp.UserResponse, 0, len(list))
	for _, u := range list {
		res = append(res, interfaceHttp.UserResponse{ID: u.ID, Name: u.Name})
	}

	return c.JSON(http.StatusOK, interfaceHttp.UserIndexResponse{Users: res})
}

func (uc *UserController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	res, err := uc.app.Get(ctx, appsvc.UserGetCommand{ID: id})

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "failed to get user"})
	}

	return c.JSON(http.StatusOK, interfaceHttp.UserGetResponse{
		User: interfaceHttp.UserResponse{ID: res.ID, Name:res.Name},
	})

}

func (uc *UserController) Post(c echo.Context) error {
	ctx := c.Request().Context()
	
	var req interfaceHttp.UserPostRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.UserName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "userName is required"})
	}

	if err := uc.app.Register(ctx, appsvc.UserRegisterCommand{Name: req.UserName}); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (uc *UserController) Put(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	var req interfaceHttp.UserPutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}

	_, err := uc.app.Update(ctx, appsvc.UserUpdateCommand{
		ID:   id,
		Name: req.Name,
	})
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if err := uc.app.Delete(ctx, appsvc.UserDeleteCommand{ID: id}); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}