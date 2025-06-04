package api

import (
    "bytes"
    "database/sql"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    mockdb "github.com/OmSingh2003/simple-bank/db/mockdb"
    db "github.com/OmSingh2003/simple-bank/db/sqlc"
    "github.com/OmSingh2003/simple-bank/util"
    "github.com/stretchr/testify/require"
    "go.uber.org/mock/gomock"
)

func TestGetAccountApi(t *testing.T) {
    account := randomAccount()

    testCases := []struct {
        name          string
        accountID     int64
        buildStubs    func(store *mockdb.MockStore)
        checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
    }{
        {
            name:      "OK",
            accountID: account.ID,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    GetAccount(gomock.Any(), gomock.Eq(account.ID)).
                    Times(1).
                    Return(account, nil)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusOK, recorder.Code)
                requireBodyMatchAccount(t, recorder.Body, account)
            },
        },
        {
            name:      "NotFound",
            accountID: account.ID,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    GetAccount(gomock.Any(), gomock.Eq(account.ID)).
                    Times(1).
                    Return(db.Account{}, sql.ErrNoRows)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusNotFound, recorder.Code)
            },
        },
        {
            name:      "InvalidID",
            accountID: 0,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    GetAccount(gomock.Any(), gomock.Any()).
                    Times(0)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusBadRequest, recorder.Code)
            },
        },
        {
            name:      "DatabaseError",
            accountID: account.ID,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    GetAccount(gomock.Any(), gomock.Eq(account.ID)).
                    Times(1).
                    Return(db.Account{}, sql.ErrConnDone)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusInternalServerError, recorder.Code)
            },
        },
    }

    for i := range testCases {
        tc := testCases[i]

        t.Run(tc.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            store := mockdb.NewMockStore(ctrl)
            tc.buildStubs(store)

            server := NewServer(store)
            recorder := httptest.NewRecorder()

            url := fmt.Sprintf("/accounts/%d", tc.accountID)
            request, err := http.NewRequest(http.MethodGet, url, nil)
            require.NoError(t, err)

            server.ServeHTTP(recorder, request)
            tc.checkResponse(t, recorder)
        })
    }
}

func TestCreateAccountApi(t *testing.T) {
    account := randomAccount()

    testCases := []struct {
        name          string
        body          map[string]interface{}
        buildStubs    func(store *mockdb.MockStore)
        checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
    }{
        {
            name: "OK",
            body: map[string]interface{}{
                "owner":    account.Owner,
                "currency": account.Currency,
            },
            buildStubs: func(store *mockdb.MockStore) {
                arg := db.CreateAccountParams{
                    Owner:    account.Owner,
                    Currency: account.Currency,
                    Balance:  0,
                }

                store.EXPECT().
                    CreateAccount(gomock.Any(), gomock.Eq(arg)).
                    Times(1).
                    Return(account, nil)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusOK, recorder.Code)
                requireBodyMatchAccount(t, recorder.Body, account)
            },
        },
        {
            name: "InvalidCurrency",
            body: map[string]interface{}{
                "owner":    account.Owner,
                "currency": "invalid",
            },
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    CreateAccount(gomock.Any(), gomock.Any()).
                    Times(0)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusBadRequest, recorder.Code)
            },
        },
        {
            name: "DatabaseError",
            body: map[string]interface{}{
                "owner":    account.Owner,
                "currency": account.Currency,
            },
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    CreateAccount(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(db.Account{}, sql.ErrConnDone)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusInternalServerError, recorder.Code)
            },
        },
    }

    for i := range testCases {
        tc := testCases[i]

        t.Run(tc.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            store := mockdb.NewMockStore(ctrl)
            tc.buildStubs(store)

            server := NewServer(store)
            recorder := httptest.NewRecorder()

            data, err := json.Marshal(tc.body)
            require.NoError(t, err)

            url := "/accounts"
            request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
            require.NoError(t, err)

            server.ServeHTTP(recorder, request)
            tc.checkResponse(t, recorder)
        })
    }
}

func TestListAccountsApi(t *testing.T) {
    n := 5
    accounts := make([]db.Account, n)
    for i := 0; i < n; i++ {
        accounts[i] = randomAccount()
    }

    testCases := []struct {
        name          string
        pageID       int
        pageSize     int
        buildStubs    func(store *mockdb.MockStore)
        checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
    }{
        {
            name:     "OK",
            pageID:   1,
            pageSize: n,
            buildStubs: func(store *mockdb.MockStore) {
                arg := db.ListAccountsParams{
                    Limit:  int32(n),
                    Offset: 0,
                }

                store.EXPECT().
                    ListAccounts(gomock.Any(), gomock.Eq(arg)).
                    Times(1).
                    Return(accounts, nil)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusOK, recorder.Code)
                requireBodyMatchAccounts(t, recorder.Body, accounts)
            },
        },
        {
            name:     "InvalidPageID",
            pageID:   -1,
            pageSize: n,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    ListAccounts(gomock.Any(), gomock.Any()).
                    Times(0)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusBadRequest, recorder.Code)
            },
        },
        {
            name:     "InvalidPageSize",
            pageID:   1,
            pageSize: 100000,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    ListAccounts(gomock.Any(), gomock.Any()).
                    Times(0)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusBadRequest, recorder.Code)
            },
        },
        {
            name:     "DatabaseError",
            pageID:   1,
            pageSize: n,
            buildStubs: func(store *mockdb.MockStore) {
                store.EXPECT().
                    ListAccounts(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return([]db.Account{}, sql.ErrConnDone)
            },
            checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                require.Equal(t, http.StatusInternalServerError, recorder.Code)
            },
        },
    }

    for i := range testCases {
        tc := testCases[i]

        t.Run(tc.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            store := mockdb.NewMockStore(ctrl)
            tc.buildStubs(store)

            server := NewServer(store)
            recorder := httptest.NewRecorder()

            url := fmt.Sprintf("/accounts?page_id=%d&page_size=%d", tc.pageID, tc.pageSize)
            request, err := http.NewRequest(http.MethodGet, url, nil)
            require.NoError(t, err)

            server.ServeHTTP(recorder, request)
            tc.checkResponse(t, recorder)
        })
    }
}

func randomAccount() db.Account {
    return db.Account{
        ID:       util.RandomInt(1, 1000),
        Owner:    util.RandomOwner(),
        Balance:  util.RandomMoney(),
        Currency: util.RandomCurrency(),
    }
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
    data, err := io.ReadAll(body)
    require.NoError(t, err)

    var gotAccount db.Account
    err = json.Unmarshal(data, &gotAccount)
    require.NoError(t, err)
    require.Equal(t, account, gotAccount)
}

func requireBodyMatchAccounts(t *testing.T, body *bytes.Buffer, accounts []db.Account) {
    data, err := io.ReadAll(body)
    require.NoError(t, err)

    var gotAccounts []db.Account
    err = json.Unmarshal(data, &gotAccounts)
    require.NoError(t, err)
    require.Equal(t, accounts, gotAccounts)
}
