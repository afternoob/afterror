package afterror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_IsTypeEqual_True(t *testing.T) {
	err := NotFound("NotFoundUser", "User not found")
	errUserNotFound := NotFound("NotFoundUser", "User2 not found")

	require.True(t, err.IsTypeEqual(errUserNotFound))
}

func TestError_IsTypeEqual_False(t *testing.T) {
	err := NotFound("NotFoundUser", "user not found")
	errUnableGetUser := InternalServer("UnableGetUser", "Unable get user")

	require.False(t, err.IsTypeEqual(errUnableGetUser))
}

func TestError_Error(t *testing.T) {
	errUnauthorized := Unauthorized("Unauthorized", "Access is denied")

	require.Equal(t, "Unauthorized: Access is denied", errUnauthorized.Error())
}

func TestError_ErrorWithNoMessage(t *testing.T) {
	errForbidden := Forbidden("Forbidden", "")

	require.Equal(t, "Unknown Error", errForbidden.Error())
}

func TestError_Wrap(t *testing.T) {
	defaultErr := errors.New("FirstName is required")
	errUnableGetUser := BadRequest("InvalidInput", "Invalid request")

	_ = errUnableGetUser.Wrap(defaultErr)
	require.Equal(t, defaultErr, errUnableGetUser.Err)
}
