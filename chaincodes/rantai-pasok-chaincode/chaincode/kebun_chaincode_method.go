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
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var kebunCreateRequest web.KebunCreateRequest
	if err := json.Unmarshal([]byte(payload), &kebunCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebun := domain.Kebun{
		Id:                            kebunCreateRequest.Id,
		AssetType:                     constant.AssetTypeKebun,
		IdPetani:                      kebunCreateRequest.IdPetani,
		Alamat:                        kebunCreateRequest.Alamat,
		Latitude:                      kebunCreateRequest.Latitude,
		Longitude:                     kebunCreateRequest.Longitude,
		Luas:                          kebunCreateRequest.Luas,
		KemampuanProduksiHarian:       kebunCreateRequest.KemampuanProduksiHarian,
		NomorSuratKeteranganLurah:     kebunCreateRequest.NomorSuratKeteranganLurah,
		CidSuratKeteranganLurah:       kebunCreateRequest.CidSuratKeteranganLurah,
		NomorSuratKeteranganGantiRugi: kebunCreateRequest.NomorSuratKeteranganGantiRugi,
		CidSuratKeteranganGantiRugi:   kebunCreateRequest.CidSuratKeteranganGantiRugi,
		NomorSertifikatHakMilik:       kebunCreateRequest.NomorSertifikatHakMilik,
		CidSertifikatHakMilik:         kebunCreateRequest.CidSertifikatHakMilik,
		NomorSuratTandaBudidaya:       kebunCreateRequest.NomorSuratTandaBudidaya,
		CidSuratTandaBudidaya:         kebunCreateRequest.CidSuratTandaBudidaya,
		NomorSertifikatRspo:           kebunCreateRequest.NomorSertifikatRspo,
		CidSertifikatRspo:             kebunCreateRequest.CidSertifikatRspo,
		NomorSertifikatIspo:           kebunCreateRequest.NomorSertifikatIspo,
		CidSertifikatIspo:             kebunCreateRequest.CidSertifikatIspo,
		NomorSertifikatIscc:           kebunCreateRequest.NomorSertifikatIscc,
		CidSertifikatIscc:             kebunCreateRequest.CidSertifikatIscc,
		CreatedAt:                     kebunCreateRequest.CreatedAt,
		UpdatedAt:                     kebunCreateRequest.UpdatedAt,
	}

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebunResponse := helper.ToKebunResponse(ctx, nil, &kebun)

	return helper.ToWebResponse(http.StatusCreated, kebunResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) KebunUpdate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var kebunUpdateRequest web.KebunUpdateRequest
	if err := json.Unmarshal([]byte(payload), &kebunUpdateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebunPrevBytes, err := ctx.GetStub().GetState(kebunUpdateRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if kebunPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var kebun domain.Kebun
	if err = json.Unmarshal(kebunPrevBytes, &kebun); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebun.Alamat = kebunUpdateRequest.Alamat
	kebun.Latitude = kebunUpdateRequest.Latitude
	kebun.Longitude = kebunUpdateRequest.Longitude
	kebun.Luas = kebunUpdateRequest.Luas
	kebun.KemampuanProduksiHarian = kebunUpdateRequest.KemampuanProduksiHarian
	kebun.NomorSuratKeteranganLurah = kebunUpdateRequest.NomorSuratKeteranganLurah
	kebun.CidSuratKeteranganLurah = kebunUpdateRequest.CidSuratKeteranganLurah
	kebun.NomorSuratKeteranganGantiRugi = kebunUpdateRequest.NomorSuratKeteranganGantiRugi
	kebun.CidSuratKeteranganGantiRugi = kebunUpdateRequest.CidSuratKeteranganGantiRugi
	kebun.NomorSertifikatHakMilik = kebunUpdateRequest.NomorSertifikatHakMilik
	kebun.CidSertifikatHakMilik = kebunUpdateRequest.CidSertifikatHakMilik
	kebun.NomorSuratTandaBudidaya = kebunUpdateRequest.NomorSuratTandaBudidaya
	kebun.CidSuratTandaBudidaya = kebunUpdateRequest.CidSuratTandaBudidaya
	kebun.NomorSertifikatRspo = kebunUpdateRequest.NomorSertifikatRspo
	kebun.CidSertifikatRspo = kebunUpdateRequest.CidSertifikatRspo
	kebun.NomorSertifikatIspo = kebunUpdateRequest.NomorSertifikatIspo
	kebun.CidSertifikatIspo = kebunUpdateRequest.CidSertifikatIspo
	kebun.NomorSertifikatIscc = kebunUpdateRequest.NomorSertifikatIscc
	kebun.CidSertifikatIscc = kebunUpdateRequest.CidSertifikatIscc
	kebun.UpdatedAt = kebunUpdateRequest.UpdatedAt

	kebunJSON, err := json.Marshal(kebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(kebun.Id, kebunJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebunResponse := helper.ToKebunResponse(ctx, nil, &kebun)

	return helper.ToWebResponse(http.StatusOK, kebunResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) KebunFindAll(ctx contractapi.TransactionContextInterface, idPetani string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType": constant.AssetTypeKebun,
			"idPetani":  idPetani,
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

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, nil, &kebun))
	}

	return helper.ToWebResponse(http.StatusOK, kebunResponses, nil)
}

func (c *RantaiPasokChaincodeImpl) KebunFindOne(ctx contractapi.TransactionContextInterface, idKebun string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	kebunPrevBytes, err := ctx.GetStub().GetState(idKebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err, nil)
	}

	if kebunPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var kebun domain.Kebun
	err = json.Unmarshal(kebunPrevBytes, &kebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kebunResponse := helper.ToKebunResponse(nil, nil, &kebun)

	return helper.ToWebResponse(http.StatusOK, kebunResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) KebunFindOneHistory(ctx contractapi.TransactionContextInterface, idKebun string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user", "koperasi.user", "pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(idKebun)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var kebunResponses []*web.KebunResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var kebun domain.Kebun
		if err = json.Unmarshal(response.Value, &kebun); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		kebunResponses = append(kebunResponses, helper.ToKebunResponse(nil, response, &kebun))
	}

	return helper.ToWebResponse(http.StatusOK, kebunResponses, nil)
}
