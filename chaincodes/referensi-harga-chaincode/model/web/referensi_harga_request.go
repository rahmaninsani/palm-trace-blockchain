package web

type ReferensiHargaCreateRequest struct {
	Id               string  `json:"id"`
	IdDinas          string  `json:"idDinas"`
	UmurTanam        int     `json:"umurTanam"`
	Harga            float64 `json:"harga"`
	TanggalPembaruan string  `json:"tanggalPembaruan"`
}

type ReferensiHargaUpdateRequest struct {
	Id               string  `json:"id"`
	IdDinas          string  `json:"idDinas"`
	UmurTanam        int     `json:"umurTanam"`
	Harga            float64 `json:"harga"`
	TanggalPembaruan string  `json:"tanggalPembaruan"`
}

type ReferensiHargaGetRequest struct {
	Id string `json:"id"`
}
