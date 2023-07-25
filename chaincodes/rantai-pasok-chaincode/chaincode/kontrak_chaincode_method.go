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

func (c *RantaiPasokChaincodeImpl) KontrakCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var kontrakCreateRequest web.KontrakCreateRequest
	if err := json.Unmarshal([]byte(payload), &kontrakCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrak := domain.Kontrak{
		Id:                kontrakCreateRequest.Id,
		AssetType:         constant.AssetTypeKontrak,
		IdPks:             kontrakCreateRequest.IdPks,
		IdKoperasi:        kontrakCreateRequest.IdKoperasi,
		Nomor:             kontrakCreateRequest.Nomor,
		TanggalPembuatan:  kontrakCreateRequest.TanggalPembuatan,
		TangalMulai:       kontrakCreateRequest.TangalMulai,
		TanggalSelesai:    kontrakCreateRequest.TanggalSelesai,
		Kuantitas:         kontrakCreateRequest.Kuantitas,
		Harga:             kontrakCreateRequest.Harga,
		Status:            constant.PenawaranKontrakMenungguKonfirmasi,
		Pesan:             "",
		TanggalRespons:    "",
		KuantitasTepenuhi: 0,
		KuantitasTersisa:  0,
		CreatedAt:         kontrakCreateRequest.CreatedAt,
		UpdatedAt:         kontrakCreateRequest.UpdatedAt,
	}

	kontrakJSON, err := json.Marshal(kontrak)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(kontrak.Id, kontrakJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrakResponse := helper.ToKontrakResponse(ctx, nil, kontrak)

	return helper.ToWebResponse(http.StatusCreated, "Created", kontrakResponse)
}

func (c *RantaiPasokChaincodeImpl) KontrakConfirm(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var kontrakConfirmRequest web.KontrakConfirmRequest
	if err := json.Unmarshal([]byte(payload), &kontrakConfirmRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrakPrevBytes, err := ctx.GetStub().GetState(kontrakConfirmRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if kontrakPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var kontrak domain.Kontrak
	if err = json.Unmarshal(kontrakPrevBytes, &kontrak); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrak.Status = kontrakConfirmRequest.Status
	kontrak.Pesan = kontrakConfirmRequest.Pesan
	kontrak.TanggalRespons = kontrakConfirmRequest.TanggalRespons
	kontrak.KuantitasTersisa = kontrak.Kuantitas
	kontrak.UpdatedAt = kontrakConfirmRequest.UpdatedAt

	kontrakJSON, err := json.Marshal(kontrak)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(kontrak.Id, kontrakJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrakResponse := helper.ToKontrakResponse(ctx, nil, kontrak)

	return helper.ToWebResponse(http.StatusOK, "OK", kontrakResponse)
}

func (c *RantaiPasokChaincodeImpl) KontrakFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var kontrakFindAllRequest web.KontrakFindAllRequest
	if err := json.Unmarshal([]byte(payload), &kontrakFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType": constant.AssetTypeKontrak,
		},
	}

	if kontrakFindAllRequest.IdPks != "" {
		query["selector"].(map[string]interface{})["idPks"] = kontrakFindAllRequest.IdPks
	}

	if kontrakFindAllRequest.IdKoperasi != "" {
		query["selector"].(map[string]interface{})["idKoperasi"] = kontrakFindAllRequest.IdKoperasi
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

	var kontrakResponses []*web.KontrakResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var kontrak domain.Kontrak
		if err = json.Unmarshal(response.Value, &kontrak); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		kontrakResponses = append(kontrakResponses, helper.ToKontrakResponse(nil, nil, kontrak))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", kontrakResponses)
}

func (c *RantaiPasokChaincodeImpl) KontrakFindOne(ctx contractapi.TransactionContextInterface, idKontrak string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	kontrakPrevBytes, err := ctx.GetStub().GetState(idKontrak)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if kontrakPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var kontrak domain.Kontrak
	err = json.Unmarshal(kontrakPrevBytes, &kontrak)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	kontrakResponse := helper.ToKontrakResponse(nil, nil, kontrak)

	return helper.ToWebResponse(http.StatusOK, "OK", kontrakResponse)
}

func (c *RantaiPasokChaincodeImpl) KontrakFindOneHistory(ctx contractapi.TransactionContextInterface, idKontrak string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(idKontrak)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if resultsIterator == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	defer resultsIterator.Close()

	var kontrakResponses []*web.KontrakResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var kontrak domain.Kontrak
		if err = json.Unmarshal(response.Value, &kontrak); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		kontrakResponses = append(kontrakResponses, helper.ToKontrakResponse(nil, response, kontrak))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", kontrakResponses)
}
