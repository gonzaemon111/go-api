package controllers

import (
    "app/models"
    "app/interfaces/database"
		"app/usecase"
		"strconv"
)


type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
			Interactor: usecase.UserInteractor{
					UserRepository: &database.UserRepository{
							SqlHandler: sqlHandler,
					},
			},
	}
}

func (controller *UserController) Create(c Context) {
	user := models.User{}
	c.Bind(&user)
	new_user, err := controller.Interactor.Add(user)
	if err != nil {
			c.JSON(500, "Internal Server Error")
			return
	}
	c.JSON(201, new_user)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
			c.JSON(500, "Internal Server Error")
			return
	}
	c.JSON(200, user)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, users)
}
