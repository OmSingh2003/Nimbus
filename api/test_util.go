package api

import (
	"testing"
	"time"

	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/util"
	mockwk "github.com/OmSingh2003/vaultguard-api/worker/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	
	// Create a mock task distributor for testing
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	taskDistributor := mockwk.NewMockTaskDistributor(ctrl)
	
	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

