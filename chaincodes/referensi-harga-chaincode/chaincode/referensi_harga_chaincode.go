package chaincode

import (
	"encoding/json"
	"fmt"

	"referensi-harga-chaincode/helper"
	"referensi-harga-chaincode/model/domain"
	"referensi-harga-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ReferensiHargaChaincode struct {
	contractapi.Contract
}

func (c *ReferensiHargaChaincode) Create(ctx contractapi.TransactionContextInterface, payload string) (*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have dinas.user affiliation/role")
	}

	var referensiHargaCreateRequest web.ReferensiHargaCreateRequest
	err = json.Unmarshal([]byte(payload), &referensiHargaCreateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, referensiHargaCreateRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", referensiHargaCreateRequest.Id)
	}

	referensiHarga := domain.ReferensiHarga{
		Id:        referensiHargaCreateRequest.Id,
		IdDinas:   referensiHargaCreateRequest.IdDinas,
		UmurTanam: referensiHargaCreateRequest.UmurTanam,
		Harga:     referensiHargaCreateRequest.Harga,
	}

	referensiHargaJSON, err := json.Marshal(referensiHarga)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal asset: %v", err)
	}

	err = ctx.GetStub().PutState(referensiHarga.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	referensiHargaResponse := web.ReferensiHargaResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    referensiHarga.Id,
		IdDinas:               referensiHarga.IdDinas,
		UmurTanam:             referensiHarga.UmurTanam,
		Harga:                 referensiHarga.Harga,
	}

	return &referensiHargaResponse, nil
}

func (c *ReferensiHargaChaincode) Update(ctx contractapi.TransactionContextInterface, payload string) (*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to update asset, does not have dinas.user affiliation/role")
	}

	var referensiHargaUpdateRequest web.ReferensiHargaUpdateRequest
	err = json.Unmarshal([]byte(payload), &referensiHargaUpdateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, referensiHargaUpdateRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", referensiHargaUpdateRequest.Id)
	}

	referensiHarga := domain.ReferensiHarga{
		Id:        referensiHargaUpdateRequest.Id,
		IdDinas:   referensiHargaUpdateRequest.IdDinas,
		UmurTanam: referensiHargaUpdateRequest.UmurTanam,
		Harga:     referensiHargaUpdateRequest.Harga,
	}

	referensiHargaJSON, err := json.Marshal(referensiHarga)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated asset: %v", err)
	}

	err = ctx.GetStub().PutState(referensiHarga.Id, referensiHargaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	referensiHargaResponse := web.ReferensiHargaResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    referensiHarga.Id,
		IdDinas:               referensiHarga.IdDinas,
		UmurTanam:             referensiHarga.UmurTanam,
		Harga:                 referensiHarga.Harga,
	}

	return &referensiHargaResponse, nil
}

func (c *ReferensiHargaChaincode) GetAll(ctx contractapi.TransactionContextInterface) ([]*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var referensiHargaResponses []*web.ReferensiHargaResponse
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		txID := ctx.GetStub().GetTxID()

		var referensiHarga domain.ReferensiHarga
		err = json.Unmarshal(queryResponse.Value, &referensiHarga)
		if err != nil {
			return nil, err
		}

		referensiHargaResponse := web.ReferensiHargaResponse{
			IdTransaksiBlockchain: txID,
			Id:                    referensiHarga.Id,
			IdDinas:               referensiHarga.IdDinas,
			UmurTanam:             referensiHarga.UmurTanam,
			Harga:                 referensiHarga.Harga,
		}
		referensiHargaResponses = append(referensiHargaResponses, &referensiHargaResponse)
	}

	return referensiHargaResponses, nil
}

func (c *ReferensiHargaChaincode) GetHistoryById(ctx contractapi.TransactionContextInterface, payload string) ([]*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	var referensiHargaGetRequest web.ReferensiHargaGetRequest
	err = json.Unmarshal([]byte(payload), &referensiHargaGetRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(referensiHargaGetRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for key %s: %v", referensiHargaGetRequest.Id, err)
	}
	defer resultsIterator.Close()

	var referensiHargaResponses []*web.ReferensiHargaResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", referensiHargaGetRequest.Id, err)
		}

		var referensiHarga domain.ReferensiHarga
		err = json.Unmarshal(response.Value, &referensiHarga)
		if err != nil {
			return nil, err
		}

		referensiHargaResponse := web.ReferensiHargaResponse{
			IdTransaksiBlockchain: response.GetTxId(),
			Id:                    referensiHarga.Id,
			IdDinas:               referensiHarga.IdDinas,
			UmurTanam:             referensiHarga.UmurTanam,
			Harga:                 referensiHarga.Harga,
		}
		referensiHargaResponses = append(referensiHargaResponses, &referensiHargaResponse)
	}

	return referensiHargaResponses, nil
}
