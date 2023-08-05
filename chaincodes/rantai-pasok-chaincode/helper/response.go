package helper

import (
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

func ToWebResponse(status int, data interface{}, err error) *web.WebResponse {
	webResponse := &web.WebResponse{
		Status: status,
	}

	if data != nil {
		webResponse.Data = data
	}

	if err != nil {
		webResponse.Message = err.Error()
	}

	return webResponse
}

func ToKebunResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kebun *domain.Kebun) *web.KebunResponse {
	kebunResponse := &web.KebunResponse{
		Id:                            kebun.Id,
		IdPetani:                      kebun.IdPetani,
		Alamat:                        kebun.Alamat,
		Latitude:                      kebun.Latitude,
		Longitude:                     kebun.Longitude,
		Luas:                          kebun.Luas,
		KemampuanProduksiHarian:       kebun.KemampuanProduksiHarian,
		NomorSuratKeteranganLurah:     kebun.NomorSuratKeteranganLurah,
		CidSuratKeteranganLurah:       kebun.CidSuratKeteranganLurah,
		NomorSuratKeteranganGantiRugi: kebun.NomorSuratKeteranganGantiRugi,
		CidSuratKeteranganGantiRugi:   kebun.CidSuratKeteranganGantiRugi,
		NomorSertifikatHakMilik:       kebun.NomorSertifikatHakMilik,
		CidSertifikatHakMilik:         kebun.CidSertifikatHakMilik,
		NomorSuratTandaBudidaya:       kebun.NomorSuratTandaBudidaya,
		CidSuratTandaBudidaya:         kebun.CidSuratTandaBudidaya,
		NomorSertifikatRspo:           kebun.NomorSertifikatRspo,
		CidSertifikatRspo:             kebun.CidSertifikatRspo,
		NomorSertifikatIspo:           kebun.NomorSertifikatIspo,
		CidSertifikatIspo:             kebun.CidSertifikatIspo,
		NomorSertifikatIscc:           kebun.NomorSertifikatIscc,
		CidSertifikatIscc:             kebun.CidSertifikatIscc,
		CreatedAt:                     kebun.CreatedAt,
		UpdatedAt:                     kebun.UpdatedAt,
	}

	if ctx != nil {
		kebunResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kebunResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kebunResponse
}

func ToKontrakResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kontrak *domain.Kontrak) *web.KontrakResponse {
	kontrakResponse := &web.KontrakResponse{
		Id:                 kontrak.Id,
		IdPks:              kontrak.IdPks,
		IdKoperasi:         kontrak.IdKoperasi,
		Nomor:              kontrak.Nomor,
		TanggalPembuatan:   kontrak.TanggalPembuatan,
		TangalMulai:        kontrak.TangalMulai,
		TanggalSelesai:     kontrak.TanggalSelesai,
		Kuantitas:          kontrak.Kuantitas,
		Harga:              kontrak.Harga,
		Status:             kontrak.Status.String(),
		Pesan:              kontrak.Pesan,
		TanggalKonfirmasi:  kontrak.TanggalKonfirmasi,
		KuantitasTerpenuhi: kontrak.KuantitasTerpenuhi,
		KuantitasTersisa:   kontrak.KuantitasTersisa,
		CreatedAt:          kontrak.CreatedAt,
		UpdatedAt:          kontrak.UpdatedAt,
	}

	if ctx != nil {
		kontrakResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kontrakResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kontrakResponse
}

func ToDeliveryOrderResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, deliveryOrder *domain.DeliveryOrder) *web.DeliveryOrderResponse {
	deliveryOrderResponse := &web.DeliveryOrderResponse{
		Id:                 deliveryOrder.Id,
		IdKontrak:          deliveryOrder.IdKontrak,
		Nomor:              deliveryOrder.Nomor,
		TanggalPembuatan:   deliveryOrder.TanggalPembuatan,
		Periode:            deliveryOrder.Periode,
		Kuantitas:          deliveryOrder.Kuantitas,
		Harga:              deliveryOrder.Harga,
		Rendemen:           deliveryOrder.Rendemen,
		Status:             deliveryOrder.Status.String(),
		Pesan:              deliveryOrder.Pesan,
		TanggalKonfirmasi:  deliveryOrder.TanggalKonfirmasi,
		KuantitasTerpenuhi: deliveryOrder.KuantitasTerpenuhi,
		KuantitasTersisa:   deliveryOrder.KuantitasTersisa,
		CreatedAt:          deliveryOrder.CreatedAt,
		UpdatedAt:          deliveryOrder.UpdatedAt,
	}

	if ctx != nil {
		deliveryOrderResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		deliveryOrderResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return deliveryOrderResponse
}

func ToTransaksiResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, transaksi *domain.Transaksi) *web.TransaksiResponse {
	transaksiResponse := &web.TransaksiResponse{
		Id:                        transaksi.Id,
		IdDeliveryOrder:           transaksi.IdDeliveryOrder,
		IdPetani:                  transaksi.IdPetani,
		Nomor:                     transaksi.Nomor,
		TanggalPembuatan:          transaksi.TanggalPembuatan,
		StatusKoperasi:            transaksi.StatusKoperasi.String(),
		PesanKoperasi:             transaksi.PesanKoperasi,
		TanggalKonfirmasiKoperasi: transaksi.TanggalKonfirmasiKoperasi,
		StatusPks:                 transaksi.StatusPks.String(),
		PesanPks:                  transaksi.PesanPks,
		TanggalKonfirmasiPks:      transaksi.TanggalKonfirmasiPks,
		Status:                    transaksi.Status.String(),
		CreatedAt:                 transaksi.CreatedAt,
		UpdatedAt:                 transaksi.UpdatedAt,
	}

	if ctx != nil {
		transaksiResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		transaksiResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return transaksiResponse
}

func ToTransaksiItemResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, transaksiItem *domain.TransaksiItem) *web.TransaksiItemResponse {
	transaksiItemResponse := &web.TransaksiItemResponse{
		Id:          transaksiItem.Id,
		IdTransaksi: transaksiItem.IdTransaksi,
		IdKebun:     transaksiItem.IdKebun,
		Kuantitas:   transaksiItem.Kuantitas,
		Harga:       transaksiItem.Harga,
		UmurTanam:   transaksiItem.UmurTanam,
		CreatedAt:   transaksiItem.CreatedAt,
		UpdatedAt:   transaksiItem.UpdatedAt,
	}

	if ctx != nil {
		transaksiItemResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		transaksiItemResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return transaksiItemResponse
}

func ToPengirimanResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, pengiriman *domain.Pengiriman) *web.PengirimanResponse {
	pengirimanResponse := &web.PengirimanResponse{
		Id:                   pengiriman.Id,
		IdTransaksi:          pengiriman.IdTransaksi,
		JenisUser:            pengiriman.JenisUser,
		Nomor:                pengiriman.Nomor,
		Tanggal:              pengiriman.Tanggal,
		NamaSopir:            pengiriman.NamaSopir,
		NomorTeleponSopir:    pengiriman.NomorTeleponSopir,
		NamaKendaraan:        pengiriman.NamaKendaraan,
		NomorPolisiKendaraan: pengiriman.NomorPolisiKendaraan,
		CreatedAt:            pengiriman.CreatedAt,
		UpdatedAt:            pengiriman.UpdatedAt,
	}

	if ctx != nil {
		pengirimanResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pengirimanResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pengirimanResponse
}

func ToPenerimaanResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, penerimaan *domain.Penerimaan) *web.PenerimaanResponse {
	pengirimanResponse := &web.PenerimaanResponse{
		Id:          penerimaan.Id,
		IdTransaksi: penerimaan.IdTransaksi,
		JenisUser:   penerimaan.JenisUser,
		Nomor:       penerimaan.Nomor,
		Tanggal:     penerimaan.Tanggal,
		Kuantitas:   penerimaan.Kuantitas,
		CreatedAt:   penerimaan.CreatedAt,
		UpdatedAt:   penerimaan.UpdatedAt,
	}

	if ctx != nil {
		pengirimanResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pengirimanResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pengirimanResponse
}

func ToPembayaranResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, pembayaran *domain.Pembayaran) *web.PembayaranResponse {
	pembayaranResponse := &web.PembayaranResponse{
		Id:                    pembayaran.Id,
		IdTransaksi:           pembayaran.IdTransaksi,
		JenisUser:             pembayaran.JenisUser,
		Nomor:                 pembayaran.Nomor,
		Tanggal:               pembayaran.Tanggal,
		NomorRekeningPengirim: pembayaran.NomorRekeningPengirim,
		NomorRekeningPenerima: pembayaran.NomorRekeningPenerima,
		CidBuktiPembayaran:    pembayaran.CidBuktiPembayaran,
		CreatedAt:             pembayaran.CreatedAt,
		UpdatedAt:             pembayaran.UpdatedAt,
	}

	if ctx != nil {
		pembayaranResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pembayaranResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pembayaranResponse
}
