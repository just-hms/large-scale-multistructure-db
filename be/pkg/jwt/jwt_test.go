package jwt_test

import (
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/jwt"

	"github.com/stretchr/testify/assert"
)

func TestInvertibility(t *testing.T) {

	// tokenize an ID
	startID := "1"
	token, err := jwt.CreateToken(startID)

	// check that there is no error
	assert.Nil(t, err)

	// get the ID from the token
	resID, err := jwt.ExtractTokenID(token)
	assert.Nil(t, err)

	// check if the extracted token is the same as the start
	assert.Equal(t, startID, resID)

}

// TODO : test
// - the expiration time
// - firm with wrong secret
