package user

import (
	"github.com/avalonprod/gasstrem/src/pkg/hasher"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDeps struct {
	hasher   *hasher.Hasher
	Database *mongo.Database
}

func Init(deps UserDeps, router *gin.RouterGroup) {
	sotorage := NewUserStorage(deps.Database)
	service := NewUserService(deps, *sotorage)

	NewUserController(router, service)

}
