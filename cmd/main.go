package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/masred/my-gram/app/controller/http/api"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/repository"
	"github.com/masred/my-gram/app/service"
	"github.com/masred/my-gram/config"
)

func main() {
	config.New()

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "ez"
	}

	validate := validator.New()
	jwtService := helper.NewUserJWTService([]byte(secretKey))
	db, err := config.NewPostgresDatabase()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate, jwtService)
	userController := api.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository, validate)
	photoController := api.NewPhotoController(photoService)

	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	userController.Route(r)
	photoController.Route(r)

	fmt.Println("Server started at " + host + ":" + port)
	http.ListenAndServe(host+":"+port, r)
}
