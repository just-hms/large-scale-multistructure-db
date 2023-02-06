package jwt

import (
	"fmt"
	"large-scale-multistructure-db/be/config/constants"
	"time"

	jwtdriver "github.com/golang-jwt/jwt"
)

// move this part into a separate area, and don't use token no more inside the usecases

// given an userID
// returns a token
func CreateToken(userID string) (string, error) {
	claims := jwtdriver.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * constants.TOKEN_HOUR_LIFE_SPAN).Unix()

	token := jwtdriver.NewWithClaims(jwtdriver.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.API_SECRET))
}

// extract the userID from which the token was generated
// return an error if the token is not valid or expired
func ExtractTokenID(tokenString string) (string, error) {

	token, err := jwtdriver.Parse(tokenString, func(token *jwtdriver.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtdriver.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(constants.API_SECRET), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwtdriver.MapClaims)

	if !ok || !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	userID, ok := claims["userID"].(string)

	if ok {
		return userID, nil
	}

	return "", fmt.Errorf("Invalid token")

}
