package zbank

import (
	"context"
	"database/sql"
	"fmt"
	"users/internal/models/entities"

	"github.com/Masterminds/squirrel"

	"users/internal/errs"
	"users/internal/infrastructure/database/postgres"
	"users/internal/ports/repository"
)

const (
	moneyTable = "money"
	userID     = "GOLBUTSA-1337-1487-911Z-Salla4VO2022"
)

type moneyRepository struct {
	db *postgres.Postgres
}

func NewMoneyRepository(
	db *postgres.Postgres,
) repository.MoneyRepository {
	return &moneyRepository{
		db: db,
	}
}

func (r *moneyRepository) AddMoney(ctx context.Context, amount int) error {
	qb := r.db.Builder.
		Update(moneyTable).
		Set("money", squirrel.Expr("money + ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := r.db.SqlxDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *moneyRepository) WithdrawMoney(ctx context.Context, amount int) error {
	qb := r.db.Builder.
		Update(moneyTable).
		Set("money", squirrel.Expr("money - ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	query, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("to sql: %w", err)
	}

	res, err := r.db.SqlxDB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *moneyRepository) DepositMoney(ctx context.Context, amount int) error {
	tx, err := r.db.SqlxDB().BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// Withdraw from money
	withdrawQB := r.db.Builder.
		Update(moneyTable).
		Set("money", squirrel.Expr("money - ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	withdrawQuery, withdrawArgs, err := withdrawQB.ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("to sql: %w", err)
	}

	if _, err := tx.ExecContext(ctx, withdrawQuery, withdrawArgs...); err != nil {
		tx.Rollback()
		return fmt.Errorf("exec: %w", err)
	}

	// Add to deposit
	depositQB := r.db.Builder.
		Update(moneyTable).
		Set("deposit", squirrel.Expr("deposit + ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	depositQuery, depositArgs, err := depositQB.ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("to sql: %w", err)
	}

	if _, err := tx.ExecContext(ctx, depositQuery, depositArgs...); err != nil {
		tx.Rollback()
		return fmt.Errorf("exec: %w", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

func (r *moneyRepository) WithdrawFromDeposit(ctx context.Context, amount int) error {
	tx, err := r.db.SqlxDB().BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	withdrawQB := r.db.Builder.
		Update(moneyTable).
		Set("deposit", squirrel.Expr("deposit - ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	withdrawQuery, withdrawArgs, err := withdrawQB.ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("to sql: %w", err)
	}

	if _, err := tx.ExecContext(ctx, withdrawQuery, withdrawArgs...); err != nil {
		tx.Rollback()
		return fmt.Errorf("exec: %w", err)
	}

	// Add to money
	addQB := r.db.Builder.
		Update(moneyTable).
		Set("money", squirrel.Expr("money + ?", amount)).
		Where(squirrel.Eq{"user_id": userID})

	addQuery, addArgs, err := addQB.ToSql()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("to sql: %w", err)
	}

	if _, err := tx.ExecContext(ctx, addQuery, addArgs...); err != nil {
		tx.Rollback()
		return fmt.Errorf("exec: %w", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

func (r *moneyRepository) GetBalanceAndDeposit(ctx context.Context) (*entities.BalanceInfo, error) {
	qb := r.db.Builder.
		Select("money", "deposit").
		From(moneyTable).
		Where(squirrel.Eq{"user_id": userID})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	var balanceInfo entities.BalanceInfo
	if err = r.db.SqlxDB().GetContext(ctx, &balanceInfo, query, args...); err != nil {
		return nil, fmt.Errorf("get context: %w", err)
	}

	return &balanceInfo, nil
}
