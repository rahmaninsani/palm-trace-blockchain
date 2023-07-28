package web

type DeliveryOrderResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdKontrak             string  `json:"idKontrak"`
	Nomor                 string  `json:"nomor"`
	TanggalPembuatan      string  `json:"tanggalPembuatan"`
	Periode               string  `json:"periode"`
	Kuantitas             float32 `json:"kuantitas"`
	Harga                 float64 `json:"harga"`
	Rendemen              float32 `json:"rendemen"`
	Status                string  `json:"status"`
	Pesan                 string  `json:"pesan"`
	TanggalRespons        string  `json:"tanggalRespons"`
	KuantitasTepenuhi     float32 `json:"kuantitasTepenuhi"`
	KuantitasTersisa      float32 `json:"kuantitasTersisa"`
	CreatedAt             string  `json:"createdAt"`
	UpdatedAt             string  `json:"updatedAt"`
}
