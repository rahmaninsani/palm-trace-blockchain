package web

type KebunCreateRequest struct {
	Id             string  `json:"id"`
	IdPetani       string  `json:"idPetani"`
	Alamat         string  `json:"alamat"`
	Latitude       string  `json:"latitude"`
	Longitude      string  `json:"longitude"`
	Luas           float64 `json:"luas"`
	NomorRspo      string  `json:"nomorRspo"`
	SertifikatRspo string  `json:"sertifikatRspo"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

type KebunUpdateRequest struct {
	Id             string  `json:"id"`
	IdPetani       string  `json:"idPetani"`
	Alamat         string  `json:"alamat"`
	Latitude       string  `json:"latitude"`
	Longitude      string  `json:"longitude"`
	Luas           float64 `json:"luas"`
	NomorRspo      string  `json:"nomorRspo"`
	SertifikatRspo string  `json:"sertifikatRspo"`
	UpdatedAt      string  `json:"updatedAt"`
}

type KebunHistoryRequest struct {
	IdPetani string `json:"idPetani"`
	IdKebun  string `json:"idKebun"`
}
