package http

import (
	"context"
	"os"

	"github.com/digital-wallet-svc/internal/app/user"
	"github.com/digital-wallet-svc/pkg/database"
	"github.com/gin-gonic/gin"
)

type Route struct {
	ctx  context.Context
	user *user.UserModule
}

func NewRoute(ctx context.Context, db *database.Database) (*Route, error) {
	// Initialize user module
	user, err := user.NewUserModule(ctx, db)
	if err != nil {
		return nil, err
	}
	return &Route{ctx: ctx, user: user}, nil
}

func (r *Route) Run(g *gin.Engine, port string) error {
	// Register user routes
	r.user.RegisterRoutes(g)
	return g.Run(port)
}

func StartServer(ctx context.Context, port string) error {
	router := gin.Default()
	// Initialize database connection
	db, err := database.NewDatabase(ctx, os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"))
	if err != nil {
		return err
	}
	defer db.CloseConnection(ctx)
	// Initialize and run the server
	server, err := NewRoute(ctx, db)
	if err != nil {
		return err
	}
	return server.Run(router, port)
}
