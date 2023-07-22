package chaincode

import (
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have petani.user affiliation/role")
	}

	var kontrak domain.Kontrak
	err = json.Unmarshal([]byte(payload), &kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kontrakCompositeKey, err := ctx.GetStub().CreateCompositeKey("Kontrak", []string{kontrak.IdPks, kontrak.IdKoperasi, kontrak.Id})
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
	kontrak.Status = helper.MenungguKonfirmasi
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

func (c *RantaiPasokChaincode) ConfirmContractByKoperasi(ctx contractapi.TransactionContextInterface, payload string) (*web.ConfirmKontrakResponse, error) {
	err := helper.CheckAffiliation(ctx, []string{"koperasi.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to confirm contract, does not have koperasi.user affiliation/role")
	}

	var confirmContractRequest web.ConfirmContractRequest
	err = json.Unmarshal([]byte(payload), &confirmContractRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	kontrakCompositeKey, err := ctx.GetStub().CreateCompositeKey("Kontrak", []string{confirmContractRequest.IdPks, confirmContractRequest.IdKoperasi, confirmContractRequest.IdKontrak})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	kontrakJSON, err := ctx.GetStub().GetState(kontrakCompositeKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read kontrak from world state: %v", err)
	}

	if kontrakJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", kontrakCompositeKey)
	}

	var kontrak domain.Kontrak
	err = json.Unmarshal(kontrakJSON, &kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal kontrak: %v", err)
	}

	// Cek apakah kontrak sudah dikonfirmasi oleh Koperasi sebelumnya
	if kontrak.Status != helper.MenungguKonfirmasi {
		return nil, fmt.Errorf("the asset %s has been confirmed by Koperasi", kontrakCompositeKey)
	}

	// Lakukan konfirmasi kontrak oleh Koperasi
	kontrak.Status = confirmContractRequest.Status
	kontrak.Pesan = confirmContractRequest.Pesan
	kontrak.TanggalRespons = confirmContractRequest.TanggalRespons
	kontrak.KuantitasTersisa = kontrak.Kuantitas

	kontrakJSON, err = json.Marshal(kontrak)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal kontrak: %v", err)
	}

	err = ctx.GetStub().PutState(kontrakCompositeKey, kontrakJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put kontrak on ledger: %v", err)
	}

	kontrakResponse := &web.ConfirmKontrakResponse{
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
		Pesan:                 kontrak.Pesan,
		TanggalRespons:        kontrak.TanggalRespons,
		KuantitasTepenuhi:     kontrak.KuantitasTepenuhi,
		KuantitasTersisa:      kontrak.KuantitasTersisa,
	}

	return kontrakResponse, nil
}
