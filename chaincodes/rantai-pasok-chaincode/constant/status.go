package constant

type StatusRantaiPasok int

const (
	RantaiPasokBerjalan StatusRantaiPasok = iota // 0
	RantaiPasokSelesai                           // 1
)

func (status StatusRantaiPasok) String() string {
	switch status {
	case RantaiPasokBerjalan:
		return "Berjalan"
	case RantaiPasokSelesai:
		return "Selesai"
	default:
		return "Status Tidak Diketahui"
	}
}

type StatusPenawaranKontrak int

const (
	PenawaranKontrakMenungguKonfirmasi StatusPenawaranKontrak = iota // 0
	PenawaranKontrakDisetujui                                        // 1
	PenawaranKontrakDitolak                                          // 2
)

func (statusPenawaranKontrak StatusPenawaranKontrak) String() string {
	switch statusPenawaranKontrak {
	case PenawaranKontrakMenungguKonfirmasi:
		return "Menunggu Konfirmasi"
	case PenawaranKontrakDisetujui:
		return "Disetujui"
	case PenawaranKontrakDitolak:
		return "Ditolak"
	default:
		return "Status Tidak Diketahui"
	}
}

type StatusPenawaranDeliveryOrder int

const (
	PenawaranDeliveryOrderMenungguKonfirmasi StatusPenawaranDeliveryOrder = iota // 0
	PenawaranDeliveryOrderDisetujui                                              // 1
	PenawaranDeliveryOrderDitolak                                                // 2
)

func (statusPenawaranKontrak StatusPenawaranDeliveryOrder) String() string {
	switch statusPenawaranKontrak {
	case PenawaranDeliveryOrderMenungguKonfirmasi:
		return "Menunggu Konfirmasi"
	case PenawaranDeliveryOrderDisetujui:
		return "Disetujui"
	case PenawaranDeliveryOrderDitolak:
		return "Ditolak"
	default:
		return "Status Tidak Diketahui"
	}
}

type StatusPenawaranTransaksi int

const (
	PenawaranTransaksiMenungguKonfirmasi StatusPenawaranTransaksi = iota // 0
	PenawaranTransaksiDisetujui                                          // 1
	PenawaranTransaksiDitolak                                            // 2
)

func (statusPenawaranTransaksi StatusPenawaranTransaksi) String() string {
	switch statusPenawaranTransaksi {
	case PenawaranTransaksiMenungguKonfirmasi:
		return "Menunggu Konfirmasi"
	case PenawaranTransaksiDisetujui:
		return "Disetujui"
	case PenawaranTransaksiDitolak:
		return "Ditolak"
	default:
		return "Status Tidak Diketahui"
	}
}

type StatusTransaksi int

const (
	TransaksiDitawarkanPetani  StatusTransaksi = iota // 0
	TransaksiDisetujuiKoperasi                        // 1
	TransaksiDitolakKoperasi                          // 2
	TransaksiDisetujuiPks                             // 3
	TransaksiDitolakPks                               // 4
	TransaksiDikirimPetani                            // 5
	TransaksiDiterimaKoperasi                         // 6
	TransaksiDikirimKoperasi                          // 7
	TransaksiDiterimaPks                              // 8
	TransaksiDibayarPks                               // 9
	TransaksiDibayarKoperasi                          // 10
	TransaksiSelesai                                  // 11
)

func (status StatusTransaksi) String() string {
	switch status {
	case TransaksiDitawarkanPetani:
		return "Ditawarkan Petani"
	case TransaksiDisetujuiKoperasi:
		return "Disetujui Koperasi"
	case TransaksiDitolakKoperasi:
		return "Ditolak Koperasi"
	case TransaksiDisetujuiPks:
		return "Disetujui Pabrik Kelapa Sawit"
	case TransaksiDitolakPks:
		return "Ditolak Pabrik Kelapa Sawit"
	case TransaksiDikirimPetani:
		return "Dikirim Petani"
	case TransaksiDiterimaKoperasi:
		return "Diterima Koperasi"
	case TransaksiDikirimKoperasi:
		return "Dikirim Koperasi"
	case TransaksiDiterimaPks:
		return "Diterima Pabrik Kelapa Sawit"
	case TransaksiDibayarPks:
		return "Dibayar Pabrik Kelapa Sawit"
	case TransaksiDibayarKoperasi:
		return "Dibayar Koperasi"
	case TransaksiSelesai:
		return "Selesai"
	default:
		return "Status Tidak Diketahui"
	}
}
