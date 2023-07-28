package web

type TransaksiResponse struct {
	IdTransaksiBlockchain string `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string `json:"id"`
	IdDeliveryOrder       string `json:"idDeliveryOrder"`
	IdPetani              string `json:"idPetani"`
	Nomor                 string `json:"nomor"`
	TanggalPembuatan      string `json:"tanggalPembuatan"`
	// TransaksiItems         []*TransaksiItemResponse `json:"transaksiItem,omitempty" metadata:",optional"`
	StatusKoperasi         string `json:"statusKoperasi"`
	PesanKoperasi          string `json:"pesanKoperasi"`
	TanggalResponsKoperasi string `json:"tanggalResponsKoperasi"`
	StatusPks              string `json:"statusPks"`
	PesanPks               string `json:"pesanPks"`
	TanggalResponsPks      string `json:"tanggalResponsPks"`
	Status                 string `json:"status"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
}
