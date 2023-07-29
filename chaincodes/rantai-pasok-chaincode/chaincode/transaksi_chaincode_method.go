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

func (c *RantaiPasokChaincodeImpl) TransaksiCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiCreateRequest web.TransaksiCreateRequest
	if err := json.Unmarshal([]byte(payload), &transaksiCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksi := domain.Transaksi{
		Id:                     transaksiCreateRequest.Id,
		AssetType:              constant.AssetTypeTransaksi,
		IdDeliveryOrder:        transaksiCreateRequest.IdDeliveryOrder,
		IdPetani:               transaksiCreateRequest.IdPetani,
		Nomor:                  transaksiCreateRequest.Nomor,
		TanggalPembuatan:       transaksiCreateRequest.TanggalPembuatan,
		StatusKoperasi:         constant.PenawaranTransaksiMenungguKonfirmasi,
		PesanKoperasi:          "",
		TanggalResponsKoperasi: "",
		StatusPks:              constant.PenawaranTransaksiMenungguKonfirmasi,
		PesanPks:               "",
		TanggalResponsPks:      "",
		Status:                 constant.TransaksiMenungguKonfirmasiKoperasi,
		CreatedAt:              transaksiCreateRequest.CreatedAt,
		UpdatedAt:              transaksiCreateRequest.UpdatedAt,
	}

	transaksiJSON, err := json.Marshal(transaksi)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(transaksi.Id, transaksiJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiResponse := helper.ToTransaksiResponse(ctx, nil, &transaksi)

	return helper.ToWebResponse(http.StatusCreated, transaksiResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) TransaksiConfirm(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiConfirmRequest web.TransaksiConfirmRequest
	if err := json.Unmarshal([]byte(payload), &transaksiConfirmRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiPrevBytes, err := ctx.GetStub().GetState(transaksiConfirmRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if transaksiPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var transaksi domain.Transaksi
	if err = json.Unmarshal(transaksiPrevBytes, &transaksi); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if transaksiConfirmRequest.StatusKoperasi != -1 {
		transaksi.StatusKoperasi = transaksiConfirmRequest.StatusKoperasi
		transaksi.PesanKoperasi = transaksiConfirmRequest.PesanKoperasi
		transaksi.TanggalResponsKoperasi = transaksiConfirmRequest.TanggalResponsKoperasi
		if transaksiConfirmRequest.StatusKoperasi == constant.PenawaranTransaksiDisetujui {
			transaksi.Status = constant.TransaksiMenungguKonfirmasiPks
		} else {
			transaksi.Status = constant.TransaksiDitolakKoperasi
		}
	}

	if transaksiConfirmRequest.StatusPks != -1 {
		transaksi.StatusPks = transaksiConfirmRequest.StatusPks
		transaksi.PesanPks = transaksiConfirmRequest.PesanPks
		transaksi.TanggalResponsPks = transaksiConfirmRequest.TanggalResponsPks
		if transaksiConfirmRequest.StatusPks == constant.PenawaranTransaksiDisetujui {
			transaksi.Status = constant.TransaksiMenungguDikirimPetani
		} else {
			transaksi.Status = constant.TransaksiDitolakPks
		}
	}

	transaksiJSON, err := json.Marshal(transaksi)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(transaksi.Id, transaksiJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiResponse := helper.ToTransaksiResponse(ctx, nil, &transaksi)
	return helper.ToWebResponse(http.StatusOK, transaksiResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) TransaksiUpdateStatus(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiUpdateStatusRequest web.TransaksiUpdateStatusRequest
	if err := json.Unmarshal([]byte(payload), &transaksiUpdateStatusRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiPrevBytes, err := ctx.GetStub().GetState(transaksiUpdateStatusRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if transaksiPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var transaksi domain.Transaksi
	if err = json.Unmarshal(transaksiPrevBytes, &transaksi); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksi.Status = transaksiUpdateStatusRequest.Status
	transaksi.UpdatedAt = transaksiUpdateStatusRequest.UpdatedAt

	transaksiJSON, err := json.Marshal(transaksi)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(transaksi.Id, transaksiJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kontrakResponse := helper.ToTransaksiResponse(ctx, nil, &transaksi)

	return helper.ToWebResponse(http.StatusOK, kontrakResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) TransaksiFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var transaksiFindAllRequest web.TransaksiFindAllRequest
	if err := json.Unmarshal([]byte(payload), &transaksiFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType":       constant.AssetTypeTransaksi,
			"idDeliveryOrder": transaksiFindAllRequest.IdDeliveryOrder,
		},
	}

	if transaksiFindAllRequest.IdPetani != "" {
		query["selector"].(map[string]interface{})["idPetani"] = transaksiFindAllRequest.IdPetani
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

	var transaksiResponses []*web.TransaksiResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var transaksi domain.Transaksi
		if err = json.Unmarshal(response.Value, &transaksi); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		transaksiResponses = append(transaksiResponses, helper.ToTransaksiResponse(nil, nil, &transaksi))
	}

	return helper.ToWebResponse(http.StatusOK, transaksiResponses, nil)
}

func (c *RantaiPasokChaincodeImpl) TransaksiFindOne(ctx contractapi.TransactionContextInterface, idTransaksi string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	transaksiPrevBytes, err := ctx.GetStub().GetState(idTransaksi)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if transaksiPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var transaksi domain.Transaksi
	err = json.Unmarshal(transaksiPrevBytes, &transaksi)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	transaksiResponse := helper.ToTransaksiResponse(nil, nil, &transaksi)

	return helper.ToWebResponse(http.StatusOK, transaksiResponse, nil)
}
