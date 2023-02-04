package jwt

import (
	"fmt"
	"large-scale-multistructure-db/be/config/constants"
	"strconv"
	"time"

	jwtdriver "github.com/golang-jwt/jwt"
)

// move this part into a separate area, and don't use token no more inside the usecases

// given an userID
// returns a token
func CreateToken(userID uint) (string, error) {
	claims := jwtdriver.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * constants.TOKEN_HOUR_LIFE_SPAN).Unix()

	token := jwtdriver.NewWithClaims(jwtdriver.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.API_SECRET))
}

// extract the userID from which the token was generated
// return an error if the token is not valid or expired
func ExtractTokenID(tokenString string) (uint, error) {

	token, err := jwtdriver.Parse(tokenString, func(token *jwtdriver.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtdriver.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(constants.API_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwtdriver.MapClaims)

	if ok && token.Valid {

		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userID"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
