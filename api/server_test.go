package api

import (
    "testing"

    mockdb "github.com/OmSingh2003/vaultguard-api/db/mockdb"
    "github.com/stretchr/testify/require"
    "go.uber.org/mock/gomock"
)

func TestNewServer(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mockdb.NewMockStore(ctrl)
    server := NewTestServer(t, store)

    require.NotNil(t, server)
    require.NotNil(t, server.store)
    require.NotNil(t, server.router)

    // Test if routes are properly set up
    routes := server.router.Routes()
    require.NotEmpty(t, routes)

    // Verify specific routes exist
    routeExists := func(method, path string) bool {
        for _, route := range routes {
            if route.Method == method && route.Path == path {
                return true
            }
        }
        return false
    }

    require.True(t, routeExists("POST", "/accounts"))
    require.True(t, routeExists("GET", "/accounts/:id"))
    require.True(t, routeExists("GET", "/accounts"))
}

func TestStartServer(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mockdb.NewMockStore(ctrl)
    server := NewTestServer(t, store)

    // Test with invalid address
    err := server.Start("invalid:address")
    require.Error(t, err)
}
