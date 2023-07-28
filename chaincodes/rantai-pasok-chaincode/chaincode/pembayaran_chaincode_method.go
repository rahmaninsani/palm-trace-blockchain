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

func (c *RantaiPasokChaincodeImpl) PembayaranCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var pembayaranCreateRequest web.PembayaranCreateRequest
	if err := json.Unmarshal([]byte(payload), &pembayaranCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	pembayaran := domain.Pembayaran{
		Id:          pembayaranCreateRequest.Id,
		AssetType:   constant.AssetTypePenerimaan,
		IdTransaksi: pembayaranCreateRequest.IdTransaksi,
		Nomor:       pembayaranCreateRequest.Nomor,
		Tanggal:     pembayaranCreateRequest.Tanggal,
		Jumlah:      pembayaranCreateRequest.Jumlah,
		HashBukti:   pembayaranCreateRequest.HashBukti,
		CreatedAt:   pembayaranCreateRequest.CreatedAt,
		UpdatedAt:   pembayaranCreateRequest.UpdatedAt,
	}

	pembayaranJSON, err := json.Marshal(pembayaran)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(pembayaran.Id, pembayaranJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	pembayaranResponse := helper.ToPembayaranResponse(ctx, nil, &pembayaran)

	return helper.ToWebResponse(http.StatusCreated, pembayaranResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) PembayaranFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var pembayaranFindAllRequest web.PembayaranFindAllRequest
	if err := json.Unmarshal([]byte(payload), &pembayaranFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType":   constant.AssetTypePembayaran,
			"idTransaksi": pembayaranFindAllRequest.IdTransaksi,
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

	var pembayaranResponses []*web.PembayaranResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var pembayaran domain.Pembayaran
		if err = json.Unmarshal(response.Value, &pembayaran); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		pembayaranResponses = append(pembayaranResponses, helper.ToPembayaranResponse(nil, nil, &pembayaran))
	}

	return helper.ToWebResponse(http.StatusOK, pembayaranResponses, nil)
}
