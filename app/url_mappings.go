package app

import (
	"bookstore/bookstore_users-api/controllers/ping"
	"bookstore/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
