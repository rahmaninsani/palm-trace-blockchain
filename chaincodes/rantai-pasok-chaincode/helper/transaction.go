package helper

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// func CreateCompositeKeyWithUnderscore(ctx contractapi.TransactionContextInterface, objectType string, components []string) (string, error) {
// 	separator := "_"
// 	compositeKey, err := ctx.GetStub().CreateCompositeKey(objectType, components)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create composite key: %v", err)
// 	}
// 	compositeKeyWithUnderscore := compositeKey[:len(compositeKey)-1] + separator + string(compositeKey[len(compositeKey)-1])
// 	return compositeKeyWithUnderscore, nil
// }
