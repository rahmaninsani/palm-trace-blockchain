package web

type PembayaranResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdTransaksi           string  `json:"idTransaksi"`
	JenisUser             string  `json:"jenisUser"`
	Nomor                 string  `json:"nomor"`
	Tanggal               string  `json:"tanggal"`
	JumlahBayar           float64 `json:"jumlahBayar"`
	HashBukti             string  `json:"hashBukti"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}
