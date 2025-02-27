package model

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID int64 `db:"id"`

	Name         sql.NullString `db:"name"`
	TgID         sql.NullInt64  `db:"tg_id"`
	Username     sql.NullString `db:"username"`
	Balance      sql.NullInt64  `db:"balance"`
	ClickBooster sql.NullInt64  `db:"click_booster"`
	AutoClicker  bool           `db:"auto_clicker"`
	UpdatedAt    sql.NullTime   `db:"updated_at"`
	CreatedAt    time.Time      `db:"created_at"`
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
}
