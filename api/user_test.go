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

func getRandomUser(t *testing.T) *db.User {
	password, err := util.HashPassword("secret")
	require.NoError(t, err)
	return &db.User{
		Username:       util.GetRandomString(6),
		HashedPassword: password,
		FullName:       util.GetRandomString(10),
		Email:          "hzshaoyuheng@gmail.com",
	}
}

type EqCreateUserParamsMatcher struct {
	user *db.User
}

// Matches returns whether x is a match.
func (e EqCreateUserParamsMatcher) Matches(x interface{}) bool {
	params := x.(db.CreateUserParams)
	if err := util.CheckPassword("secret", params.HashedPassword); err != nil {
		return false
	}
	flag := e.user.Username == params.Username
	flag = flag && (e.user.FullName == params.FullName)
	flag = flag && (e.user.Email == params.Email)
	return flag
}

// String describes what the matcher matches.
func (e EqCreateUserParamsMatcher) String() string {
	return "EqCreateUserParamsMatcher"
}

func EqCreateUserParams(usr *db.User) EqCreateUserParamsMatcher {
	return EqCreateUserParamsMatcher{usr}
}

func TestCreateUserAPI(t *testing.T) {
	t.Run("POST200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)

		usr := getRandomUser(t)

		store.EXPECT().
			CreateUser(gomock.Any(), EqCreateUserParams(usr)).
			Times(1).
			Return(*usr, nil)

		server := NewServer(store)
		recorder := httptest.NewRecorder()

		req := createUserRequest{
			Username: usr.Username,
			Password: "secret",
			FullName: usr.FullName,
			Email:    usr.Email,
		}
		data, err := json.Marshal(req)
		require.NoError(t, err)

		url := fmt.Sprintf("/user")
		request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		server.router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusOK, recorder.Code)
		requireMatchUser(t, usr, recorder.Body)
	})
}

func requireMatchUser(t *testing.T, usr *db.User, buffer *bytes.Buffer) {
	data, err := ioutil.ReadAll(buffer)
	require.NoError(t, err)
	var gotRsp userResponse
	err = json.Unmarshal(data, &gotRsp)
	require.NoError(t, err)
	require.Equal(t, usr.Username, gotRsp.Username)
	require.Equal(t, usr.FullName, gotRsp.FullName)
	require.Equal(t, usr.Email, gotRsp.Email)
}
