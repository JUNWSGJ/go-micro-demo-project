package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	Id            uint64
	Mobile        string
	Nickname      string
	WxOpenid      string
	LastLoginTime time.Time
	CreateTime    time.Time
}

type UserAuth struct {
	Mobile   string
	AuthType string
	AuthCode string
}

type UserRepo interface {
	SaveUser(ctx context.Context, u *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper("usecase/user", logger)}
}

func (uc *UserUsecase) RegisterOrLogin(ctx context.Context, u *User) (*User, error) {
	return uc.repo.SaveUser(ctx, u)
}
