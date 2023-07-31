package domain

import (
	"rantai-pasok-chaincode/constant"
)

type Kontrak struct {
	Id                 string                          `json:"id"`
	AssetType          constant.AssetType              `json:"assetType"`
	IdPks              string                          `json:"idPks"`
	IdKoperasi         string                          `json:"idKoperasi"`
	Nomor              string                          `json:"nomor"`
	TanggalPembuatan   string                          `json:"tanggalPembuatan"`
	TangalMulai        string                          `json:"tanggalMulai"`
	TanggalSelesai     string                          `json:"tanggalSelesai"`
	Kuantitas          float32                         `json:"kuantitas"`
	Harga              float64                         `json:"harga"`
	Status             constant.StatusPenawaranKontrak `json:"status"`
	Pesan              string                          `json:"pesan"`
	TanggalRespons     string                          `json:"tanggalRespons"`
	KuantitasTerpenuhi float32                         `json:"kuantitasTerpenuhi"`
	KuantitasTersisa   float32                         `json:"kuantitasTersisa"`
	CreatedAt          string                          `json:"createdAt"`
	UpdatedAt          string                          `json:"updatedAt"`
}
