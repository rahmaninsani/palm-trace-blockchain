package domain

type TransaksiItem struct {
	Id          string  `json:"id"`
	IdTransaksi string  `json:"idTransaksi"`
	Kuantitas   float32 `json:"kuantitas"`
	Harga       float64 `json:"harga"`
	Kebun       Kebun   `json:"kebun"`
}
