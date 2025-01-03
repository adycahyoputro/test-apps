package router

import (
	"backend/delivery"
	"backend/repository"
	"backend/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)

	r.POST("/register", userDelivery.CreateUser)
	return r
}