package jwt

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTModel struct {
	ThirdParty  string  `json:"third_party"`
	ExpiredTime float64 `json:"expired_time"`
}

const (
	hmacSecret = "hahahahahhawkwkwkwkwkw"
)

func GenerateToken(userId uint) string {
	user_id := fmt.Sprint(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"third_party":  user_id,
		"expired_time": time.Now().Add(time.Hour * 24 * 30 * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return err.Error()
	} else {
		return tokenString
	}
}

func CheckToken(unparsedToken string) (bool, string) {

	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})
	if err != nil {
		return false, err.Error()
	}
	raw, err := json.Marshal(token.Claims)
	if err != nil {
		return false, err.Error()
	}
	var tokenParsing *JWTModel
	err = json.Unmarshal(raw, &tokenParsing)
	if err != nil {
		return false, err.Error()
	} else {
		if time.Now().Unix() > int64(tokenParsing.ExpiredTime) {
			return false, "Token has been expired."
		} else {
			return true, tokenParsing.ThirdParty
		}
	}
}
