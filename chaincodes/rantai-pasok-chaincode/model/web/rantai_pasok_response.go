package web

import (
	"rantai-pasok-chaincode/model/domain"
)

type CreateKontrakResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain"`
	Id                    string  `json:"id"`
	Nomor                 string  `json:"nomor"`
	TanggalPembuatan      string  `json:"tanggalPembuatan"`
	TangalMulai           string  `json:"tanggalMulai"`
	TanggalSelesai        string  `json:"tanggalSelesai"`
	IdPks                 string  `json:"idPks"`
	IdKoperasi            string  `json:"idKoperasi"`
	Kuantitas             float32 `json:"kuantitas"`
	Harga                 float64 `json:"harga"`
	Status                string  `json:"status"`
}

type ConfirmKontrakResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain"`
	Id                    string  `json:"id"`
	Nomor                 string  `json:"nomor"`
	TanggalPembuatan      string  `json:"tanggalPembuatan"`
	TangalMulai           string  `json:"tanggalMulai"`
	TanggalSelesai        string  `json:"tanggalSelesai"`
	IdPks                 string  `json:"idPks"`
	IdKoperasi            string  `json:"idKoperasi"`
	Kuantitas             float32 `json:"kuantitas"`
	Harga                 float64 `json:"harga"`
	Status                string  `json:"status"`
	Pesan                 string  `json:"pesan"`
	TanggalRespons        string  `json:"tanggalRespons"`
	KuantitasTepenuhi     float32 `json:"kuantitasTepenuhi"`
	KuantitasTersisa      float32 `json:"kuantitasTersisa"`
}

type KontrakResponse struct {
	IdTransaksiBlockchain string                 `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string                 `json:"id"`
	Nomor                 string                 `json:"nomor"`
	TanggalPembuatan      string                 `json:"tanggalPembuatan"`
	TangalMulai           string                 `json:"tanggalMulai"`
	TanggalSelesai        string                 `json:"tanggalSelesai"`
	IdPks                 string                 `json:"idPks"`
	IdKoperasi            string                 `json:"idKoperasi"`
	Kuantitas             float32                `json:"kuantitas"`
	Harga                 float64                `json:"harga"`
	Status                string                 `json:"status"`
	Pesan                 string                 `json:"pesan"`
	TanggalRespons        string                 `json:"tanggalRespons"`
	DeliveryOrders        []domain.DeliveryOrder `json:"deliveryOrders"`
	KuantitasTepenuhi     float32                `json:"kuantitasTepenuhi"`
	KuantitasTersisa      float32                `json:"kuantitasTersisa"`
}
