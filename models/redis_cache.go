package models

import (
	"context"
	"fmt"
	"gochat/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// SetUserOnlineInfo 将在线用户存到缓存, 使用string
func SetUserOnlineInfo(userId int, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	key := "online_" + strconv.Itoa(userId)
	err := utils.Rdb.Set(ctx, key, val, timeTTL).Err()
	if err != nil {
		fmt.Println("rdb.set err:", err)
		return
	}
}

// DeleteUserOnline 用户下线后，将用户从缓存中删除
func DeleteUserOnline(userId uint) {
	ctx := context.Background()
	key := "online_" + strconv.Itoa(int(userId))
	utils.Rdb.Del(ctx, key)
}

// GetOnlineUser 判断目标用户是否在线
func GetOnlineUser(targetId uint) bool {
	ctx := context.Background()
	result, err := utils.Rdb.Get(ctx, "online_"+strconv.Itoa(int(targetId))).Result()
	if err != nil {
		fmt.Println("rdb.get err:", err)
		return false
	}
	if result == "" {
		return false
	}
	return true
}

// SetMessage 缓存离线消息
func SetMessage(userId uint, targetId uint, msg []byte) {
	ctx := context.Background()
	//双方互相发送的消息共用一个key
	var key string
	if userId < targetId {
		key = fmt.Sprintf("msg_%d_%d", userId, targetId)
	} else {
		key = fmt.Sprintf("msg_%d_%d", targetId, userId)
	}
	result, _ := utils.Rdb.ZRange(ctx, key, 0, -1).Result()
	len := float64(len(result) + 1)
	utils.Rdb.ZAdd(ctx, key, &redis.Z{len, msg})
}

func GetRedisMsg(userId uint, targetId uint) (r []string, err error) {
	ctx := context.Background()
	var key string
	if userId < targetId {
		key = fmt.Sprintf("msg_%d_%d", userId, targetId)
	} else {
		key = fmt.Sprintf("msg_%d_%d", targetId, userId)
	}
	r, err = utils.Rdb.ZRange(ctx, key, 0, -1).Result()
	return
}
