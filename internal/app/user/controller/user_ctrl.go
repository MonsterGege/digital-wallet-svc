package controller

import (
	"context"
	"log"

	"github.com/digital-wallet-svc/internal/app/user/implementations"
	"github.com/digital-wallet-svc/internal/app/user/models"
	"github.com/digital-wallet-svc/internal/app/user/services"
	"github.com/digital-wallet-svc/pkg/database"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(ctx context.Context, db *database.Database) *UserController {
	impl := implementations.NewUserImplementation(ctx, db)
	svc := services.NewUserService(ctx, impl)
	return &UserController{UserService: svc}
}

func (uc *UserController) RegisterRoutes(g *gin.RouterGroup) {
	// Implement user handling route here
	g.GET("/withdraw_funds", uc.WithdrawFunds)
	g.GET("/balance_wallet", uc.BalanceWallet)
}

func (uc *UserController) WithdrawFunds(g *gin.Context) {
	// Implement withdraw funds logic here
	var param models.UserParam
	if err := g.ShouldBind(&param); err != nil {
		g.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}
	data, err := uc.UserService.WithdrawFunds(g.Request.Context(), param)
	if err != nil {
		log.Println("Error Withdraw Funds :", err)
		g.JSON(500, gin.H{"error": err.Error()})
	} else {
		g.JSON(200, gin.H{"message": data})
	}
}

func (uc *UserController) BalanceWallet(g *gin.Context) {
	// Implement balance wallet logic here
	var param models.UserParam
	if err := g.ShouldBind(&param); err != nil {
		g.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}
	data, err := uc.UserService.BalanceWallet(g.Request.Context(), param)
	if err != nil {
		log.Println("Error Balance Wallet :", err)
		g.JSON(500, gin.H{"error": err.Error()})
	} else {
		g.JSON(200, data)
	}
}
