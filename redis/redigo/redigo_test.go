package redigo

import "github.com/Sagleft/redsync/redis"

var _ redis.Conn = (*conn)(nil)

var _ redis.Pool = (*pool)(nil)
