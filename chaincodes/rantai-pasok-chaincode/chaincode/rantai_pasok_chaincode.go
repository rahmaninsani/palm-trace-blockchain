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

type RantaiPasokChaincode struct {
	contractapi.Contract
}

func (c *RantaiPasokChaincode) CreateKontrak(ctx contractapi.TransactionContextInterface, payload string) (*web.CreateKontrakResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have pabrikkelapasawit.user affiliation/role")
	}

	var kontrak domain.Kontrak
	err = json.Unmarshal([]byte(payload), &kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	compositeKeyType := "kontrak~id"
	compositeKeyComponents := []string{kontrak.Id}
	kontrakCompositeKey, err := ctx.GetStub().CreateCompositeKey(compositeKeyType, compositeKeyComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	exists, err := helper.AssetExists(ctx, kontrakCompositeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("the asset %s already exists", kontrakCompositeKey)
	}

	// Set nilai default atau kosong untuk field yang belum terisi
	kontrak.Status = constant.PenawaranKontrakMenungguKonfirmasi
	kontrak.DeliveryOrders = []domain.DeliveryOrder{}

	kontrakJSON, err := json.Marshal(kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kontrak: %v", err)
	}

	err = ctx.GetStub().PutState(kontrakCompositeKey, kontrakJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put kontrak on ledger: %v", err)
	}

	kontrakResponse := &web.CreateKontrakResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    kontrak.Id,
		Nomor:                 kontrak.Nomor,
		TanggalPembuatan:      kontrak.TanggalPembuatan,
		TangalMulai:           kontrak.TangalMulai,
		TanggalSelesai:        kontrak.TanggalSelesai,
		IdPks:                 kontrak.IdPks,
		IdKoperasi:            kontrak.IdKoperasi,
		Kuantitas:             kontrak.Kuantitas,
		Harga:                 kontrak.Harga,
		Status:                kontrak.Status.String(),
	}

	return kontrakResponse, nil
}

// func (c *RantaiPasokChaincode) ConfirmContract(ctx contractapi.TransactionContextInterface, payload string) (*web.ConfirmKontrakResponse, error) {
// 	err := helper.CheckAffiliation(ctx, []string{"koperasi.user"})
// 	if err != nil {
// 		return nil, fmt.Errorf("submitting client not authorized to confirm contract, does not have koperasi.user affiliation/role")
// 	}

// 	var confirmContractRequest web.ConfirmContractRequest
// 	err = json.Unmarshal([]byte(payload), &confirmContractRequest)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
// 	}

// 	compositeKeyType := "Kontrak"
// 	compositeKeyComponents := []string{confirmContractRequest.IdPks, confirmContractRequest.IdKoperasi, confirmContractRequest.IdKontrak}
// 	kontrakCompositeKey, err := helper.CreateCompositeKeyWithUnderscore(ctx, compositeKeyType, compositeKeyComponents)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create composite key: %v", err)
// 	}

// 	kontrakJSON, err := ctx.GetStub().GetState(kontrakCompositeKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read kontrak from world state: %v", err)
// 	}

// 	if kontrakJSON == nil {
// 		return nil, fmt.Errorf("the asset %s does not exist", kontrakCompositeKey)
// 	}

// 	var kontrak domain.Kontrak
// 	err = json.Unmarshal(kontrakJSON, &kontrak)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal kontrak: %v", err)
// 	}

// 	// Cek apakah kontrak sudah dikonfirmasi oleh Koperasi sebelumnya
// 	if kontrak.Status != helper.MenungguKonfirmasi {
// 		return nil, fmt.Errorf("the asset %s has been confirmed by Koperasi", kontrakCompositeKey)
// 	}

// 	// Lakukan konfirmasi kontrak oleh Koperasi
// 	kontrak.Status = confirmContractRequest.Status
// 	kontrak.Pesan = confirmContractRequest.Pesan
// 	kontrak.TanggalRespons = confirmContractRequest.TanggalRespons
// 	kontrak.KuantitasTersisa = kontrak.Kuantitas

// 	kontrakJSON, err = json.Marshal(kontrak)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal kontrak: %v", err)
// 	}

// 	err = ctx.GetStub().PutState(kontrakCompositeKey, kontrakJSON)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to put kontrak on ledger: %v", err)
// 	}

// 	kontrakResponse := &web.ConfirmKontrakResponse{
// 		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
// 		Id:                    kontrak.Id,
// 		Nomor:                 kontrak.Nomor,
// 		TanggalPembuatan:      kontrak.TanggalPembuatan,
// 		TangalMulai:           kontrak.TangalMulai,
// 		TanggalSelesai:        kontrak.TanggalSelesai,
// 		IdPks:                 kontrak.IdPks,
// 		IdKoperasi:            kontrak.IdKoperasi,
// 		Kuantitas:             kontrak.Kuantitas,
// 		Harga:                 kontrak.Harga,
// 		Status:                kontrak.Status.String(),
// 		Pesan:                 kontrak.Pesan,
// 		TanggalRespons:        kontrak.TanggalRespons,
// 		KuantitasTepenuhi:     kontrak.KuantitasTepenuhi,
// 		KuantitasTersisa:      kontrak.KuantitasTersisa,
// 	}

// 	return kontrakResponse, nil
// }

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
