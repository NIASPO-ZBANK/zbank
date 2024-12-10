package usecase

import (
	"context"
	"users/internal/models/dto"
)

type MoneyUseCase interface {
	AddFunds(ctx context.Context, amount int) error
	WithdrawFunds(ctx context.Context, amount int) error
	DepositFunds(ctx context.Context, amount int) error
	WithdrawFromDeposit(ctx context.Context, amount int) error
	GetBalanceAndDeposit(ctx context.Context) (*dto.BalanceInfo, error)
}
