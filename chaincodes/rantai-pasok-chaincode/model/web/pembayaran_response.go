package web

type PembayaranResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdTransaksi           string  `json:"idTransaksi"`
	Nomor                 string  `json:"nomor"`
	Tanggal               string  `json:"tanggal"`
	Jumlah                float64 `json:"jumlah"`
	HashBukti             string  `json:"hashBukti"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}
