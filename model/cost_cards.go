package model

import (
	"context"
	"database/sql"
	"time"
)

type CostCard struct {
	CardID            int64        `db:"card_id"`
	CardName          string       `db:"card_name"`
	BasePrice         int64        `db:"base_price"`
	GrowthCoefficient float64      `db:"growth_coefficient"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
	CreatedAt         time.Time    `db:"created_at"`
}

type ICostCardRepository interface {
	GetCostCardByCardID(ctx context.Context, cardID int64) (CostCard, error)
}

type ICostCardUsecase interface {
	GetCostCardByCardID(ctx context.Context, cardID int64) (CostCard, error)
}
