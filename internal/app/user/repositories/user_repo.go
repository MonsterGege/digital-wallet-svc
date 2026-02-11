package repositories

import (
	"context"

	"github.com/digital-wallet-svc/internal/app/user/implementations"
	"github.com/digital-wallet-svc/internal/app/user/models"
)

type UserRepository interface {
	// Define necessary methods here
	CheckUserExists(ctx context.Context, userID int) (bool, error)
	WithdrawFunds(ctx context.Context, walletID int, userID int, amount float64) error
	UpdateTransactionHistory(ctx context.Context, walletId int, amount float64) error
	BalanceWallet(ctx context.Context, walletId int, userID int) (models.ResponseBalance, error)
	UsingTransactions(ctx context.Context) (*implementations.UserImplemantation, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}
