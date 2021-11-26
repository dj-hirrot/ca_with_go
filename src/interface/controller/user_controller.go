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

// Show godoc
// @Summary      Show an user
// @Description  Get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int           true  "User ID"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{id} [get]
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

// Index godoc
// @Summary      List users
// @Description  Get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Users
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users [get]
func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// Create godoc
// @Summary      Create user
// @Description  Create user by body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        parameter body      domain.User true "User attributes"
// @Success      201       {object}  domain.User
// @Failure      400       {object}  Error
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /users [post]
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

// Update godoc
// @Summary      Update user
// @Description  Update user by body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id        path      int         true "User ID"
// @Param        parameter body      domain.User true "User attributes"
// @Success      204       {object}  domain.User
// @Failure      400       {object}  Error
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /users/{id} [put]
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

// Delete godoc
// @Summary      Delete user
// @Description  Delete user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int           true  "User ID"
// @Success      201  {object}  domain.User
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{id} [delete]
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
