package main

import (
	"log"

	"referensi-harga-chaincode/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	referensiHargaChaincode, err := contractapi.NewChaincode(&chaincode.ReferensiHargaChaincode{})
	if err != nil {
		log.Panicf("Error creating ReferensiHargaChaincode: %v", err)
	}

	if err := referensiHargaChaincode.Start(); err != nil {
		log.Panicf("Error starting ReferensiHargaChaincode: %v", err)
	}
}
