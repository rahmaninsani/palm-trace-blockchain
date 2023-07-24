package chaincode

import (
	"encoding/json"
	"fmt"
	"rantai-pasok-chaincode/constant"
	"rantai-pasok-chaincode/helper"
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *RantaiPasokChaincodeImpl) KebunCreate(ctx contractapi.TransactionContextInterface, payload string) (*web.KebunResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	var kebunCreateRequest web.KebunCreateRequest
	if err := json.Unmarshal([]byte(payload), &kebunCreateRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kebunPrev, err := helper.GetAsset(ctx, kebunCreateRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if kebunPrev != nil {
		return nil, fmt.Errorf("the asset %s already exists", kebunCreateRequest.Id)
	}

	kebun := domain.Kebun{
		Id:             kebunCreateRequest.Id,
		AssetType:      constant.AssetTypeKebun,
		IdPetani:       kebunCreateRequest.IdPetani,
		Alamat:         kebunCreateRequest.Alamat,
		Latitude:       kebunCreateRequest.Latitude,
		Longitude:      kebunCreateRequest.Longitude,
		Luas:           kebunCreateRequest.Luas,
		NomorRspo:      kebunCreateRequest.NomorRspo,
		SertifikatRspo: kebunCreateRequest.SertifikatRspo,
		CreatedAt:      kebunCreateRequest.CreatedAt,
		UpdatedAt:      kebunCreateRequest.UpdatedAt,
	}

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kebun: %v", err)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return nil, fmt.Errorf("failed to put kebun on ledger: %v", err)
	}

	return helper.ToKebunResponse(ctx, nil, kebun), nil
}

func (c *RantaiPasokChaincodeImpl) KebunUpdate(ctx contractapi.TransactionContextInterface, payload string) (*web.KebunResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	var kebunUpdateRequest web.KebunUpdateRequest
	if err := json.Unmarshal([]byte(payload), &kebunUpdateRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kebunPrev, err := helper.GetAsset(ctx, kebunUpdateRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if kebunPrev == nil {
		return nil, fmt.Errorf("the asset %s does not exist", kebunUpdateRequest.Id)
	}

	var kebun domain.Kebun
	if err = json.Unmarshal(kebunPrev, &kebun); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kebun: %v", err)
	}

	if kebun.IdPetani != kebunUpdateRequest.IdPetani {
		return nil, fmt.Errorf("the asset %s is not assigned to the petani %s", kebun.Id, kebunUpdateRequest.IdPetani)
	}

	kebun.Alamat = kebunUpdateRequest.Alamat
	kebun.Latitude = kebunUpdateRequest.Latitude
	kebun.Longitude = kebunUpdateRequest.Longitude
	kebun.Luas = kebunUpdateRequest.Luas
	kebun.NomorRspo = kebunUpdateRequest.NomorRspo
	kebun.SertifikatRspo = kebunUpdateRequest.SertifikatRspo
	kebun.UpdatedAt = kebunUpdateRequest.UpdatedAt

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kebun: %v", err)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return nil, fmt.Errorf("failed to put kebun on ledger: %v", err)
	}

	return helper.ToKebunResponse(ctx, nil, kebun), nil
}

func (c *RantaiPasokChaincodeImpl) KebunGetAllByIdPetani(ctx contractapi.TransactionContextInterface, idPetani string) ([]*web.KebunResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	queryString := fmt.Sprintf(`{
		"selector": {
			"assetType": %d,
			"idPetani": "%s"
		}
	}`, constant.AssetTypeKebun, idPetani)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get kebun for petani: %v", err)
	}

	if resultsIterator == nil {
		return nil, fmt.Errorf("kebun for petani with ID %s does not exist", idPetani)
	}

	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %v", err)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			return nil, fmt.Errorf("failed to unmarshal kebun response: %v", err)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, nil, kebun))
	}

	return kebunResponses, nil
}

func (c *RantaiPasokChaincodeImpl) KebunGetHistoryById(ctx contractapi.TransactionContextInterface, payload string) ([]*web.KebunResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	var kebunHistoryRequest web.KebunHistoryRequest
	if err := json.Unmarshal([]byte(payload), &kebunHistoryRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	queryString := fmt.Sprintf(`{
		"selector": {
			"assetType": %d,
			"idPetani": "%s",
			"idKebun": "%s"
		}
	}`, constant.AssetTypeKebun, kebunHistoryRequest.IdPetani, kebunHistoryRequest.IdKebun)

	kebunPrev, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get kebun for petani: %v", err)
	}

	if kebunPrev == nil {
		return nil, fmt.Errorf("kebun with ID %s does not exist", kebunHistoryRequest.IdKebun)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(kebunHistoryRequest.IdKebun)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for kebun with ID %s: %v", kebunHistoryRequest.IdKebun, err)
	}

	if resultsIterator == nil {
		return nil, fmt.Errorf("kebun with ID %s does not exist", kebunHistoryRequest.IdKebun)
	}

	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", kebunHistoryRequest.IdKebun, err)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			return nil, fmt.Errorf("failed to unmarshal kebun response: %v", err)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, response, kebun))
	}

	return kebunResponses, nil
}
