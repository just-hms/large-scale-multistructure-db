package tokenapi

import (
	"fmt"
	"time"

	jwtdriver "github.com/golang-jwt/jwt"
)

// TODO
// - maybe pass create a struct and pass the env to a constructor ???

// move this part into a separate area, and don't use token no more inside the usecases

// given an userID
// returns a token

type JWT struct {
	tokenLifespan time.Duration
	apisecret     string
}

func New(tokenLifespan time.Duration, apisecret string) *JWT {
	return &JWT{
		tokenLifespan: tokenLifespan,
		apisecret:     apisecret,
	}
}

func (j *JWT) CreateToken(userID string) (string, error) {

	if userID == "" {
		return "", fmt.Errorf("cannot create token from empty string")
	}
	claims := jwtdriver.MapClaims{}

	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(j.tokenLifespan).Unix()

	token := jwtdriver.NewWithClaims(jwtdriver.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.apisecret))
}

// extract the userID from which the token was generated
// return an error if the token is not valid or expired
func (j *JWT) ExtractTokenID(tokenString string) (string, error) {

	token, err := jwtdriver.Parse(tokenString, func(token *jwtdriver.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtdriver.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.apisecret), nil
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
