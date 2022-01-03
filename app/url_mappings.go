package app

import (
	users "bookstore/bookstore_users-api/controllers"
	"bookstore/bookstore_users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
