package tests

import (
	"testing"

	"github.com/lucidconnect/inverse/database"
)

var db = database.SetupDB("postgres://localhost:5432/lucid_nft_test?sslmode=disable")

func TestDelete(t *testing.T) {
	if err := db.RemoveFarcasterCriteria("44643b76-cf35-4f06-abef-066524fb6431"); err != nil {
		t.Log(err)
		t.Fail()
	}
}
