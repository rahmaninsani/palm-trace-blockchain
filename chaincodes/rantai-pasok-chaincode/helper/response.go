package helper

import (
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

func ToKontrakResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kontrak domain.Kontrak) *web.KontrakResponse {
	kontrakResponse := &web.KontrakResponse{
		Id:                kontrak.Id,
		IdPks:             kontrak.IdPks,
		IdKoperasi:        kontrak.IdKoperasi,
		Nomor:             kontrak.Nomor,
		TanggalPembuatan:  kontrak.TanggalPembuatan,
		TangalMulai:       kontrak.TangalMulai,
		TanggalSelesai:    kontrak.TanggalSelesai,
		Kuantitas:         kontrak.Kuantitas,
		Harga:             kontrak.Harga,
		Status:            kontrak.Status.String(),
		Pesan:             kontrak.Pesan,
		TanggalRespons:    kontrak.TanggalRespons,
		KuantitasTepenuhi: kontrak.KuantitasTepenuhi,
		KuantitasTersisa:  kontrak.KuantitasTersisa,
		CreatedAt:         kontrak.CreatedAt,
		UpdatedAt:         kontrak.UpdatedAt,
	}

	if ctx != nil {
		kontrakResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kontrakResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kontrakResponse
}

func ToKebunResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kebun domain.Kebun) *web.KebunResponse {
	kebunResponse := &web.KebunResponse{
		Id:             kebun.Id,
		IdPetani:       kebun.IdPetani,
		Alamat:         kebun.Alamat,
		Latitude:       kebun.Latitude,
		Longitude:      kebun.Longitude,
		Luas:           kebun.Luas,
		NomorRspo:      kebun.NomorRspo,
		SertifikatRspo: kebun.SertifikatRspo,
		CreatedAt:      kebun.CreatedAt,
		UpdatedAt:      kebun.UpdatedAt,
	}

	if ctx != nil {
		kebunResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kebunResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kebunResponse
}
