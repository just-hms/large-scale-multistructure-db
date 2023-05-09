package tokenapi_test

import (
	"testing"
	"time"

	jwt "github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/tokenapi"
	"github.com/stretchr/testify/assert"
)

func TestInvertibility(t *testing.T) {

	j := jwt.New(time.Hour, "api_secret")
	// tokenize an ID
	startID := "1"
	token, err := j.CreateToken(startID)

	// check that there is no error
	assert.Nil(t, err)

	// get the ID from the token
	resID, err := j.ExtractTokenID(token)
	assert.Nil(t, err)

	// check if the extracted token is the same as the start
	assert.Equal(t, startID, resID)

}

// TODO : test
// - the expiration time
// - firm with wrong secret
