package chaincode

import (
	"encoding/json"
	"net/http"
	"rantai-pasok-chaincode/constant"
	"rantai-pasok-chaincode/helper"
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *RantaiPasokChaincodeImpl) KebunCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var kebunCreateRequest web.KebunCreateRequest
	if err := json.Unmarshal([]byte(payload), &kebunCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
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
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kebunResponse := helper.ToKebunResponse(ctx, nil, kebun)

	return helper.ToWebResponse(http.StatusCreated, "Created", kebunResponse)
}

func (c *RantaiPasokChaincodeImpl) KebunUpdate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var kebunUpdateRequest web.KebunUpdateRequest
	if err := json.Unmarshal([]byte(payload), &kebunUpdateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kebunPrevBytes, err := ctx.GetStub().GetState(kebunUpdateRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if kebunPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var kebun domain.Kebun
	if err = json.Unmarshal(kebunPrevBytes, &kebun); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
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
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kebunResponse := helper.ToKebunResponse(ctx, nil, kebun)

	return helper.ToWebResponse(http.StatusOK, "OK", kebunResponse)
}

func (c *RantaiPasokChaincodeImpl) KebunFindAll(ctx contractapi.TransactionContextInterface, idPetani string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType": constant.AssetTypeKebun,
			"idPetani":  idPetani,
		},
	}

	queryString, err := helper.BuildQueryString(query)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if resultsIterator == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, nil, kebun))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", kebunResponses)
}

func (c *RantaiPasokChaincodeImpl) KebunFindOne(ctx contractapi.TransactionContextInterface, idKebun string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	kebunPrevBytes, err := ctx.GetStub().GetState(idKebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if kebunPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var kebun domain.Kebun
	err = json.Unmarshal(kebunPrevBytes, &kebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kebunResponse := helper.ToKebunResponse(nil, nil, kebun)

	return helper.ToWebResponse(http.StatusOK, "OK", kebunResponse)
}

func (c *RantaiPasokChaincodeImpl) KebunFindOneHistory(ctx contractapi.TransactionContextInterface, idKebun string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(idKebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if resultsIterator == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	defer resultsIterator.Close()

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, response, kebun))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", kebunResponses)
}
