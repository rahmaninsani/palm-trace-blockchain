package web

type KebunResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdPetani              string  `json:"idPetani"`
	Alamat                string  `json:"alamat"`
	Latitude              string  `json:"latitude"`
	Longitude             string  `json:"longitude"`
	Luas                  float64 `json:"luas"`
	NomorRspo             string  `json:"nomorRspo"`
	SertifikatRspo        string  `json:"sertifikatRspo"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}
