package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GMD_DB *gorm.DB
	GMD_RD *redis.Client
)
