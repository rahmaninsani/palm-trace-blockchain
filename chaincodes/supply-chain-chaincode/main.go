package main

import (
	"log"

	"supply-chain-chaincode/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	supplyChainChaincode, err := contractapi.NewChaincode(&chaincode.SupplyChainChaincode{})
	if err != nil {
		log.Panicf("Error creating SupplyChainChaincode: %v", err)
	}

	if err := supplyChainChaincode.Start(); err != nil {
		log.Panicf("Error starting SupplyChainChaincode: %v", err)
	}
}
