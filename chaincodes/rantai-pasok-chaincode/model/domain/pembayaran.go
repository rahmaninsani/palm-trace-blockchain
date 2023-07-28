package domain

import "rantai-pasok-chaincode/constant"

type Pembayaran struct {
	Id          string             `json:"id"`
	AssetType   constant.AssetType `json:"assetType"`
	IdTransaksi string             `json:"idTransaksi"`
	Nomor       string             `json:"nomor"`
	Tanggal     string             `json:"tanggal"`
	Jumlah      float64            `json:"jumlah"`
	HashBukti   string             `json:"hashBukti"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
}
