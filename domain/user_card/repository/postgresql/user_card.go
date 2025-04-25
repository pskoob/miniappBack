package postgresql

import (
	"context"

	"github.com/pskoob/miniappBack/database/postgresql"
	"github.com/pskoob/miniappBack/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type UserCardRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserCardRepository {
	return &UserCardRepository{sqalxConn: sqalxConn}
}

var UserCardFields = []string{
	"user_id",
	"card_id",
	"card_name",
	"current_level",
	"auto_clicker",
	"updated_at",
	"created_at",
}

func (r *UserCardRepository) CreateUserCardByUserID(ctx context.Context, userCard model.UserCard) (int64, error) {
	query, params, err := postgresql.Builder.Insert("users_cards").
		Columns(
			"user_id",
			"card_name",
			"card_id",
		).
		Values(
			userCard.UserID,
			userCard.CardName,
			userCard.CardID,
		).
		Suffix("RETURNING card_id").ToSql()

	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var card_id int64
	err = r.sqalxConn.GetContext(ctx, &card_id, query, params...)
	if err != nil {
		return 0, err
	}

	return card_id, nil
}

func (r *UserCardRepository) GetUserCardsByUserID(ctx context.Context, userID int64) ([]model.UserCard, error) {
	var userCards []model.UserCard

	query, params, err := postgresql.Builder.Select(UserCardFields...).
		From("users_cards").
		Where(sq.Eq{"user_id": userID}).
		ToSql()

	if err != nil {
		return userCards, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	err = r.sqalxConn.SelectContext(ctx, &userCards, query, params...)
	if err != nil {
		return userCards, err
	}

	return userCards, nil
}

func (r *UserCardRepository) UpdateUserCardByUserID(ctx context.Context, userCard model.UserCard) error {
	query, params, err := postgresql.Builder.Update("users_cards").
		Set("auto_clicker", userCard.AutoClicker.Bool).
		Set("current_level", userCard.CurrentLevel).
		Set("updated_at", userCard.UpdatedAt).
		Where(sq.Eq{"user_id": userCard.UserID, "card_id": userCard.CardID}).
		ToSql()

	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}
