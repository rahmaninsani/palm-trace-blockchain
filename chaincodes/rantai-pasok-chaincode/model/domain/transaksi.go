package domain

import "rantai-pasok-chaincode/helper"

type Transaksi struct {
	Id                     string                 `json:"id"`
	IdDeliveryOrder        string                 `json:"idDeliveryOrder"`
	Nomor                  string                 `json:"nomor"`
	Tanggal                string                 `json:"tanggalPembuatan"`
	IdPetani               string                 `json:"idPetani"`
	TransaksiItem          []TransaksiItem        `json:"detailTransaksi"`
	TotalKuantitas         float32                `json:"totalKuantitas"`
	TotalHarga             float64                `json:"totalHarga"`
	StatusKoperasi         helper.Status          `json:"statusKoperasi"`
	PesanKoperasi          string                 `json:"pesanKoperasi"`
	TanggalResponsKoperasi string                 `json:"tanggalResponsKoperasi"`
	StatusPks              helper.Status          `json:"statusPks"`
	PesanPks               string                 `json:"pesanPks"`
	TanggalResponsPks      string                 `json:"tanggalResponsPks"`
	PengirimanPetani       Pengiriman             `json:"pengirimanPetani"`
	PenerimaanKoperasi     Penerimaan             `json:"penerimaanKoperasi"`
	PengirimanKoperasi     Pengiriman             `json:"pengirimanKoperasi"`
	PenerimaanPks          Penerimaan             `json:"penerimaanPks"`
	PembayaranPks          Pembayaran             `json:"pembayaranPks"`
	PembayaranKoperasi     Pembayaran             `json:"pembayaranKoperasi"`
	Status                 helper.StatusTransaksi `json:"status"`
}
