package jwt

import (
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/env"

	jwtdriver "github.com/golang-jwt/jwt"
)

// TODO
// - maybe pass create a struct and pass the env to a constructor ???

// move this part into a separate area, and don't use token no more inside the usecases

// given an userID
// returns a token
func CreateToken(userID string) (string, error) {

	if userID == "" {
		return "", fmt.Errorf("cannot create token from empty string")
	}
	claims := jwtdriver.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID

	lifespan, err := env.GetString("TOKEN_LIFE_SPAN")
	if err != nil {
		return "", err
	}

	duration, err := time.ParseDuration(lifespan)
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Add(duration).Unix()

	token := jwtdriver.NewWithClaims(jwtdriver.SigningMethodHS256, claims)

	apiSecret, err := env.GetString("TOKEN_API_SECRET")
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(apiSecret))
}

// extract the userID from which the token was generated
// return an error if the token is not valid or expired
func ExtractTokenID(tokenString string) (string, error) {

	token, err := jwtdriver.Parse(tokenString, func(token *jwtdriver.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtdriver.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secret, err := env.GetString("TOKEN_API_SECRET")
		if err != nil {
			return nil, err
		}
		return []byte(secret), nil
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
