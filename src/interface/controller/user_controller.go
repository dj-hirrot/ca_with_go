package controller

import (
	"strconv"

	"github.com/dj-hirrot/ca_with_go/src/domain"
	"github.com/dj-hirrot/ca_with_go/src/interface/db"
	"github.com/dj-hirrot/ca_with_go/src/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler db.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &db.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.CreateUser(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.UpdateUser(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		Id: id,
	}
	err = controller.Interactor.DeleteUser(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
