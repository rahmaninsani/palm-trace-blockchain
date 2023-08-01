package web

type TransaksiResponse struct {
	IdTransaksiBlockchain string `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string `json:"id"`
	IdDeliveryOrder       string `json:"idDeliveryOrder"`
	IdPetani              string `json:"idPetani"`
	Nomor                 string `json:"nomor"`
	TanggalPembuatan      string `json:"tanggalPembuatan"`
	// TransaksiItems         []*TransaksiItemResponse `json:"transaksiItem,omitempty" metadata:",optional"`
	StatusKoperasi            string `json:"statusKoperasi"`
	PesanKoperasi             string `json:"pesanKoperasi"`
	TanggalKonfirmasiKoperasi string `json:"tanggalKonfirmasiKoperasi"`
	StatusPks                 string `json:"statusPks"`
	PesanPks                  string `json:"pesanPks"`
	TanggalKonfirmasiPks      string `json:"tanggalKonfirmasiPks"`
	Status                    string `json:"status"`
	CreatedAt                 string `json:"createdAt"`
	UpdatedAt                 string `json:"updatedAt"`
}
