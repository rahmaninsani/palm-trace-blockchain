package web

type PenerimaanCreateRequest struct {
	Id          string  `json:"id"`
	IdTransaksi string  `json:"idTransaksi"`
	Nomor       string  `json:"nomor"`
	Tanggal     string  `json:"tanggal"`
	Kuantitas   float32 `json:"kuantitas"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type PenerimaanFindAllRequest struct {
	IdTransaksi string `json:"idTransaksi"`
}
