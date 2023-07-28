package domain

import "rantai-pasok-chaincode/constant"

type Penerimaan struct {
	Id          string             `json:"id"`
	AssetType   constant.AssetType `json:"assetType"`
	IdTransaksi string             `json:"idTransaksi"`
	Nomor       string             `json:"nomor"`
	Tanggal     string             `json:"tanggal"`
	Kuantitas   float32            `json:"kuantitas"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
}
