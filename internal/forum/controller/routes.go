package controller

import (
	"net/http"

	"github.com/ynuraddi/router"
)

func routes() http.Handler {
	r := router.NewRouter()

	r.GET("/", nil)
	r.POST("/sign-up", nil)
	r.POST("/login", nil)
	r.POST("/logout", nil)

	r.GET("/community", nil)
	r.POST("/community/subscribe", nil)

	r.POST("/community/create-post", nil)
	r.PATCH("/community/update-post", nil)
	r.DELETE("/community/delete-post", nil)

	r.POST("/community/comment-post", nil)
	r.POST("/community/vote-post", nil)

	return r
}
