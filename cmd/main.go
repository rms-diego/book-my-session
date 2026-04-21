package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	s.ListenAndServe()
}
