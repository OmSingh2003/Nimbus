package api

import (
	"testing"
	"time"

	db "github.com/OmSingh2003/simple-bank/db/sqlc"
	"github.com/OmSingh2003/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

