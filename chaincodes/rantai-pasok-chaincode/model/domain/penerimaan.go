package domain

type Penerimaan struct {
	Id                string  `json:"id"`
	IdTransaksi       string  `json:"idTransaksi"`
	Nomor             string  `json:"nomor"`
	Tanggal           string  `json:"tanggal"`
	KuantitasDiterima float32 `json:"kuantitasDiterima"`
}
