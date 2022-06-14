package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/yuhengfdada/go-bank/db"
	mockdb "github.com/yuhengfdada/go-bank/db/mock"
	"github.com/yuhengfdada/go-bank/util"
)

func getRandomAccount() *db.Account {
	return &db.Account{
		ID:       util.GetRandomInt64(100, 1000),
		Owner:    util.GetRandomString(10),
		Balance:  util.GetRandomInt64(100, 1000),
		Currency: util.GetRandomString(3),
	}
}

func TestGetAccountAPI(t *testing.T) {
	t.Run("GET200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)

		acc := getRandomAccount()

		store.EXPECT().
			GetAccount(gomock.Any(), gomock.Eq(acc.ID)).
			Times(1).
			Return(*acc, nil)

		server := NewServer(store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/accounts/%d", acc.ID)
		request := httptest.NewRequest(http.MethodGet, url, recorder.Body)
		server.router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusOK, recorder.Code)
		requireMatchAccount(t, acc, recorder.Body)
	})
	t.Run("GET400", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)

		store.EXPECT().
			GetAccount(gomock.Any(), gomock.Any()).
			Times(0)

		server := NewServer(store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/accounts/0")
		request := httptest.NewRequest(http.MethodGet, url, recorder.Body)
		server.router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func requireMatchAccount(t *testing.T, acc *db.Account, buffer *bytes.Buffer) {
	data, err := ioutil.ReadAll(buffer)
	require.NoError(t, err)
	var gotAcc db.Account
	err = json.Unmarshal(data, &gotAcc)
	require.NoError(t, err)
	require.Equal(t, *acc, gotAcc)
}
