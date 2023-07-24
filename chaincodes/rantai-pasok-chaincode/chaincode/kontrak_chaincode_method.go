package chaincode

import (
	"encoding/json"
	"fmt"
	"rantai-pasok-chaincode/constant"
	"rantai-pasok-chaincode/helper"
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *RantaiPasokChaincodeImpl) KontrakCreate(ctx contractapi.TransactionContextInterface, payload string) (*web.KontrakResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	var kontrakCreateRequest web.KontrakCreateRequest
	if err := json.Unmarshal([]byte(payload), &kontrakCreateRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kontrakPrev, err := helper.GetAsset(ctx, kontrakCreateRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if kontrakPrev != nil {
		return nil, fmt.Errorf("the asset %s already exists", kontrakCreateRequest.Id)
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
		return nil, fmt.Errorf("failed to marshal kontrak: %v", err)
	}

	if err = ctx.GetStub().PutState(kontrak.Id, kontrakJSON); err != nil {
		return nil, fmt.Errorf("failed to put kontrak on ledger: %v", err)
	}

	return helper.ToKontrakResponse(ctx, nil, kontrak), nil
}

func (c *RantaiPasokChaincodeImpl) KontrakConfirm(ctx contractapi.TransactionContextInterface, payload string) (*web.KontrakResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	var kontrakConfirmRequest web.KontrakConfirmRequest
	if err := json.Unmarshal([]byte(payload), &kontrakConfirmRequest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kontrakPrev, err := helper.GetAsset(ctx, kontrakConfirmRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if kontrakPrev == nil {
		return nil, fmt.Errorf("the asset %s does not exist", kontrakConfirmRequest.Id)
	}

	var kontrak domain.Kontrak
	if err = json.Unmarshal(kontrakPrev, &kontrak); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kontrak: %v", err)
	}

	if kontrak.IdKoperasi != kontrakConfirmRequest.IdKoperasi {
		return nil, fmt.Errorf("the asset %s is not assigned to the koperasi %s", kontrak.Id, kontrakConfirmRequest.IdKoperasi)
	}

	if kontrak.Status != constant.PenawaranKontrakMenungguKonfirmasi {
		return nil, fmt.Errorf("the asset %s has been confirmed by Koperasi", kontrak.IdKoperasi)
	}

	// Konfirmasi kontrak
	kontrak.Status = kontrakConfirmRequest.Status
	kontrak.Pesan = kontrakConfirmRequest.Pesan
	kontrak.TanggalRespons = kontrakConfirmRequest.TanggalRespons
	kontrak.KuantitasTersisa = kontrak.Kuantitas
	kontrak.UpdatedAt = kontrakConfirmRequest.UpdatedAt

	kontrakJSON, err := json.Marshal(kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kontrak: %v", err)
	}

	if err = ctx.GetStub().PutState(kontrak.Id, kontrakJSON); err != nil {
		return nil, fmt.Errorf("failed to put kontrak on ledger: %v", err)
	}

	return helper.ToKontrakResponse(ctx, nil, kontrak), nil
}

func (c *RantaiPasokChaincodeImpl) KontrakGetAllByIdPks(ctx contractapi.TransactionContextInterface, idPks string) ([]*web.KontrakResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	queryString := fmt.Sprintf(`{
		"selector": {
			"assetType": %d,
			"idPks": "%s"
		}
	}`, constant.AssetTypeKontrak, idPks)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get kontrak for pks: %v", err)
	}

	if resultsIterator == nil {
		return nil, fmt.Errorf("kontrak for pks with ID %s does not exist", idPks)
	}

	defer resultsIterator.Close()

	var kontrakResponses []*web.KontrakResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %v", err)
		}

		var kontrak domain.Kontrak
		if err = json.Unmarshal(response.Value, &kontrak); err != nil {
			return nil, fmt.Errorf("failed to unmarshal kontrak response: %v", err)
		}

		kontrakResponses = append(kontrakResponses, helper.ToKontrakResponse(nil, nil, kontrak))
	}

	return kontrakResponses, nil
}

func (c *RantaiPasokChaincodeImpl) KontrakGetAllByIdKoperasi(ctx contractapi.TransactionContextInterface, idKoperasi string) ([]*web.KontrakResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"koperasi.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	queryString := fmt.Sprintf(`{
		"selector": {
			"assetType": %d,
			"idKoperasi": "%s"
		}
	}`, constant.AssetTypeKontrak, idKoperasi)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get kontrak for koperasi: %v", err)
	}

	if resultsIterator == nil {
		return nil, fmt.Errorf("kontrak for koperasi with ID %s does not exist", idKoperasi)
	}

	defer resultsIterator.Close()

	var kontrakResponses []*web.KontrakResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %v", err)
		}

		var kontrak domain.Kontrak
		if err = json.Unmarshal(response.Value, &kontrak); err != nil {
			return nil, fmt.Errorf("failed to unmarshal kontrak response: %v", err)
		}

		kontrakResponses = append(kontrakResponses, helper.ToKontrakResponse(nil, nil, kontrak))
	}

	return kontrakResponses, nil
}

func (c *RantaiPasokChaincodeImpl) KontrakGetAllForPetani(ctx contractapi.TransactionContextInterface, idKoperasi string) ([]*web.KontrakResponse, error) {
	if err := helper.CheckAffiliation(ctx, []string{"petani.user"}); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	queryString := fmt.Sprintf(`{
		"selector": {
			"assetType": %d,
			"idKoperasi": "%s",
			"status": %d
		}
	}`, constant.AssetTypeKontrak, idKoperasi, constant.PenawaranKontrakDisetujui)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get kontrak for koperasi: %v", err)
	}

	if resultsIterator == nil {
		return nil, fmt.Errorf("kontrak for koperasi with ID %s does not exist", idKoperasi)
	}

	defer resultsIterator.Close()

	var kontrakResponses []*web.KontrakResponse
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %v", err)
		}

		var kontrak domain.Kontrak
		if err = json.Unmarshal(response.Value, &kontrak); err != nil {
			return nil, fmt.Errorf("failed to unmarshal kontrak response: %v", err)
		}

		kontrakResponses = append(kontrakResponses, helper.ToKontrakResponse(nil, nil, kontrak))
	}

	return kontrakResponses, nil
}
