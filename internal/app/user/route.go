package user

import (
	"context"

	ctrl_user "github.com/digital-wallet-svc/internal/app/user/controller"
	"github.com/digital-wallet-svc/pkg/database"
	"github.com/gin-gonic/gin"
)

type UserModule struct {
	ctx       context.Context
	user_ctrl *ctrl_user.UserController
}

func NewUserModule(ctx context.Context, db *database.Database) (*UserModule, error) {
	user_ctrl := ctrl_user.NewUserController(ctx, db)
	return &UserModule{ctx: ctx, user_ctrl: user_ctrl}, nil
}

func (u *UserModule) RegisterRoutes(g *gin.Engine) {
	// Define user-related routes here
	user := g.Group("/user")
	{
		u.user_ctrl.RegisterRoutes(user)
	}
}
