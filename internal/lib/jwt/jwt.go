package jwt

import (
	"se_project/internal/consts"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}

func GenAccessToken(UserId string) (tokenString string, err error) {
	atc := &AccessTokenClaims{
		UserId: UserId,
		Role:   consts.TypeUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	return token.SignedString([]byte(consts.JwtKey))
}

func ParseAccessToken(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*AccessTokenClaims); !ok {
		return nil, gerror.New("In ParseAccessToken 类型断言失败")
	}
	return token, nil
}
