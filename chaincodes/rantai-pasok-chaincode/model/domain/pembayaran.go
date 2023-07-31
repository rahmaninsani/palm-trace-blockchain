package domain

import "rantai-pasok-chaincode/constant"

type Pembayaran struct {
	Id          string             `json:"id"`
	AssetType   constant.AssetType `json:"assetType"`
	IdTransaksi string             `json:"idTransaksi"`
	JenisUser   string             `json:"jenisUser"`
	Nomor       string             `json:"nomor"`
	Tanggal     string             `json:"tanggal"`
	JumlahBayar float64            `json:"jumlahBayar"`
	HashBukti   string             `json:"hashBukti"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
}
