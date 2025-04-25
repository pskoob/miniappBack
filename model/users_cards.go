package model

import (
	"context"
	"database/sql"
	"time"
)

type UserCard struct {
	UserID       int64        `db:"user_id"`
	CardID       int64        `db:"card_id"`
	CardName     string       `db:"card_name"`
	CurrentLevel int64        `db:"current_level"`
	AutoClicker  sql.NullBool `db:"auto_clicker"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	CreatedAt    time.Time    `db:"created_at"`
}

type IUserCardRepository interface {
	CreateUserCardByUserID(ctx context.Context, userCard UserCard) (int64, error)
	GetUserCardsByUserID(ctx context.Context, userID int64) ([]UserCard, error)
	UpdateUserCardByUserID(ctx context.Context, userCard UserCard) error
}

type IUserCardUsecase interface {
	CreateUserCardByUserID(ctx context.Context, userCard UserCard) (int64, error)
	GetUserCardsByUserID(ctx context.Context, userID int64) ([]UserCard, error)
	UpdateUserCardByUserID(ctx context.Context, userCard UserCard) error
}
