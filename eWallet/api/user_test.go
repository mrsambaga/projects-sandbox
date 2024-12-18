package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/mrsambaga/projects-sandbox/eWallet/db/mock"
	db "github.com/mrsambaga/projects-sandbox/eWallet/db/sqlc"
	"github.com/mrsambaga/projects-sandbox/eWallet/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserAPI(t *testing.T) {
	user, password := createRandomUser(t)

	testCases := []struct{
		name string
		body gin.H
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":    	user.Username,
				"password": 	password,
				"full_name": 	user.FullName,
				"email": 		user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
			arg := db.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
			}
			store.EXPECT().
				CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
				Times(1).
				Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "Invalid Email",
			body: gin.H{
				"username":    	user.Username,
				"password": 	password,
				"full_name": 	user.FullName,
				"email": 		"invalidemail",
			},
			buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Password Too Short",
			body: gin.H{
				"username":    	user.Username,
				"password": 	"x",
				"full_name": 	user.FullName,
				"email": 		user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Duplicate User",
			body: gin.H{
				"username":    	user.Username,
				"password": 	password,
				"full_name": 	user.FullName,
				"email": 		user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Times(1).
				Return(db.User{}, db.ErrUniqueViolation)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"username":    	user.Username,
				"password": 	password,
				"full_name": 	user.FullName,
				"email": 		user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Times(1).
				Return(db.User{}, sql.ErrConnDone)
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
		
			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
		
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func createRandomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6);
	hashedPassword, err := util.HashPassword(password);
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:  		util.RandomOwner(),
		Email: 			util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
