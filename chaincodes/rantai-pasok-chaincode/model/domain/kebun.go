package domain

import "rantai-pasok-chaincode/constant"

type Kebun struct {
	Id             string             `json:"id"`
	AssetType      constant.AssetType `json:"assetType"`
	IdPetani       string             `json:"idPetani"`
	Alamat         string             `json:"alamat"`
	Latitude       string             `json:"latitude"`
	Longitude      string             `json:"longitude"`
	Luas           float64            `json:"luas"`
	NomorRspo      string             `json:"nomorRspo"`
	SertifikatRspo string             `json:"sertifikatRspo"`
	CreatedAt      string             `json:"createdAt"`
	UpdatedAt      string             `json:"updatedAt"`
}
