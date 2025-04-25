package model

import (
	"context"
	"database/sql"
	"time"
)

const (
	EnergyBooster = "energy_booster"
	AutoClicker   = "auto_clicker"
	PowerClick    = "power_click"
)

type Card struct {
	ID int64 `db:"id"`

	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	MaxLevel    sql.NullInt64  `db:"max_level"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	CreatedAt   time.Time      `db:"created_at"`
}

type ICardRepository interface {
	GetCardByID(ctx context.Context, cardID int64) (Card, error)
	GetCardByName(ctx context.Context, cardName string) (Card, error)
}

type ICardUsecase interface {
	GetCardByID(ctx context.Context, cardID int64) (Card, error)
	GetCardByName(ctx context.Context, cardName string) (Card, error)
}
