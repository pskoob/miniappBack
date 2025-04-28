package model

import (
	"context"
	"database/sql"
	"time"
)

var (
	UserDataSaved             = "user data was saved"
	NoSuchUser                = "no such user"
	BalanceTooLow             = "user balance too low"
	UserAlreadyHasAutoClicker = "user already has auto clicker"
	WalletWasUpdated          = "user wallet was updated"
	UserHaveReferralUser      = "user have referral user"

	CardHasMaxLvl = "card has max lvl"

	StartedEnergyCollect = "energy collect was started"
	StoppedEnergyCollect = "energy collect was stopped"

	StartedAutoClicker = "auto clicker was started"
	StoppedAutoClicker = "auto clicker was stopped"
	HasNotAutoClicker  = "user has not auto clicker"

	CryptoDidNotSend = "crypto did not send"
)

type User struct {
	ID int64 `db:"id"`

	TgID                    sql.NullInt64  `db:"tg_id"`
	Username                sql.NullString `db:"username"`
	Wallet                  sql.NullString `db:"wallet"`
	Balance                 int64          `db:"balance"`
	Energy                  int64          `db:"energy"`
	ReferralUserID          sql.NullInt64  `db:"referral_user_id"`
	ReferralLink            string         `db:"referral_link"`
	TotalCoinEarned         int64          `db:"total_coin_earned"`
	TotalCoinSentToReferral int64          `db:"total_coin_referral_spent"`
	UpdatedAt               sql.NullTime   `db:"updated_at"`
	CreatedAt               time.Time      `db:"created_at"`
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user User) (int64, error)

	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
	GetUserByReferralLink(ctx context.Context, referralLink string) (User, error)
	GetUsersByReferralUser(ctx context.Context, userID int64) ([]User, error)

	UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error
	UpdateEnergyByTgID(ctx context.Context, energy int64, tgID int64) error
	UpdateWalletByTgID(ctx context.Context, wallet string, tgID int64) error
	UpdateUserReferralByTgID(ctx context.Context, referralUserID int64, tgID int64) error
	UpdateTotalCoinEarnedByTgID(ctx context.Context, totalCoinEarned int64, tgID int64) error
	UpdateTotalCoinSentToReferralByTgID(ctx context.Context, totalCoinSentToReferral int64, tgID int64) error
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user User) (int64, error)

	GetUserByTgID(ctx context.Context, tgID int64) (User, error)
	GetUserByReferralLink(ctx context.Context, referralLink string) (User, error)
	GetUsersByReferralUser(ctx context.Context, userID int64) ([]User, error)

	UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error
	UpdateEnergyByTgID(ctx context.Context, energy int64, tgID int64) error
	UpdateWalletByTgID(ctx context.Context, wallet string, tgID int64) error
	UpdateUserReferralByTgID(ctx context.Context, referralUserID int64, tgID int64) error
	UpdateTotalCoinEarnedByTgID(ctx context.Context, totalCoinEarned int64, tgID int64) error
	UpdateTotalCoinSentToReferralByTgID(ctx context.Context, totalCoinSentToReferral int64, tgID int64) error
}
