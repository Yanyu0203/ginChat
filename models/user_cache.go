package models

import (
	"context"
	"ginchat/utils"
	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTLL time.Duration) {
	ctx := context.Background()
	utils.Red.Set(ctx, key, val, timeTLL)
}
