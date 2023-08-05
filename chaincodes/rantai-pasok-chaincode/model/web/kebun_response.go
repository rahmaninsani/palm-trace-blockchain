package web

type KebunResponse struct {
	IdTransaksiBlockchain         string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                            string  `json:"id"`
	IdPetani                      string  `json:"idPetani"`
	Alamat                        string  `json:"alamat"`
	Latitude                      string  `json:"latitude"`
	Longitude                     string  `json:"longitude"`
	Luas                          float32 `json:"luas"`
	KemampuanProduksiHarian       float32 `json:"kemampuanProduksiHarian"`
	NomorSuratKeteranganLurah     string  `json:"nomorSuratKeteranganLurah"`
	CidSuratKeteranganLurah       string  `json:"cidSuratKeteranganLurah"`
	NomorSuratKeteranganGantiRugi string  `json:"nomorSuratKeteranganGantiRugi"`
	CidSuratKeteranganGantiRugi   string  `json:"cidSuratKeteranganGantiRugi"`
	NomorSertifikatHakMilik       string  `json:"nomorSertifikatHakMilik"`
	CidSertifikatHakMilik         string  `json:"cidSertifikatHakMilik"`
	NomorSuratTandaBudidaya       string  `json:"nomorSuratTandaBudidaya"`
	CidSuratTandaBudidaya         string  `json:"cidSuratTandaBudidaya"`
	NomorSertifikatRspo           string  `json:"nomorSertifikatRspo"`
	CidSertifikatRspo             string  `json:"cidSertifikatRspo"`
	NomorSertifikatIspo           string  `json:"nomorSertifikatIspo"`
	CidSertifikatIspo             string  `json:"cidSertifikatIspo"`
	NomorSertifikatIscc           string  `json:"nomorSertifikatIscc"`
	CidSertifikatIscc             string  `json:"cidSertifikatIscc"`
	CreatedAt                     string  `json:"createdAt"`
	UpdatedAt                     string  `json:"updatedAt"`
}
