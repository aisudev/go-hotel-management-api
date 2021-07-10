package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"poke/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Mock context user UUID

			token := c.Request().Header.Get("Authorization")
			tokenSplit := strings.Split(token, " ")

			uuid, err := VerifyToken("access", tokenSplit[1], []byte(viper.GetString("jwt.access_secret")))

			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			c.Set("uuid", uuid)

			return next(c)
		}
	}
}

// signin -> hash uuid -> set it to secret and save it in redis
type Claim struct {
	UUID   string `json:"uuid"`
	Secret string `json:"secret"`
	jwt.StandardClaims
}

var ctx = context.Background()
var rd = utils.RedisInit()

func GenerateToken(code string, uid string, expire int64, jwtKey []byte) (string, error) {
	if code != "access" && code != "refresh" {
		return "", errors.New("code invalid.")
	}

	secret, _ := uuid.DefaultGenerator.NewV4()
	hashSecret := utils.GenerateHash(secret.String())

	rd_key := fmt.Sprintf("%s_%s", code, uid)
	rd.Set(ctx, rd_key, secret.String(), 0)

	claim := &Claim{
		UUID:   uid,
		Secret: hashSecret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)

	return token.SignedString(jwtKey)
}

func VerifyToken(code string, token string, jwtKey []byte) (string, error) {
	if code != "access" && code != "refresh" {
		return "", errors.New("code invalid.")
	}

	claim := &Claim{}

	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", errors.New("token invalid.")
	}

	expireDiff := claim.ExpiresAt - time.Now().Unix()
	if expireDiff <= 0 {
		return "", errors.New("token expired.")
	}

	rd_key := fmt.Sprintf("%s_%s", code, claim.UUID)
	rd_val := rd.Get(ctx, rd_key)

	err = utils.CompareHash(rd_val.Val(), claim.Secret)
	if err != nil {
		return "", errors.New("token secret wrong.")
	}

	return claim.UUID, nil
}
