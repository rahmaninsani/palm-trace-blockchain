package domain

type ReferensiHarga struct {
	Id               string  `json:"id"`
	IdDinas          string  `json:"idDinas"`
	UmurTanam        int     `json:"umurTanam"`
	Harga            float64 `json:"harga"`
	TanggalPembaruan string  `json:"tanggalPembaruan"`
}
