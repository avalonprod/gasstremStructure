package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type userStorage struct {
	db *mongo.Collection
}

type iUserStorage interface {
	Create(ctx context.Context)
}

func NewUserStorage(db *mongo.Database) *userStorage {
	return &userStorage{
		db: db.Collection("users"),
	}
}

func (u *userStorage) Create(ctx context.Context) {

}
