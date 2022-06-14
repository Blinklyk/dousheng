package global

import (
	"time"
)

const (
	SECRETKEY              = "secrete key"
	REDIS_USER_PREFIX      = "login:session:"
	REDIS_USER_TTL         = time.Minute * 60
	LOCAL_FILE_PATH_PREFIX = "public/"
)
