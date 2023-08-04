package web

type TransaksiItemCreateRequest struct {
	Id          string  `json:"id"`
	IdTransaksi string  `json:"idTransaksi"`
	IdKebun     string  `json:"idKebun"`
	Kuantitas   float32 `json:"kuantitas"`
	Harga       float64 `json:"harga"`
	UmurTanam   int     `json:"umurTanam"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type TransaksiItemFindAllRequest struct {
	IdTransaksi string `json:"idTransaksi"`
}
