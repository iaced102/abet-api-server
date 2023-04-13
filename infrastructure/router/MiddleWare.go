package router

import (
	aAuth "aBet/adapters/auth"
	iAuth "aBet/infrastructure/auth"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
Hàm lấy thông tin token và kiểm tra
- hợp lệ
- Xác thực
*/
func getMiddleWareConfig(authObject *aAuth.AuthObject) middleware.JWTConfig {
	return middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				signingKey := iAuth.GetPublicKey()

				// signingK, e := jwt.ParseRSAPublicKeyFromPEM(signingKey)
				// fmt.Println(e, "signinggggggggggggggggggggggggg")
				return []byte(signingKey), nil
			}
			// var re = regexp.MustCompile(`[a-z0-9A-X]*::`)
			// newAuth := re.ReplaceAllString(auth, "")
			token, err := jwt.Parse(auth, keyFunc)
			fmt.Println(err, "rrrrrrrrrrrrrrrrrrrrrrr")
			fmt.Println(token.Valid)
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			claims, _ := json.Marshal(token.Claims)

			*authObject, _ = iAuth.NewAuthObject(claims, "Bearer "+token.Raw)

			return token, nil
		},
	}
}
