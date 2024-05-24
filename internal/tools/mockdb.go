package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username: "alex",
	},
	"john": {
		AuthToken: "5678",
		Username: "john",
	},
	"jane": {
		AuthToken: "91011",
		Username: "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"john": {
		Coins: 200,
		Username: "john",
	},
	"jane": {
		Coins: 300,
		Username: "jane",
	},
}

// now we need to imp. getUserLoginDetails and getUserCoinDetails

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
