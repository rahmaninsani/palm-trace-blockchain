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

func (c *RantaiPasokChaincodeImpl) TransaksiItemCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiItemCreateRequest web.TransaksiItemCreateRequest
	if err := json.Unmarshal([]byte(payload), &transaksiItemCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiItem := domain.TransaksiItem{
		Id:          transaksiItemCreateRequest.Id,
		AssetType:   constant.AssetTypeTransaksiItem,
		IdTransaksi: transaksiItemCreateRequest.IdTransaksi,
		IdKebun:     transaksiItemCreateRequest.IdKebun,
		// TODO IdTransaksiBlockchainKebun: "",
		Kuantitas: transaksiItemCreateRequest.Kuantitas,
		Harga:     transaksiItemCreateRequest.Harga,
		CreatedAt: transaksiItemCreateRequest.CreatedAt,
		UpdatedAt: transaksiItemCreateRequest.UpdatedAt,
	}

	transaksiItemJSON, err := json.Marshal(transaksiItem)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(transaksiItem.Id, transaksiItemJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiItemResponse := helper.ToTransaksiItemResponse(ctx, nil, &transaksiItem)

	return helper.ToWebResponse(http.StatusCreated, transaksiItemResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) TransaksiItemFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiItemFindAllRequest web.TransaksiItemFindAllRequest
	if err := json.Unmarshal([]byte(payload), &transaksiItemFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType":   constant.AssetTypeTransaksiItem,
			"idTransaksi": transaksiItemFindAllRequest.IdTransaksi,
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

	var transaksiItemResponses []*web.TransaksiItemResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var transaksiItem domain.TransaksiItem
		if err = json.Unmarshal(response.Value, &transaksiItem); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		transaksiItemResponses = append(transaksiItemResponses, helper.ToTransaksiItemResponse(nil, nil, &transaksiItem))
	}

	return helper.ToWebResponse(http.StatusOK, transaksiItemResponses, nil)
}
