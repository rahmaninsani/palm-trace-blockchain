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
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var deliveryOrderCreateRequest web.DeliveryOrderCreateRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderCreateRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryOrder := domain.DeliveryOrder{
		Id:               deliveryOrderCreateRequest.Id,
		AssetType:        constant.AssetTypeDeliveryOrder,
		IdKontrak:        deliveryOrderCreateRequest.IdKontrak,
		Nomor:            deliveryOrderCreateRequest.Nomor,
		TanggalPembuatan: deliveryOrderCreateRequest.TanggalPembuatan,
		Periode:          deliveryOrderCreateRequest.Periode,
		Kuantitas:        deliveryOrderCreateRequest.Kuantitas,
		Harga:            deliveryOrderCreateRequest.Harga,
		Rendemen:         deliveryOrderCreateRequest.Rendemen,
		Status:           constant.PenawaranDeliveryOrderMenungguKonfirmasi,
		Pesan:            "",
		TanggalRespons:   "",
		CreatedAt:        deliveryOrderCreateRequest.CreatedAt,
		UpdatedAt:        deliveryOrderCreateRequest.UpdatedAt,
	}

	deliveryOrderJSON, err := json.Marshal(deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(deliveryOrder.Id, deliveryOrderJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryResponse := helper.ToDeliveryOrderResponse(ctx, nil, deliveryOrder)

	return helper.ToWebResponse(http.StatusCreated, "Created", deliveryResponse)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderConfirm(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var deliveryOrderConfirmRequest web.DeliveryOrderConfirmRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderConfirmRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryOrderPrevBytes, err := ctx.GetStub().GetState(deliveryOrderConfirmRequest.Id)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if deliveryOrderPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var deliveryOrder domain.DeliveryOrder
	if err = json.Unmarshal(deliveryOrderPrevBytes, &deliveryOrder); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryOrder.Status = deliveryOrderConfirmRequest.Status
	deliveryOrder.Pesan = deliveryOrderConfirmRequest.Pesan
	deliveryOrder.TanggalRespons = deliveryOrderConfirmRequest.TanggalRespons
	deliveryOrder.UpdatedAt = deliveryOrderConfirmRequest.UpdatedAt

	deliveryOrderJSON, err := json.Marshal(deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if err = ctx.GetStub().PutState(deliveryOrder.Id, deliveryOrderJSON); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryOrderResponse := helper.ToDeliveryOrderResponse(ctx, nil, deliveryOrder)

	return helper.ToWebResponse(http.StatusOK, "OK", deliveryOrderResponse)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindAll(ctx contractapi.TransactionContextInterface, payload string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	var deliveryOrderFindAllRequest web.DeliveryOrderFindAllRequest
	if err := json.Unmarshal([]byte(payload), &deliveryOrderFindAllRequest); err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
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

	var deliveryOrderResponses []*web.DeliveryOrderResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var deliveryOrder domain.DeliveryOrder
		if err = json.Unmarshal(response.Value, &deliveryOrder); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		deliveryOrderResponses = append(deliveryOrderResponses, helper.ToDeliveryOrderResponse(nil, nil, deliveryOrder))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", deliveryOrderResponses)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindOne(ctx contractapi.TransactionContextInterface, idDeliveryOrder string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	deliveryOrderPrevBytes, err := ctx.GetStub().GetState(idDeliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if deliveryOrderPrevBytes == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	var deliveryOrder domain.DeliveryOrder
	err = json.Unmarshal(deliveryOrderPrevBytes, &deliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	deliveryResponse := helper.ToDeliveryOrderResponse(nil, nil, deliveryOrder)

	return helper.ToWebResponse(http.StatusOK, "OK", deliveryResponse)
}

func (c *RantaiPasokChaincodeImpl) DeliveryOrderFindOneHistory(ctx contractapi.TransactionContextInterface, idDeliveryOrder string) *web.WebResponse {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user", "koperasi.user", "petani.user"}); err != nil {
		return helper.ToWebResponse(http.StatusUnauthorized, err.Error(), nil)
	}

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(idDeliveryOrder)
	if err != nil {
		return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
	}

	if resultsIterator == nil {
		return helper.ToWebResponse(http.StatusNotFound, err.Error(), nil)
	}

	defer resultsIterator.Close()

	var deliveryOrderResponses []*web.DeliveryOrderResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		var deliveryOrder domain.DeliveryOrder
		if err = json.Unmarshal(response.Value, &deliveryOrder); err != nil {
			return helper.ToWebResponse(http.StatusInternalServerError, err.Error(), nil)
		}

		deliveryOrderResponses = append(deliveryOrderResponses, helper.ToDeliveryOrderResponse(nil, response, deliveryOrder))
	}

	return helper.ToWebResponse(http.StatusOK, "OK", deliveryOrderResponses)
}
