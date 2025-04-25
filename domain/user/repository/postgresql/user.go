package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/pskoob/miniappBack/database/postgresql"
	"github.com/pskoob/miniappBack/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (int64, error) {
	query, params, err := postgresql.Builder.Insert("users").
		Columns(
			"tg_id",
			"username",
		).
		Values(
			user.TgID,
			user.Username,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	var userID int64
	err = r.sqalxConn.GetContext(ctx, &userID, query, params...)
	return userID, err
}

func (r *UserRepository) GetUserByTgID(ctx context.Context, tgID int64) (model.User, error) {
	var user model.User
	query, params, err := postgresql.Builder.Select(
		"users.id",
		"users.tg_id",
		"users.username",
		"users.wallet",
		"users.balance",
		"users.updated_at",
		"users.created_at",
	).
		From("users").
		Where(sq.Eq{"users.tg_id": tgID}).ToSql()
	if err != nil {
		return user, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &user, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
	}

	return user, err
}

func (r *UserRepository) UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error {
	query, params, err := postgresql.Builder.Update("users").
		Set("balance", balance).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"tg_id": tgID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *UserRepository) SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, energyBooster int64, balance int64, tgID int64) error {
	query, params, err := postgresql.Builder.Update("users").
		Set("balance", balance).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"tg_id": tgID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
