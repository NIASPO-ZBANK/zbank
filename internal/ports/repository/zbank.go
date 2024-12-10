package repository

import (
	"context"

	"users/internal/models/entities"
)

type MoneyRepository interface {
	AddMoney(ctx context.Context, amount int) error
	WithdrawMoney(ctx context.Context, amount int) error
	DepositMoney(ctx context.Context, amount int) error
	WithdrawFromDeposit(ctx context.Context, amount int) error
	GetBalanceAndDeposit(ctx context.Context) (*entities.BalanceInfo, error)
}
