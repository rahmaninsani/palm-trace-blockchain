package web

import "rantai-pasok-chaincode/constant"

type TransaksiCreateRequest struct {
	Id               string `json:"id"`
	IdDeliveryOrder  string `json:"idDeliveryOrder"`
	IdPetani         string `json:"idPetani"`
	Nomor            string `json:"nomor"`
	TanggalPembuatan string `json:"tanggalPembuatan"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

type TransaksiUpdateStatusRequest struct {
	Id        string                   `json:"id"`
	Status    constant.StatusTransaksi `json:"status"`
	UpdatedAt string                   `json:"updatedAt"`
}

type TransaksiConfirmRequest struct {
	Id                     string                            `json:"id"`
	StatusKoperasi         constant.StatusPenawaranTransaksi `json:"statusKoperasi"`
	PesanKoperasi          string                            `json:"pesanKoperasi"`
	TanggalResponsKoperasi string                            `json:"tanggalResponsKoperasi"`
	StatusPks              constant.StatusPenawaranTransaksi `json:"statusPks"`
	PesanPks               string                            `json:"pesanPks"`
	TanggalResponsPks      string                            `json:"tanggalResponsPks"`
	UpdatedAt              string                            `json:"updatedAt"`
}

type TransaksiFindAllRequest struct {
	IdDeliveryOrder string `json:"idDeliveryOrder"`
	IdPetani        string `json:"idPetani"`
}
