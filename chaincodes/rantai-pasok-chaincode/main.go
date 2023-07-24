package main

import (
	"log"

	"rantai-pasok-chaincode/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	chaincodes := []contractapi.ContractInterface{
		chaincode.NewKebunChaincode(),
		// chaincode.NewRantaiPasokChaincode(),
	}

	rantaiPasokChaincode, err := contractapi.NewChaincode(chaincodes...)

	if err != nil {
		log.Panicf("Error creating RantaiPasokChaincode: %v", err)
	}

	if err := rantaiPasokChaincode.Start(); err != nil {
		log.Panicf("Error starting RantaiPasokChaincode: %v", err)
	}
}
