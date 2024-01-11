package tools

import (
	"time"
)

type MockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "12345",
		Username:  "alex",
	},
}

var mockCoinsDetails = map[string]CoinDetails{
	"alex": {
		Coins:    200,
		Username: "alex",
	},
}

func (d *MockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *MockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinsDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *MockDB) SetUpDatabase() error {
	return nil
}
