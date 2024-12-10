package converters

import (
	"users/internal/models/dto"
	"users/internal/models/entities"
)

type MoneyConverter struct{}

func NewMoneyConverter() *MoneyConverter {
	return &MoneyConverter{}
}

func (c *MoneyConverter) ToBalanceInfoDTO(balance *entities.BalanceInfo) *dto.BalanceInfo {
	return &dto.BalanceInfo{
		Balance: balance.Balance,
		Deposit: balance.Deposit,
	}
}
