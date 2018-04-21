package controllers

import (
	"app/interfaces/database"
	"app/models"
	"strconv"

	"app/usecase"
)

// UserController manages responce
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController genrerates the instance of UserController
func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Create is savig a new user
func (controller *UserController) Create(c Context) {
	user := models.User{}
	c.Bind(&user)
	newUser, err := controller.Interactor.Add(user)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(201, newUser)
}

// Show is getting a user by id
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, user)
}

// Index is getting all users
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, users)
}
