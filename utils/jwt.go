package utils

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

type MyClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

// TokenExpireDuration
// TODO set a longer time for testing
const TokenExpireDuration = time.Minute

var MySecret = []byte("secret key")

// GenToken gen JWT token and set the configration
func GenToken(ID int64) (string, error) {
	// create our own claims map
	c := MyClaims{
		ID, // store ID in Preload
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TokenExpireDuration)},
			Issuer:    "douyin-demo",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenStr, existGet := c.GetQuery("token")
		if !existGet {
			tokenStr = c.PostForm("token")
		}

		//jwt framework parsetoken version
		//mc, err := ParseToken(tokenStr)

		UserStr, err := RedisParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token/Token验证错误/Token过期",
			})
			c.Abort()
		}
		c.Set("UserStr", UserStr)
		c.Next()

		// 将当前请求的username信息保存到请求的上下文c上
		//c.Set("ID", mc.ID)
		//c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// RedisParseToken redis parse token version
// key:login:session:"+tokenStr, value:user TTL:10min
func RedisParseToken(tokenStr string) (string, error) {
	res := global.App.DY_REDIS.Get(context.Background(), global.REDIS_USER_PREFIX+tokenStr)
	log.Println("jwt check from redis:", res)
	UserStr := res.Val()
	// if token already expired, abort
	if UserStr == "" {
		return "", errors.New("error: parse token ")
	}
	// if token in use, then refresh the TTL˜
	global.App.DY_REDIS.Expire(context.Background(), global.REDIS_USER_PREFIX+tokenStr, global.REDIS_USER_TTL)
	return UserStr, nil
}
