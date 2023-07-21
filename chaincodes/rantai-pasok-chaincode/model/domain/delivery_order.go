package domain

import "rantai-pasok-chaincode/helper"

type DeliveryOrder struct {
	Id               string        `json:"id"`
	IdKontrak        string        `json:"idKontrak"`
	Nomor            string        `json:"nomor"`
	TanggalPembuatan string        `json:"tanggalPembuatan"`
	Periode          string        `json:"periode"`
	Kuantitas        float32       `json:"kuantitas"`
	Harga            float64       `json:"harga"`
	Rendemen         float32       `json:"rendemen"`
	Status           helper.Status `json:"status"`
	Pesan            string        `json:"pesan"`
	TanggalRespons   string        `json:"tanggalRespons"`
	Transactions     []Transaksi   `json:"transactions"`
}
