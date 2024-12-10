package entities

type BalanceInfo struct {
	Balance int `db:"money"`
	Deposit int `db:"deposit"`
}
