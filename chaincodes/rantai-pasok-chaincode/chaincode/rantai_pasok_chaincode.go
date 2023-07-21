package chaincode

import (
	"encoding/json"
	"fmt"

	"rantai-pasok-chaincode/helper"
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RantaiPasokChaincode struct {
	contractapi.Contract
}

func (c *RantaiPasokChaincode) Create(ctx contractapi.TransactionContextInterface, payload string) (*web.RantaiPasokResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have dinas.user affiliation/role")
	}

	var rantaiPasok domain.RantaiPasok
	err = json.Unmarshal([]byte(payload), &rantaiPasok)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, rantaiPasok.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", rantaiPasok.Id)
	}

	referensiHargaJSON, err := json.Marshal(rantaiPasok)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal asset: %v", err)
	}

	err = ctx.GetStub().PutState(rantaiPasok.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	RantaiPasokResponse := web.RantaiPasokResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    rantaiPasok.Id,
		IdDinas:               rantaiPasok.IdDinas,
		UmurTanam:             rantaiPasok.UmurTanam,
		Harga:                 rantaiPasok.Harga,
		TanggalPembaruan:      rantaiPasok.TanggalPembaruan,
	}

	return &RantaiPasokResponse, nil
}

func (c *RantaiPasokChaincode) Update(ctx contractapi.TransactionContextInterface, payload string) (*web.RantaiPasokResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to update asset, does not have dinas.user affiliation/role")
	}

	var rantaiPasok domain.RantaiPasok
	err = json.Unmarshal([]byte(payload), &rantaiPasok)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, rantaiPasok.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", rantaiPasok.Id)
	}

	referensiHargaJSON, err := json.Marshal(rantaiPasok)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated asset: %v", err)
	}

	err = ctx.GetStub().PutState(rantaiPasok.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	RantaiPasokResponse := web.RantaiPasokResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    rantaiPasok.Id,
		IdDinas:               rantaiPasok.IdDinas,
		UmurTanam:             rantaiPasok.UmurTanam,
		Harga:                 rantaiPasok.Harga,
		TanggalPembaruan:      rantaiPasok.TanggalPembaruan,
	}

	return &RantaiPasokResponse, nil
}

func (c *RantaiPasokChaincode) GetAll(ctx contractapi.TransactionContextInterface) ([]*web.RantaiPasokResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var RantaiPasokResponses []*web.RantaiPasokResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var rantaiPasok domain.RantaiPasok
		err = json.Unmarshal(response.Value, &rantaiPasok)
		if err != nil {
			return nil, err
		}

		RantaiPasokResponse := web.RantaiPasokResponse{
			Id:               rantaiPasok.Id,
			IdDinas:          rantaiPasok.IdDinas,
			UmurTanam:        rantaiPasok.UmurTanam,
			Harga:            rantaiPasok.Harga,
			TanggalPembaruan: rantaiPasok.TanggalPembaruan,
		}
		RantaiPasokResponses = append(RantaiPasokResponses, &RantaiPasokResponse)
	}

	return RantaiPasokResponses, nil
}

func (c *RantaiPasokChaincode) GetHistoryById(ctx contractapi.TransactionContextInterface, id string) ([]*web.RantaiPasokResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for key %s: %v", id, err)
	}
	defer resultsIterator.Close()

	var RantaiPasokResponses []*web.RantaiPasokResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", id, err)
		}

		var rantaiPasok domain.RantaiPasok
		err = json.Unmarshal(response.Value, &rantaiPasok)
		if err != nil {
			return nil, err
		}

		RantaiPasokResponse := web.RantaiPasokResponse{
			IdTransaksiBlockchain: response.GetTxId(),
			Id:                    rantaiPasok.Id,
			IdDinas:               rantaiPasok.IdDinas,
			UmurTanam:             rantaiPasok.UmurTanam,
			Harga:                 rantaiPasok.Harga,
			TanggalPembaruan:      rantaiPasok.TanggalPembaruan,
		}
		RantaiPasokResponses = append(RantaiPasokResponses, &RantaiPasokResponse)
	}

	return RantaiPasokResponses, nil
}
