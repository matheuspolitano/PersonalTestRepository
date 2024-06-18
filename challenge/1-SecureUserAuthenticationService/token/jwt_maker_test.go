package token

import (
	"testing"
	"time"

	"github.com/matheuspolitano/PersonalTestRepository/challenge/1-SecureUserAuthenticationService/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	secretKey := utils.RandString(20)
	username := utils.RandString(10)
	role := utils.RandString(10)

	maker := JWTMaker{secretKey: secretKey}

	payload, err := NewPayload(username, role, time.Minute*15)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	token, payload, err := maker.CreateToken(payload)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, token)
}

func TestCheckToken(t *testing.T) {
	secretKey := utils.RandString(20)
	username := utils.RandString(10)
	role := utils.RandString(10)

	maker := JWTMaker{secretKey: secretKey}

	payloadCreated, err := NewPayload(username, role, time.Minute*15)
	require.NoError(t, err)
	require.NotEmpty(t, payloadCreated)

	token, payloadCreated, err := maker.CreateToken(payloadCreated)
	require.NoError(t, err)
	require.NotEmpty(t, payloadCreated)
	require.NotEmpty(t, token)

	payloadFromToken, err := maker.CheckToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payloadFromToken)

	require.Equal(t, payloadCreated.ID, payloadFromToken.ID)
	require.Equal(t, payloadCreated.Username, payloadFromToken.Username)
	require.Equal(t, payloadCreated.Role, payloadFromToken.Role)
}

func TestCheckTokenExpired(t *testing.T) {
	secretKey := utils.RandString(20)
	username := utils.RandString(10)
	role := utils.RandString(10)

	maker := JWTMaker{secretKey: secretKey}

	payloadCreated, err := NewPayload(username, role, time.Millisecond*1)
	require.NoError(t, err)
	require.NotEmpty(t, payloadCreated)

	token, payloadCreated, err := maker.CreateToken(payloadCreated)
	require.NoError(t, err)
	require.NotEmpty(t, payloadCreated)
	require.NotEmpty(t, token)

	time.Sleep(time.Millisecond * 10)
	_, err = maker.CheckToken(token)
	require.Error(t, err, "error token is expired")

}
