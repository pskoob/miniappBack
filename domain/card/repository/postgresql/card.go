package postgresql

import (
	"context"

	"database/sql"
	"errors"

	"github.com/pskoob/miniappBack/database/postgresql"
	"github.com/pskoob/miniappBack/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type CardRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ICardRepository {
	return &CardRepository{sqalxConn: sqalxConn}
}

var CardFields = []string{
	"cards.id",
	"cards.name",
	"cards.description",
	"cards.max_level",
	"cards.updated_at",
	"cards.created_at",
}

func (r *CardRepository) GetCardByID(ctx context.Context, cardID int64) (model.Card, error) {
	var card model.Card

	query, params, err := postgresql.Builder.Select(CardFields...).
		From("cards").
		Where(sq.Eq{"cards.id": cardID}).ToSql()

	if err != nil {
		return card, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	if err = r.sqalxConn.GetContext(ctx, &card, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return card, err
		}
		return card, err
	}

	return card, nil
}

func (r *CardRepository) GetCardByName(ctx context.Context, cardName string) (model.Card, error) {
	var card model.Card

	query, params, err := postgresql.Builder.Select(CardFields...).
		From("cards").
		Where(sq.Eq{"cards.name": cardName}).ToSql()

	if err != nil {
		return card, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	if err = r.sqalxConn.GetContext(ctx, &card, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return card, err
		}
		return card, err
	}

	return card, nil
}
