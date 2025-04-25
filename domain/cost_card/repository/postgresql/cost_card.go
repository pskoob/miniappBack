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

type CostCardRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ICostCardRepository {
	return &CostCardRepository{sqalxConn: sqalxConn}
}

func (r *CostCardRepository) GetCostCardByCardID(ctx context.Context, cardID int64) (model.CostCard, error) {
	var costCard model.CostCard
	query, params, err := postgresql.Builder.Select(
		"cost_cards.card_id",
		"cost_cards.card_name",
		"cost_cards.base_price",
		"cost_cards.growth_coefficient",
		"cost_cards.updated_at",
		"cost_cards.created_at",
	).
		From("cost_cards").
		Where(sq.Eq{"cost_cards.card_id": cardID}).ToSql()
	if err != nil {
		return costCard, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &costCard, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return costCard, err
		}
	}

	return costCard, err
}
