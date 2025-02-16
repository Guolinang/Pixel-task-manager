package auf

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userid int) (string, error) {

	exp := time.Hour * 24 * 30
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":    strconv.Itoa(userid),
		"expiredAt": time.Now().Add(exp).Unix(),
	})
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return tokenStr, nil
}

func ParseJWT(secret []byte, token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("authentication failed: %v", err)

	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {

		if expiredAt, ok := (claims["expiredAt"]).(float64); ok {
			tm := time.Unix(int64(expiredAt), 0)
			if time.Now().After(tm) {
				return 0, fmt.Errorf("token expired")
			}
		}
		if str, ok := (claims["userid"]).(string); ok {
			n, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("authentication failed: %v", err)
			}
			return int(n), nil
		}
	}
	return 0, fmt.Errorf("authentication failed")
}
