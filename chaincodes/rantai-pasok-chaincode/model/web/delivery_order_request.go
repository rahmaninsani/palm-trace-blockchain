package web

import "rantai-pasok-chaincode/constant"

type DeliveryOrderCreateRequest struct {
	Id               string  `json:"id"`
	IdKontrak        string  `json:"idKontrak"`
	Nomor            string  `json:"nomor"`
	TanggalPembuatan string  `json:"tanggalPembuatan"`
	Periode          string  `json:"periode"`
	Kuantitas        float32 `json:"kuantitas"`
	Harga            float64 `json:"harga"`
	Rendemen         float32 `json:"rendemen"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}

type DeliveryOrderConfirmRequest struct {
	Id             string                                `json:"id"`
	Status         constant.StatusPenawaranDeliveryOrder `json:"status"`
	Pesan          string                                `json:"pesan"`
	TanggalRespons string                                `json:"tanggalRespons"`
	UpdatedAt      string                                `json:"updatedAt"`
}

type DeliveryOrderUpdateKuantitasRequest struct {
	Id                 string  `json:"id"`
	KuantitasTerpenuhi float32 `json:"kuantitasTerpenuhi"`
	UpdatedAt          string  `json:"updatedAt"`
}

type DeliveryOrderFindAllRequest struct {
	IdKontrak string                                `json:"idKontrak"`
	Status    constant.StatusPenawaranDeliveryOrder `json:"status"`
}
