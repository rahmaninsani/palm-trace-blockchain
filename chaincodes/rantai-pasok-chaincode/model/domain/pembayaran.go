package domain

type Pembayaran struct {
	Id          string  `json:"id"`
	IdTransaksi string  `json:"idTransaksi"`
	Nomor       string  `json:"nomor"`
	Tanggal     string  `json:"tanggal"`
	Jumlah      float64 `json:"jumlah"`
	Bukti       string  `json:"bukti"`
}
