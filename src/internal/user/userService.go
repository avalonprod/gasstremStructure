package user

import "context"

type userService struct {
	deps    UserDeps
	storage userStorage
}

type iUserService interface {
	Create(ctx context.Context)
}

func NewUserService(deps UserDeps, userStorage userStorage) *userService {
	return &userService{
		deps:    deps,
		storage: userStorage,
	}
}

func (u *userStorage) CreateUserService(ctx context.Context) {

}
