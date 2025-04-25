package model

import (
	"context"
	"database/sql"
	"time"
)

var (
	UserDataSaved = "user data was saved"
	NoSuchUser    = "no such user"

	StartedAutoClicker = "auto clicker was started"
	StoppedAutoClicker = "auto clicker was stopped"
	HasNotAutoClicker  = "user has not auto clicker"

	CryptoDidNotSend = "crypto did not send"
)

type User struct {
	ID int64 `db:"id"`

	TgID      sql.NullInt64  `db:"tg_id"`
	Username  sql.NullString `db:"username"`
	Wallet    sql.NullString `db:"wallet"`
	Balance   sql.NullInt64  `db:"balance"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	CreatedAt time.Time      `db:"created_at"`
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
	UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error
	SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, energyBooster int64, balance int64, tgID int64) error
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user User) (int64, error)
	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
	UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error
	SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, energyBooster int64, balance int64, tgID int64) error
}
