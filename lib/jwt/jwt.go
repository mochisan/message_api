package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var base = "4yChANSQCUac9wGKp62-Cjpw6ThE2t5HVUuk7Kj8"

// UserIDClaim struct
type UserIDClaim struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// CreateToken function
func CreateToken(userID uint) string {
	mySigningKey := []byte(base)

	claims := UserIDClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: 0,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	return ss
}

// FetchUserID function
func FetchUserID(tokenString string) uint {

	token, err := jwt.ParseWithClaims(tokenString, &UserIDClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(base), nil
	})
	if err != nil {
		return 0
	}

	claims, ok := token.Claims.(*UserIDClaim)

	if ok && token.Valid {
		return claims.UserID
	}

	return 0
}
