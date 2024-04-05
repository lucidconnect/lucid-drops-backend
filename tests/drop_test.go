package tests

import (
	"testing"

	"github.com/lucidconnect/inverse/drops"
	uuid "github.com/satori/go.uuid"
)

func TestCreateDrop(t *testing.T) {
	testDrop := &drops.Drop{
		CreatorID:      uuid.FromStringOrNil("2298ba66-f1ec-4ca7-9730-727d09e08737"),
		Name:           "test drop",
		CreatorAddress: "0x8605fFD3382850228135A4A8a780a740e9251A43",
	}

	testItem := &drops.Item{
		Name:    "test item",
		TokenID: int64(1),
	}

	if err := db.CreateDrop(testDrop, testItem); err != nil {
		t.Log(err)
		t.Fail()
	}
}
