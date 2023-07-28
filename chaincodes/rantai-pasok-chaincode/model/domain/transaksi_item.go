package domain

import "rantai-pasok-chaincode/constant"

type TransaksiItem struct {
	Id          string             `json:"id"`
	AssetType   constant.AssetType `json:"assetType"`
	IdTransaksi string             `json:"idTransaksi"`
	IdKebun     string             `json:"idKebun"`
	// TODO IdTransaksiBlockchainKebun string  `json:"idTransaksiBlockchainKebun"`
	Kuantitas float32 `json:"kuantitas"`
	Harga     float64 `json:"harga"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
