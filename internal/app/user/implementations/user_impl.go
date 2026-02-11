package implementations

import (
	"context"

	"github.com/digital-wallet-svc/internal/app/user/models"
	"github.com/digital-wallet-svc/pkg/database"
	"github.com/jackc/pgx/v5"
)

type UserImplemantation struct {
	ctx context.Context
	db  *database.Database
	tx  *pgx.Tx
}

func NewUserImplementation(ctx context.Context, db *database.Database) *UserImplemantation {
	return &UserImplemantation{ctx: ctx, db: db}
}

func (ui *UserImplemantation) CheckUserExists(ctx context.Context, userID int) (bool, error) {
	pool := ui.db.GetConnection(ctx)
	var exists bool
	q := `
		SELECT EXISTS (
		SELECT 1 FROM users
		WHERE id = $1
		)	
	`
	rows, err := pool.Query(ctx, q, userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			return false, err
		}
	}
	return exists, nil
}

func (ui *UserImplemantation) WithdrawFunds(ctx context.Context, walletId int, userID int, amount float64) error {
	// Implement the logic to withdraw funds from the user's wallet
	pool := ui.db.GetConnection(ctx)
	q := `
		UPDATE wallets
		SET balance = balance - $1
		WHERE wallet_id = $2 and user_id = $3
	`
	_, err := pool.Exec(ctx, q, amount, walletId, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ui *UserImplemantation) UpdateTransactionHistory(ctx context.Context, walletID int, amount float64) error {
	// Implement the logic to update the transaction history of the user
	q := `
		INSERT INTO transactions (wallet_id,amount, types)
		VALUES ($1, $2, 'withdraw')
		`
	err := ui.db.ExecuteTransaction(ctx, *ui.tx, q, walletID, amount)
	if err != nil {
		return err
	}
	return nil

}

func (ui *UserImplemantation) BalanceWallet(ctx context.Context, walletId int, userID int) (models.ResponseBalance, error) {
	// Implement the logic to get the balance of the user's wallet
	pool := ui.db.GetConnection(ctx)
	var balance float64
	q := `
		SELECT balance FROM wallets
		WHERE wallet_id = $1 and user_id = $2
	`
	rows, err := pool.Query(ctx, q, walletId, userID)
	if err != nil {
		return models.ResponseBalance{}, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&balance)
		if err != nil {
			return models.ResponseBalance{}, err
		}
	}
	response := models.ResponseBalance{
		Balance: balance,
	}
	return response, nil
}

func (ui *UserImplemantation) UsingTransactions(ctx context.Context) (*UserImplemantation, error) {
	tx, err := ui.db.UsingTransactions(ctx)
	if err != nil {
		return nil, err
	}
	return &UserImplemantation{ctx: ctx, db: ui.db, tx: &tx}, nil
}

func (ui *UserImplemantation) Rollback(ctx context.Context) error {
	return ui.db.Rollback(ctx, *ui.tx)
}

func (ui *UserImplemantation) Commit(ctx context.Context) error {
	return ui.db.Commit(ctx, *ui.tx)
}
