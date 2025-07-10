package main

import (
	"fmt"
	"io"
	"live/internal/config"
	"live/internal/repository"
	"live/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"live/internal/database"
	"live/internal/handler"
	"live/models"

	"github.com/sirupsen/logrus"
)

func main() {
	InitLogger()
	cfg, err := config.InitializationConfig()
	if err != nil {
		logrus.Fatalf("failed start server, error:%v", err)
	}

	db, err := database.ConnectDD(cfg.DBConfig)
	if err != nil {
		logrus.WithError(err).Error("failed connect to db")
	}

	db.AutoMigrate(&models.Comment{})

	var comments []models.Comment

	db.Find(&comments)

	fmt.Print(comments)

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	handler := handler.NewHandler(service)
	router := handler.InitRout()

	srv := http.Server{
		Handler:      router,
		Addr:         cfg.HTTPServer.Address,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
	}

	done := make(chan os.Signal, 1)

	go func() {
		srv.ListenAndServe()
	}()

	logrus.Info("Server started")
	signal.Notify(done, syscall.SIGINT, syscall.SIGINT)
	<-done
	logrus.Info("Server shutdown")
}

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, file)

	logrus.SetOutput(mw)

	logrus.SetLevel(logrus.DebugLevel)

}
