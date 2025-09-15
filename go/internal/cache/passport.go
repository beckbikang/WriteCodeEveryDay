package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/beckbikang/WriteCodeEveryDay/internal/model"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/app"
)

const (
	PassportCacheKeyPrefix = "passport:user_name:"
	PassportCacheExpire    = 30 * time.Minute
)

func SetPassportCache(userName string, passport *model.Passport) error {
	key := fmt.Sprintf("%s%s", PassportCacheKeyPrefix, userName)
	
	data, err := json.Marshal(passport)
	if err != nil {
		return err
	}
	
	return RedisClient.Set(context.Background(), key, data, PassportCacheExpire).Err()
}

func GetPassportCache(userName string) (*model.Passport, error) {
	key := fmt.Sprintf("%s%s", PassportCacheKeyPrefix, userName)
	
	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	
	var passport model.Passport
	err = json.Unmarshal([]byte(val), &passport)
	if err != nil {
		return nil, err
	}
	
	return &passport, nil
}

func DeletePassportCache(userName string) error {
	key := fmt.Sprintf("%s%s", PassportCacheKeyPrefix, userName)
	return RedisClient.Del(context.Background(), key).Err()
}
