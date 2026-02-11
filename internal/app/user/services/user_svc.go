package services

import (
	"context"
	"errors"

	"github.com/digital-wallet-svc/internal/app/user/models"
	"github.com/digital-wallet-svc/internal/app/user/repositories"
)

type UserService struct {
	// Add necessary fields here
	ctx  context.Context
	repo repositories.UserRepository
}

func NewUserService(ctx context.Context, repo repositories.UserRepository) *UserService {
	return &UserService{
		ctx:  ctx,
		repo: repo,
	}
}

func (us *UserService) WithdrawFunds(ctx context.Context, param models.UserParam) (string, error) {
	// Implement the logic to withdraw funds using the repository

	user, err := us.repo.CheckUserExists(ctx, param.ID)
	if err != nil {
		return "", errors.New("Failed to check user existence")
	}
	if user == false {
		return "", errors.New("User Not Found")
	}

	balance, err := us.repo.BalanceWallet(ctx, param.WalletID, param.ID)
	if err != nil {
		return "", errors.New("Failed to get wallet balance")
	}
	if balance.Balance < param.WithdrawalAmount {
		return "", errors.New("Insufficient Funds")
	}

	tx, err := us.repo.UsingTransactions(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	err = tx.WithdrawFunds(ctx, param.WalletID, param.ID, param.WithdrawalAmount)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	err = tx.UpdateTransactionHistory(ctx, param.WalletID, param.WithdrawalAmount)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return `Funds withdrawn successfully `, nil
}
func (us *UserService) BalanceWallet(ctx context.Context, param models.UserParam) (models.ResponseBalance, error) {
	// Implement the logic to get wallet balance using the repository
	balance, err := us.repo.BalanceWallet(ctx, param.WalletID, param.ID)
	if err != nil {
		return models.ResponseBalance{}, err
	}
	return balance, nil
}
