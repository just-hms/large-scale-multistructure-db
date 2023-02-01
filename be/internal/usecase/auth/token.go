package auth

import (
	"fmt"
	"large-scale-multistructure-db/be/config/constants"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// given an userID
// returns a token
func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * constants.TOKEN_HOUR_LIFE_SPAN).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.API_SECRET))
}

// extract the userID from which the token was generated
// return an error if the token is not valid or expired
func ExtractTokenID(tokenString string) (uint, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(constants.API_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userID"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
