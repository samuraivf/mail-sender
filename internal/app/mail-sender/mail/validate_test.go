package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateEmail(t *testing.T) {
	require.True(t, validateEmail("email@gmail.com"))
	require.False(t, validateEmail("email"))
}