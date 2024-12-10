package zbank

import (
	"context"
	"fmt"
	"users/internal/converters"
	"users/internal/models/dto"
	"users/internal/ports/repository"
	"users/internal/usecase"
)

type moneyUseCase struct {
	moneyRepository repository.MoneyRepository
	moneyConverter  *converters.MoneyConverter
}

func NewMoneyUseCase(
	moneyRepository repository.MoneyRepository,
) usecase.MoneyUseCase {
	return &moneyUseCase{
		moneyRepository: moneyRepository,
		moneyConverter:  converters.NewMoneyConverter(),
	}
}

// AddFunds adds a specified amount to the user's account
func (muc *moneyUseCase) AddFunds(ctx context.Context, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	err := muc.moneyRepository.AddMoney(ctx, amount)
	if err != nil {
		return fmt.Errorf("repo add money: %w", err)
	}

	return nil
}

// WithdrawFunds withdraws a specified amount from the user's account
func (muc *moneyUseCase) WithdrawFunds(ctx context.Context, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	err := muc.moneyRepository.WithdrawMoney(ctx, amount)
	if err != nil {
		return fmt.Errorf("repo withdraw money: %w", err)
	}

	return nil
}

// DepositFunds moves a specified amount from the user's account to their deposit
func (muc *moneyUseCase) DepositFunds(ctx context.Context, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	err := muc.moneyRepository.DepositMoney(ctx, amount)
	if err != nil {
		return fmt.Errorf("repo deposit money: %w", err)
	}

	return nil
}

// WithdrawFromDeposit withdraws a specified amount from the user's deposit back to their account
func (muc *moneyUseCase) WithdrawFromDeposit(ctx context.Context, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	err := muc.moneyRepository.WithdrawFromDeposit(ctx, amount)
	if err != nil {
		return fmt.Errorf("repo withdraw from deposit: %w", err)
	}

	return nil
}

func (muc *moneyUseCase) GetBalanceAndDeposit(ctx context.Context) (*dto.BalanceInfo, error) {
	balance, err := muc.moneyRepository.GetBalanceAndDeposit(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get balance and deposit: %w", err)
	}

	return muc.moneyConverter.ToBalanceInfoDTO(balance), nil
}
