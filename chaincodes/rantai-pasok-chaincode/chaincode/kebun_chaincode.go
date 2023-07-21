package chaincode

import (
	"encoding/json"
	"fmt"
	"rantai-pasok-chaincode/helper"
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateKebun digunakan untuk membuat kebun baru
func (c *RantaiPasokChaincode) CreateKebun(ctx contractapi.TransactionContextInterface, payload string) (*web.KebunResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"petani.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have petani.user affiliation/role")
	}

	var kebun domain.Kebun
	err = json.Unmarshal([]byte(payload), &kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	// Buat kunci komposit dengan format "Kebun_idPetani"
	kebunCompositeKey, err := ctx.GetStub().CreateCompositeKey("Kebun", []string{kebun.IdPetani, kebun.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	exists, err := helper.AssetExists(ctx, kebunCompositeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", kebunCompositeKey)
	}

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kebun: %v", err)
	}

	err = ctx.GetStub().PutState(kebunCompositeKey, kebunJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put kebun on ledger: %v", err)
	}

	kebunResponse := &web.KebunResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    kebun.Id,
		IdPetani:              kebun.IdPetani,
		Alamat:                kebun.Alamat,
		Latitude:              kebun.Latitude,
		Longitude:             kebun.Longitude,
		Luas:                  kebun.Luas,
		NomorRspo:             kebun.NomorRspo,
		SertifikatRspo:        kebun.SertifikatRspo,
		UpdatedAt:             kebun.UpdatedAt,
	}

	return kebunResponse, nil
}

// UpdateKebun digunakan untuk memperbarui data kebun
func (c *RantaiPasokChaincode) UpdateKebun(ctx contractapi.TransactionContextInterface, payload string) (*web.KebunResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"petani.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to update asset, does not have petani.user affiliation/role")
	}

	var kebun domain.Kebun
	err = json.Unmarshal([]byte(payload), &kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	// Buat kunci komposit dengan format "Kebun_idPetani"
	kebunCompositeKey, err := ctx.GetStub().CreateCompositeKey("Kebun", []string{kebun.IdPetani, kebun.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	exists, err := helper.AssetExists(ctx, kebunCompositeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", kebunCompositeKey)
	}

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated asset: %v", err)
	}

	err = ctx.GetStub().PutState(kebunCompositeKey, kebunJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put state: %v", err)
	}

	kebunResponse := &web.KebunResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    kebun.Id,
		IdPetani:              kebun.IdPetani,
		Alamat:                kebun.Alamat,
		Latitude:              kebun.Latitude,
		Longitude:             kebun.Longitude,
		Luas:                  kebun.Luas,
		NomorRspo:             kebun.NomorRspo,
		SertifikatRspo:        kebun.SertifikatRspo,
		UpdatedAt:             kebun.UpdatedAt,
	}

	return kebunResponse, nil
}

// GetAllKebunByIdPetani digunakan untuk mendapatkan data kebun berdasarkan ID Petani
func (c *RantaiPasokChaincode) GetAllKebunByIdPetani(ctx contractapi.TransactionContextInterface, idPetani string) ([]*web.KebunResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to query asset, does not have petani.user affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("Kebun", []string{idPetani})
	if err != nil {
		return nil, fmt.Errorf("failed to get kebun for petani: %v", err)
	}
	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %v", err)
		}

		var kebunResponse web.KebunResponse
		err = json.Unmarshal(response.Value, &kebunResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal kebun response: %v", err)
		}

		kebunResponses = append(kebunResponses, &kebunResponse)
	}

	return kebunResponses, nil
}

// GetKebunById digunakan untuk mendapatkan history perubahan data kebun berdasarkan ID
func (c *RantaiPasokChaincode) GetKebunHistoryById(ctx contractapi.TransactionContextInterface, payload string) ([]*web.KebunResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to query asset, does not have petani.user affiliation/role")
	}

	var kebunHistoryByIdRequest web.KebunHistoryByIdRequest
	err = json.Unmarshal([]byte(payload), &kebunHistoryByIdRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	kebunCompositeKey, err := ctx.GetStub().CreateCompositeKey("Kebun", []string{kebunHistoryByIdRequest.IdPetani, kebunHistoryByIdRequest.IdKebun})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(kebunCompositeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for kebun with ID %s: %v", kebunCompositeKey, err)
	}
	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", kebunHistoryByIdRequest.IdKebun, err)
		}

		var kebunResponse web.KebunResponse
		err = json.Unmarshal(response.Value, &kebunResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal kebun response: %v", err)
		}

		kebunResponse.IdTransaksiBlockchain = response.GetTxId()
		kebunResponses = append(kebunResponses, &kebunResponse)
	}

	return kebunResponses, nil
}
