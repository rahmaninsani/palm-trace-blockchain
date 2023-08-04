package web

type TransaksiItemResponse struct {
	IdTransaksiBlockchain string `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string `json:"id"`
	IdTransaksi           string `json:"idTransaksi"`
	IdKebun               string `json:"idKebun"`
	// TODO IdTransaksiBlockchainKebun string  `json:"idTransaksiBlockchainKebun"`
	Kuantitas float32 `json:"kuantitas"`
	Harga     float64 `json:"harga"`
	UmurTanam int     `json:"umurTanam"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
