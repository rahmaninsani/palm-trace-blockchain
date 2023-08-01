package domain

import "rantai-pasok-chaincode/constant"

type Transaksi struct {
	Id                        string                            `json:"id"`
	AssetType                 constant.AssetType                `json:"assetType"`
	IdDeliveryOrder           string                            `json:"idDeliveryOrder"`
	IdPetani                  string                            `json:"idPetani"`
	Nomor                     string                            `json:"nomor"`
	TanggalPembuatan          string                            `json:"tanggalPembuatan"`
	StatusKoperasi            constant.StatusPenawaranTransaksi `json:"statusKoperasi"`
	PesanKoperasi             string                            `json:"pesanKoperasi"`
	TanggalKonfirmasiKoperasi string                            `json:"tanggalKonfirmasiKoperasi"`
	StatusPks                 constant.StatusPenawaranTransaksi `json:"statusPks"`
	PesanPks                  string                            `json:"pesanPks"`
	TanggalKonfirmasiPks      string                            `json:"tanggalKonfirmasiPks"`
	Status                    constant.StatusTransaksi          `json:"status"`
	CreatedAt                 string                            `json:"createdAt"`
	UpdatedAt                 string                            `json:"updatedAt"`
}
