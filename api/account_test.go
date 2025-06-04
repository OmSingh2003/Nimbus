package api 
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OmSingh2003/simple-bank/db/mockdb"
	db "github.com/OmSingh2003/simple-bank/db/sqlc"
	"github.com/OmSingh2003/simple-bank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t) 
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	//build stubs 
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// start test server and send request 
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.ServeHTTP(recorder, request)

	// check response 
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
