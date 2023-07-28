package domain

import "rantai-pasok-chaincode/constant"

type Pengiriman struct {
	Id                   string             `json:"id"`
	AssetType            constant.AssetType `json:"assetType"`
	IdTransaksi          string             `json:"idTransaksi"`
	Nomor                string             `json:"nomor"`
	Tanggal              string             `json:"tanggal"`
	NamaSopir            string             `json:"namaSopir"`
	NomorTeleponSopir    string             `json:"nomorTeleponSopir"`
	NamaKendaraan        string             `json:"namaKendaraan"`
	NomorPolisiKendaraan string             `json:"nomorPolisiKendaraan"`
	CreatedAt            string             `json:"createdAt"`
	UpdatedAt            string             `json:"updatedAt"`
}
