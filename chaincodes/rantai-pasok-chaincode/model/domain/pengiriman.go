package domain

type Pengiriman struct {
	Id                   string `json:"id"`
	IdTransaksi          string `json:"idTransaksi"`
	Nomor                string `json:"nomor"`
	Tanggal              string `json:"tanggal"`
	NamaSopir            string `json:"namaSopir"`
	NomorTeleponSopir    string `json:"nomorTeleponSopir"`
	NamaKendaraan        string `json:"namaKendaraan"`
	NomorPolisiKendaraan string `json:"nomorPolisiKendaraan"`
}
