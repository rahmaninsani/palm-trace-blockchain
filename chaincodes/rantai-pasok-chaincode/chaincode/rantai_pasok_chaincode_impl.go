package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RantaiPasokChaincodeImpl struct {
	contractapi.Contract
}

func NewRantaiPasokChaincode() contractapi.ContractInterface {
	return &RantaiPasokChaincodeImpl{
		Contract: contractapi.Contract{},
	}
}
