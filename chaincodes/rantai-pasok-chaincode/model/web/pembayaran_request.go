package web

type PembayaranCreateRequest struct {
	Id                    string  `json:"id"`
	IdTransaksi           string  `json:"idTransaksi"`
	JenisUser             string  `json:"jenisUser"`
	Nomor                 string  `json:"nomor"`
	Tanggal               string  `json:"tanggal"`
	JumlahPembayaran      float64 `json:"jumlahPembayaran"`
	NamaBankPengirim      string  `json:"namaBankPengirim"`
	NomorRekeningPengirim string  `json:"nomorRekeningPengirim"`
	NamaPengirim          string  `json:"namaPengirim"`
	NamaBankPenerima      string  `json:"namaBankPenerima"`
	NomorRekeningPenerima string  `json:"nomorRekeningPenerima"`
	NamaPenerima          string  `json:"namaPenerima"`
	CidBuktiPembayaran    string  `json:"cidBuktiPembayaran"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}

type PembayaranFindAllRequest struct {
	IdTransaksi string `json:"idTransaksi"`
}
