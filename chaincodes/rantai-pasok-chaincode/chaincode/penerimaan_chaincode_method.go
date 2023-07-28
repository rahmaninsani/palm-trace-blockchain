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

func (c *RantaiPasokChaincodeImpl) PenerimaanCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var penerimaanCreateRequest web.PenerimaanCreateRequest
	if err := json.Unmarshal([]byte(payload), &penerimaanCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	penerimaan := domain.Penerimaan{
		Id:          penerimaanCreateRequest.Id,
		AssetType:   constant.AssetTypePenerimaan,
		IdTransaksi: penerimaanCreateRequest.IdTransaksi,
		Nomor:       penerimaanCreateRequest.Nomor,
		Tanggal:     penerimaanCreateRequest.Tanggal,
		Kuantitas:   penerimaanCreateRequest.Kuantitas,
		CreatedAt:   penerimaanCreateRequest.CreatedAt,
		UpdatedAt:   penerimaanCreateRequest.UpdatedAt,
	}

	penerimaanJSON, err := json.Marshal(penerimaan)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(penerimaan.Id, penerimaanJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	penerimaanResponse := helper.ToPenerimaanResponse(ctx, nil, &penerimaan)

	return helper.ToWebResponse(http.StatusCreated, penerimaanResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) PenerimaanFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var penerimaanFindAllRequest web.PenerimaanFindAllRequest
	if err := json.Unmarshal([]byte(payload), &penerimaanFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType":   constant.AssetTypePenerimaan,
			"idTransaksi": penerimaanFindAllRequest.IdTransaksi,
		},
	}

	queryString, err := helper.BuildQueryString(query)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var penerimaanResponses []*web.PenerimaanResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var penerimaan domain.Penerimaan
		if err = json.Unmarshal(response.Value, &penerimaan); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		penerimaanResponses = append(penerimaanResponses, helper.ToPenerimaanResponse(nil, nil, &penerimaan))
	}

	return helper.ToWebResponse(http.StatusOK, penerimaanResponses, nil)
}
