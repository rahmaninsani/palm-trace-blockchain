package helper

import (
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func ToKontrakResponse(ctx contractapi.TransactionContextInterface, kontrak domain.Kontrak) *web.KontrakResponse {
	kontrakResponse := &web.KontrakResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    kontrak.Id,
		IdPks:                 kontrak.IdPks,
		IdKoperasi:            kontrak.IdKoperasi,
		Nomor:                 kontrak.Nomor,
		TanggalPembuatan:      kontrak.TanggalPembuatan,
		TangalMulai:           kontrak.TangalMulai,
		TanggalSelesai:        kontrak.TanggalSelesai,
		Kuantitas:             kontrak.Kuantitas,
		Harga:                 kontrak.Harga,
		Status:                kontrak.Status.String(),
		Pesan:                 kontrak.Pesan,
		TanggalRespons:        kontrak.TanggalRespons,
		KuantitasTepenuhi:     kontrak.KuantitasTepenuhi,
		KuantitasTersisa:      kontrak.KuantitasTersisa,
		CreatedAt:             kontrak.CreatedAt,
		UpdatedAt:             kontrak.UpdatedAt,
	}

	return kontrakResponse
}

func ToKebunResponse(ctx contractapi.TransactionContextInterface, kebun domain.Kebun) *web.KebunResponse {
	kebunResponse := &web.KebunResponse{
		IdTransaksiBlockchain: ctx.GetStub().GetTxID(),
		Id:                    kebun.Id,
		IdPetani:              kebun.IdPetani,
		Alamat:                kebun.Alamat,
		Latitude:              kebun.Latitude,
		Longitude:             kebun.Longitude,
		Luas:                  kebun.Luas,
		NomorRspo:             kebun.NomorRspo,
		SertifikatRspo:        kebun.SertifikatRspo,
		CreatedAt:             kebun.CreatedAt,
		UpdatedAt:             kebun.UpdatedAt,
	}

	return kebunResponse
}

// func ToKebunResponses(response *queryresult.KeyModification) *web.KebunResponse {
// 	var kebunResponse web.KebunResponse
// 	if err := json.Unmarshal(response.Value, &kebunResponse); err != nil {
// 		return &web.KebunResponse{}
// 	}

// 	if _, found := reflect.TypeOf(response).MethodByName("GetTxId"); found {
// 		kebunResponse.IdTransaksiBlockchain = response.GetTxId()
// 	}

// 	return &kebunResponse
// }
