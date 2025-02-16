package auf

import (
	"server/config"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWT(t *testing.T) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":    strconv.Itoa(2),
		"expiredAt": time.Now().Unix(),
	})
	tokenStr, err := token.SignedString([]byte(config.Env.JWTSecret))
	if err != nil {
		t.Fatal("error generation token", err)
	}
	t.Log(tokenStr)

	_, err = ParseJWT([]byte(config.Env.JWTSecret), tokenStr)
	if err == nil {
		t.Fatal("token is expired, error is nil, should return error")
	}

	goodtoken, err := CreateJWT([]byte(config.Env.JWTSecret), 2)
	if err != nil {
		t.Fatal("error generation token", err)
	}

	userid, err := ParseJWT([]byte(config.Env.JWTSecret), goodtoken)
	if err != nil {
		t.Fatal("error pasrsing token", err)
	}

	if userid != 2 {
		t.Fatal("id is not 2, wrong parsing")
	}

}
