package gogocryptohopper_test

import (
	"os"
	"testing"

	"github.com/cryptwire/gogocryptohopper"
)

var ch *gogocryptohopper.GoGoCryptohopper

func TestMain(m *testing.M) {
	endpoint := "https://api.cryptohopper.com/v1"
	var err error
	ch, err = gogocryptohopper.New(endpoint)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestNewWithEndpoint(t *testing.T) {
	endpoint := "https://api.cryptohopper.com/v1"
	gs, err := gogocryptohopper.New(endpoint)
	if err != nil {
		t.Error(err)
	}
	if gs.APIEndpoint != endpoint {
		t.Errorf(
			"Incorrect endpoint set, expected %s, got %s", endpoint, gs.APIEndpoint,
		)
	}
}
