package domain

import "rantai-pasok-chaincode/helper"

type Kontrak struct {
	Id                string          `json:"id"`
	Nomor             string          `json:"nomor"`
	TanggalPembuatan  string          `json:"tanggalPembuatan"`
	TangalMulai       string          `json:"tanggalMulai"`
	TanggalSelesai    string          `json:"tanggalSelesai"`
	IdPks             string          `json:"idPks"`
	IdKoperasi        string          `json:"idKoperasi"`
	Kuantitas         float32         `json:"kuantitas"`
	Harga             float64         `json:"harga"`
	Status            helper.Status   `json:"status"`
	Pesan             string          `json:"pesan"`
	TanggalRespons    string          `json:"tanggalRespons"`
	DeliveryOrders    []DeliveryOrder `json:"deliveryOrders"`
	KuantitasTepenuhi float32         `json:"kuantitasTepenuhi"`
	KuantitasTersisa  float32         `json:"kuantitasTersisa"`
}
