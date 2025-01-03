package delivery

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	CreateUser(ctx *gin.Context)
}

type userDelivery struct {
	userUsecase usecase.UserUsecase
}

func NewUserDelivery(userUsecase usecase.UserUsecase) UserDelivery {
	return &userDelivery{userUsecase: userUsecase}
}

func (h *userDelivery) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.userUsecase.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": result})

}