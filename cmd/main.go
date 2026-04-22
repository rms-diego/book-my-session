package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/middleware"
	"github.com/rms-diego/book-my-session/internal/routes"
	"github.com/rms-diego/book-my-session/pkg/config"
	"github.com/rms-diego/book-my-session/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	if err := logger.Init(); err != nil {
		panic(err)
	}

	if err := config.Init(); err != nil {
		panic(err)
	}

	r.Use(middleware.LogsMiddleware(logger.Log))
	r.Use(middleware.ErrorMiddleware(logger.Log))

	routes.Init(r)

	s := &http.Server{
		Addr:    ":" + config.Env.PORT,
		Handler: r,
	}

	logger.Log.Info("Server is running")
	logger.Log.Info(fmt.Sprintf("Address: http://localhost:%v", config.Env.PORT))
	if err := s.ListenAndServe(); err != nil {
		logger.Log.Fatal("Server failed to start", zap.Error(err))
	}
}
