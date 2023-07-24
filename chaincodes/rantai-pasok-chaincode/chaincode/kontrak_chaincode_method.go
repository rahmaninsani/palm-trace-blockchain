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

	return helper.ToKontrakResponse(ctx, kontrak), nil
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
		return nil, fmt.Errorf("failed to read kontrak from world state: %v", err)
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

	return helper.ToKontrakResponse(ctx, kontrak), nil
}

// func (c *RantaiPasokChaincode) GetAllKontrakByIdPks(ctx contractapi.TransactionContextInterface, idPks string) ([]*web.KontrakResponse, error) {
// 	err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"})
// 	if err != nil {
// 		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have pabrikkelapasawit.user affiliation/role")
// 	}

// 	compositeKeyType := "Kontrak"
// 	compositeKeyComponents := []string{idPks}
// 	kontrakCompositeKey, err := helper.CreateCompositeKeyWithUnderscore(ctx, compositeKeyType, compositeKeyComponents)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create composite key: %v", err)
// 	}

// 	// Dapatkan semua hasil yang cocok dengan kunci komposit "Kontrak:idPks"
// 	kontrakResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(kontrakCompositeKey, []string{})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get kontrak by idPks: %v", err)
// 	}
// 	defer kontrakResultsIterator.Close()

// 	// Buat slice untuk menyimpan hasil kontrak
// 	var kontrakResponses []*web.KontrakResponse

// 	// Iterasi semua hasil dan tambahkan ke slice
// 	for kontrakResultsIterator.HasNext() {
// 		response, err := kontrakResultsIterator.Next()
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to iterate kontrak results: %v", err)
// 		}

// 		// Konversi nilai hasil dari byte ke struct domain.Kontrak
// 		var kontrak domain.Kontrak
// 		err = json.Unmarshal(response.Value, &kontrak)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal kontrak: %v", err)
// 		}

// 		// Buat kontrak response untuk dijadikan hasil respons
// 		kontrakResponse := &web.KontrakResponse{
// 			Id:                kontrak.Id,
// 			Nomor:             kontrak.Nomor,
// 			TanggalPembuatan:  kontrak.TanggalPembuatan,
// 			TangalMulai:       kontrak.TangalMulai,
// 			TanggalSelesai:    kontrak.TanggalSelesai,
// 			IdPks:             kontrak.IdPks,
// 			IdKoperasi:        kontrak.IdKoperasi,
// 			Kuantitas:         kontrak.Kuantitas,
// 			Harga:             kontrak.Harga,
// 			Status:            kontrak.Status.String(),
// 			Pesan:             kontrak.Pesan,
// 			TanggalRespons:    kontrak.TanggalRespons,
// 			DeliveryOrders:    kontrak.DeliveryOrders,
// 			KuantitasTepenuhi: kontrak.KuantitasTepenuhi,
// 			KuantitasTersisa:  kontrak.KuantitasTersisa,
// 		}

// 		// Tambahkan kontrak response ke slice
// 		kontrakResponses = append(kontrakResponses, kontrakResponse)
// 	}

// 	return kontrakResponses, nil
// }

// func (c *RantaiPasokChaincode) GetAllKontrakByIdKoperasi(ctx contractapi.TransactionContextInterface, idKoperasi string) ([]*web.KontrakResponse, error) {
// 	err := helper.CheckAffiliation(ctx, []string{"koperasi.user", "petani.user"})
// 	if err != nil {
// 		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have pabrikkelapasawit.user affiliation/role")
// 	}

// 	fmt.Println("idKoperasi: ", idKoperasi)

// 	kontrakResultsIterator, _ := ctx.GetStub().GetStateByPartialCompositeKey("Kontrak", []string{"*", idKoperasi})
// 	fmt.Println("1. kontrakResultsIterator.HasNext(): ", kontrakResultsIterator.HasNext())

// 	kontrakResultsIterator, err = ctx.GetStub().GetStateByPartialCompositeKey("Kontrak", []string{"*", idKoperasi, "*"})
// 	fmt.Println("2. kontrakResultsIterator.HasNext(): ", kontrakResultsIterator.HasNext())

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get kontrak by idKoperasi: %v", err)
// 	}

// 	defer kontrakResultsIterator.Close()

// 	// Buat slice untuk menyimpan hasil kontrak
// 	var kontrakResponses []*web.KontrakResponse

// 	// Iterasi semua hasil dan tambahkan ke slice
// 	for kontrakResultsIterator.HasNext() {
// 		response, err := kontrakResultsIterator.Next()
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to iterate kontrak results: %v", err)
// 		}

// 		// Konversi nilai hasil dari byte ke struct domain.Kontrak
// 		var kontrak domain.Kontrak
// 		err = json.Unmarshal(response.Value, &kontrak)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal kontrak: %v", err)
// 		}

// 		// Cek apakah kontrak telah di-assign ke koperasi yang dimaksud
// 		// if kontrak.IdKoperasi != idKoperasi {
// 		// 	continue // Jika tidak di-assign, skip ke kontrak berikutnya
// 		// }

// 		// Buat kontrak response untuk dijadikan hasil respons
// 		kontrakResponse := &web.KontrakResponse{
// 			Id:                kontrak.Id,
// 			Nomor:             kontrak.Nomor,
// 			TanggalPembuatan:  kontrak.TanggalPembuatan,
// 			TangalMulai:       kontrak.TangalMulai,
// 			TanggalSelesai:    kontrak.TanggalSelesai,
// 			IdPks:             kontrak.IdPks,
// 			IdKoperasi:        kontrak.IdKoperasi,
// 			Kuantitas:         kontrak.Kuantitas,
// 			Harga:             kontrak.Harga,
// 			Status:            kontrak.Status.String(),
// 			Pesan:             kontrak.Pesan,
// 			TanggalRespons:    kontrak.TanggalRespons,
// 			DeliveryOrders:    kontrak.DeliveryOrders,
// 			KuantitasTepenuhi: kontrak.KuantitasTepenuhi,
// 			KuantitasTersisa:  kontrak.KuantitasTersisa,
// 		}

// 		// Tambahkan kontrak response ke slice
// 		kontrakResponses = append(kontrakResponses, kontrakResponse)
// 	}

// 	return kontrakResponses, nil
// }
