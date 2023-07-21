package chaincode

import (
	"encoding/json"
	"fmt"

	"supply-chain-chaincode/helper"
	"supply-chain-chaincode/model/domain"
	"supply-chain-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SupplyChainChaincode struct {
	contractapi.Contract
}

func (c *SupplyChainChaincode) Create(ctx contractapi.TransactionContextInterface, payload string) (*web.SupplyChainResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have dinas.user affiliation/role")
	}

	var supplyChain domain.SupplyChain
	err = json.Unmarshal([]byte(payload), &supplyChain)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, supplyChain.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", supplyChain.Id)
	}

	referensiHargaJSON, err := json.Marshal(supplyChain)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal asset: %v", err)
	}

	err = ctx.GetStub().PutState(supplyChain.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	SupplyChainResponse := web.SupplyChainResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    supplyChain.Id,
		IdDinas:               supplyChain.IdDinas,
		UmurTanam:             supplyChain.UmurTanam,
		Harga:                 supplyChain.Harga,
		TanggalPembaruan:      supplyChain.TanggalPembaruan,
	}

	return &SupplyChainResponse, nil
}

func (c *SupplyChainChaincode) Update(ctx contractapi.TransactionContextInterface, payload string) (*web.SupplyChainResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to update asset, does not have dinas.user affiliation/role")
	}

	var supplyChain domain.SupplyChain
	err = json.Unmarshal([]byte(payload), &supplyChain)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, supplyChain.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", supplyChain.Id)
	}

	referensiHargaJSON, err := json.Marshal(supplyChain)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated asset: %v", err)
	}

	err = ctx.GetStub().PutState(supplyChain.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	SupplyChainResponse := web.SupplyChainResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    supplyChain.Id,
		IdDinas:               supplyChain.IdDinas,
		UmurTanam:             supplyChain.UmurTanam,
		Harga:                 supplyChain.Harga,
		TanggalPembaruan:      supplyChain.TanggalPembaruan,
	}

	return &SupplyChainResponse, nil
}

func (c *SupplyChainChaincode) GetAll(ctx contractapi.TransactionContextInterface) ([]*web.SupplyChainResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var SupplyChainResponses []*web.SupplyChainResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var supplyChain domain.SupplyChain
		err = json.Unmarshal(response.Value, &supplyChain)
		if err != nil {
			return nil, err
		}

		SupplyChainResponse := web.SupplyChainResponse{
			Id:               supplyChain.Id,
			IdDinas:          supplyChain.IdDinas,
			UmurTanam:        supplyChain.UmurTanam,
			Harga:            supplyChain.Harga,
			TanggalPembaruan: supplyChain.TanggalPembaruan,
		}
		SupplyChainResponses = append(SupplyChainResponses, &SupplyChainResponse)
	}

	return SupplyChainResponses, nil
}

func (c *SupplyChainChaincode) GetHistoryById(ctx contractapi.TransactionContextInterface, id string) ([]*web.SupplyChainResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for key %s: %v", id, err)
	}
	defer resultsIterator.Close()

	var SupplyChainResponses []*web.SupplyChainResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", id, err)
		}

		var supplyChain domain.SupplyChain
		err = json.Unmarshal(response.Value, &supplyChain)
		if err != nil {
			return nil, err
		}

		SupplyChainResponse := web.SupplyChainResponse{
			IdTransaksiBlockchain: response.GetTxId(),
			Id:                    supplyChain.Id,
			IdDinas:               supplyChain.IdDinas,
			UmurTanam:             supplyChain.UmurTanam,
			Harga:                 supplyChain.Harga,
			TanggalPembaruan:      supplyChain.TanggalPembaruan,
		}
		SupplyChainResponses = append(SupplyChainResponses, &SupplyChainResponse)
	}

	return SupplyChainResponses, nil
}
