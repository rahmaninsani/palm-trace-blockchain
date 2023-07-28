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

func (c *RantaiPasokChaincodeImpl) PengirimanCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var pengirimanCreateRequest web.PengirimanCreateRequest
	if err := json.Unmarshal([]byte(payload), &pengirimanCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	pengiriman := domain.Pengiriman{
		Id:                   pengirimanCreateRequest.Id,
		AssetType:            constant.AssetTypePengiriman,
		IdTransaksi:          pengirimanCreateRequest.IdTransaksi,
		Nomor:                pengirimanCreateRequest.Nomor,
		Tanggal:              pengirimanCreateRequest.Tanggal,
		NamaSopir:            pengirimanCreateRequest.NamaSopir,
		NomorTeleponSopir:    pengirimanCreateRequest.NomorTeleponSopir,
		NamaKendaraan:        pengirimanCreateRequest.NamaKendaraan,
		NomorPolisiKendaraan: pengirimanCreateRequest.NomorPolisiKendaraan,
		CreatedAt:            pengirimanCreateRequest.CreatedAt,
		UpdatedAt:            pengirimanCreateRequest.UpdatedAt,
	}

	pengirimanJSON, err := json.Marshal(pengiriman)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(pengiriman.Id, pengirimanJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	pengirimanResponse := helper.ToPengirimanResponse(ctx, nil, &pengiriman)

	return helper.ToWebResponse(http.StatusCreated, pengirimanResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) PengirimanFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var pengirimanFindAllRequest web.PengirimanFindAllRequest
	if err := json.Unmarshal([]byte(payload), &pengirimanFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType":   constant.AssetTypePengiriman,
			"idTransaksi": pengirimanFindAllRequest.IdTransaksi,
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

	var pengirimanResponses []*web.PengirimanResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var pengiriman domain.Pengiriman
		if err = json.Unmarshal(response.Value, &pengiriman); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		pengirimanResponses = append(pengirimanResponses, helper.ToPengirimanResponse(nil, nil, &pengiriman))
	}

	return helper.ToWebResponse(http.StatusOK, pengirimanResponses, nil)
}
