package web

type PenerimaanResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdTransaksi           string  `json:"idTransaksi"`
	JenisUser             string  `json:"jenisUser"`
	Nomor                 string  `json:"nomor"`
	Tanggal               string  `json:"tanggal"`
	Kuantitas             float32 `json:"kuantitas"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}
