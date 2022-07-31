package auth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("sub_jwt_key"))

type JWTClaim struct {
	UserId   string `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userId string, email string, name string) (tokenString string, expiresAt int64, err error) {
	expiresAt = time.Now().Add(1 * time.Hour).Unix()
	claims := &JWTClaim{
		UserId:   userId,
		Email:    email,
		UserName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {

	claims, ok := GetClaimFromToken(signedToken)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func GetClaimFromToken(signedToken string) (*JWTClaim, bool) {
	signedToken = removeWordBearer(signedToken)
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, false
	}
	claims, ok := token.Claims.(*JWTClaim)
	return claims, ok
}

func removeWordBearer(tokenString string) string {
	return strings.ReplaceAll(tokenString, "Bearer ", "")
}
