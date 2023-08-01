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

func (c *RantaiPasokChaincodeImpl) DeliveryOrderCreate(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var deliveryOrderCreateRequest web.DeliveryOrderCreateRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrder := domain.DeliveryOrder{
		Id:                 deliveryOrderCreateRequest.Id,
		AssetType:          constant.AssetTypeDeliveryOrder,
		IdKontrak:          deliveryOrderCreateRequest.IdKontrak,
		Nomor:              deliveryOrderCreateRequest.Nomor,
		TanggalPembuatan:   deliveryOrderCreateRequest.TanggalPembuatan,
		Periode:            deliveryOrderCreateRequest.Periode,
		Kuantitas:          deliveryOrderCreateRequest.Kuantitas,
		Harga:              deliveryOrderCreateRequest.Harga,
		Rendemen:           deliveryOrderCreateRequest.Rendemen,
		Status:             constant.PenawaranDeliveryOrderMenungguKonfirmasi,
		Pesan:              "",
		TanggalKonfirmasi:  "",
		KuantitasTerpenuhi: 0,
		KuantitasTersisa:   deliveryOrderCreateRequest.Kuantitas,
		CreatedAt:          deliveryOrderCreateRequest.CreatedAt,
		UpdatedAt:          deliveryOrderCreateRequest.UpdatedAt,
	}

	deliveryOrderJSON, err := json.Marshal(deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(deliveryOrder.Id, deliveryOrderJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryResponse := helper.ToDeliveryOrderResponse(ctx, nil, &deliveryOrder)

	return helper.ToWebResponse(http.StatusCreated, deliveryResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderConfirm(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var deliveryOrderConfirmRequest web.DeliveryOrderConfirmRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderConfirmRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrderPrevBytes, err := ctx.GetStub().GetState(deliveryOrderConfirmRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if deliveryOrderPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var deliveryOrder domain.DeliveryOrder
	if err = json.Unmarshal(deliveryOrderPrevBytes, &deliveryOrder); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrder.Status = deliveryOrderConfirmRequest.Status
	deliveryOrder.Pesan = deliveryOrderConfirmRequest.Pesan
	deliveryOrder.TanggalKonfirmasi = deliveryOrderConfirmRequest.TanggalKonfirmasi
	deliveryOrder.UpdatedAt = deliveryOrderConfirmRequest.UpdatedAt

	deliveryOrderJSON, err := json.Marshal(deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(deliveryOrder.Id, deliveryOrderJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrderResponse := helper.ToDeliveryOrderResponse(ctx, nil, &deliveryOrder)

	return helper.ToWebResponse(http.StatusOK, deliveryOrderResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderUpdateKuantitas(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var deliveryOrderUpdateKuantitasRequest web.DeliveryOrderUpdateKuantitasRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderUpdateKuantitasRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrderPrevBytes, err := ctx.GetStub().GetState(deliveryOrderUpdateKuantitasRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if deliveryOrderPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var deliveryOrder domain.DeliveryOrder
	if err = json.Unmarshal(deliveryOrderPrevBytes, &deliveryOrder); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryOrder.KuantitasTerpenuhi = deliveryOrder.KuantitasTerpenuhi + deliveryOrderUpdateKuantitasRequest.KuantitasTerpenuhi
	deliveryOrder.KuantitasTersisa = deliveryOrder.KuantitasTersisa - deliveryOrderUpdateKuantitasRequest.KuantitasTerpenuhi
	deliveryOrder.UpdatedAt = deliveryOrderUpdateKuantitasRequest.UpdatedAt

	deliveryOrderJSON, err := json.Marshal(deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if err = ctx.GetStub().PutState(deliveryOrder.Id, deliveryOrderJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	kontrakResponse := helper.ToDeliveryOrderResponse(ctx, nil, &deliveryOrder)

	return helper.ToWebResponse(http.StatusOK, kontrakResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	var deliveryOrderFindAllRequest web.DeliveryOrderFindAllRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"assetType": constant.AssetTypeDeliveryOrder,
			"idKontrak": deliveryOrderFindAllRequest.IdKontrak,
		},
	}

	if deliveryOrderFindAllRequest.Status != -1 {
		query["selector"].(map[string]interface{})["status"] = deliveryOrderFindAllRequest.Status
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

	var deliveryOrderResponses []*web.DeliveryOrderResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var deliveryOrder domain.DeliveryOrder
		if err = json.Unmarshal(response.Value, &deliveryOrder); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		deliveryOrderResponses = append(deliveryOrderResponses, helper.ToDeliveryOrderResponse(nil, nil, &deliveryOrder))
	}

	return helper.ToWebResponse(http.StatusOK, deliveryOrderResponses, nil)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindOne(ctx contractapi.TransactionContextInterface, idDeliveryOrder string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	deliveryOrderPrevBytes, err := ctx.GetStub().GetState(idDeliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	if deliveryOrderPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var deliveryOrder domain.DeliveryOrder
	err = json.Unmarshal(deliveryOrderPrevBytes, &deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	deliveryResponse := helper.ToDeliveryOrderResponse(nil, nil, &deliveryOrder)

	return helper.ToWebResponse(http.StatusOK, deliveryResponse, nil)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindOneHistory(ctx contractapi.TransactionContextInterface, idDeliveryOrder string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, nil, err)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(idDeliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return helper.ToWebResponse(http.StatusNotFound, nil, nil)
	}

	var deliveryOrderResponses []*web.DeliveryOrderResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		var deliveryOrder domain.DeliveryOrder
		if err = json.Unmarshal(response.Value, &deliveryOrder); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, nil, err)
		}

		deliveryOrderResponses = append(deliveryOrderResponses, helper.ToDeliveryOrderResponse(nil, response, &deliveryOrder))
	}

	return helper.ToWebResponse(http.StatusOK, deliveryOrderResponses, nil)
}
