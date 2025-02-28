package model

import (
	"context"
	"database/sql"
	"time"
)

var (
	UserDataSaved = "user data was saved"
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
	SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, balance int64, tgID int64) error
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
	SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, balance int64, tgID int64) error
}
