package app

import (
	pingController "bookstore-users-api/controllers/ping"
	userController "bookstore-users-api/controllers/user"
)

func mapURLS() {
	router.GET("/ping", pingController.Ping)

	router.GET("/users/:user_id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:user_id", userController.UpdateUser)
	router.PATCH("/users/:user_id", userController.UpdateUser)
	router.DELETE("/users/:user_id", userController.DeleteUser)

	router.GET("/internal/users/search", userController.Search) //example: [URL]?status=Active
}
