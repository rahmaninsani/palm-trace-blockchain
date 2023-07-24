package helper

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetAsset(ctx contractapi.TransactionContextInterface, id string) ([]byte, error) {
	asset, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state")
	}

	return asset, nil
}
