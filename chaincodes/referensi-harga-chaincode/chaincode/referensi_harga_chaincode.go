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

	var referensiHarga domain.ReferensiHarga
	err = json.Unmarshal([]byte(payload), &referensiHarga)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, referensiHarga.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", referensiHarga.Id)
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
		TanggalPembaruan:      referensiHarga.TanggalPembaruan,
	}

	return &referensiHargaResponse, nil
}

func (c *ReferensiHargaChaincode) Update(ctx contractapi.TransactionContextInterface, payload string) (*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to update asset, does not have dinas.user affiliation/role")
	}

	var referensiHarga domain.ReferensiHarga
	err = json.Unmarshal([]byte(payload), &referensiHarga)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := helper.AssetExists(ctx, referensiHarga.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", referensiHarga.Id)
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
		TanggalPembaruan:      referensiHarga.TanggalPembaruan,
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
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var referensiHarga domain.ReferensiHarga
		err = json.Unmarshal(response.Value, &referensiHarga)
		if err != nil {
			return nil, err
		}

		referensiHargaResponse := web.ReferensiHargaResponse{
			Id:               referensiHarga.Id,
			IdDinas:          referensiHarga.IdDinas,
			UmurTanam:        referensiHarga.UmurTanam,
			Harga:            referensiHarga.Harga,
			TanggalPembaruan: referensiHarga.TanggalPembaruan,
		}
		referensiHargaResponses = append(referensiHargaResponses, &referensiHargaResponse)
	}

	return referensiHargaResponses, nil
}

func (c *ReferensiHargaChaincode) GetHistoryById(ctx contractapi.TransactionContextInterface, id string) ([]*web.ReferensiHargaResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for key %s: %v", id, err)
	}
	defer resultsIterator.Close()

	var referensiHargaResponses []*web.ReferensiHargaResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", id, err)
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
			TanggalPembaruan:      referensiHarga.TanggalPembaruan,
		}
		referensiHargaResponses = append(referensiHargaResponses, &referensiHargaResponse)
	}

	return referensiHargaResponses, nil
}
