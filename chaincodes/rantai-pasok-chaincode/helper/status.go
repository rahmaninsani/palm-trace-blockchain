package helper

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

type Status int

const (
	MenungguKonfirmasi Status = iota // 0
	Disetujui                        // 1
	Ditolak                          // 2
)

func (status Status) String() string {
	switch status {
	case MenungguKonfirmasi:
		return "Menunggu Konfirmasi"
	case Disetujui:
		return "Disetujui"
	case Ditolak:
		return "Ditolak"
	default:
		return "Status Tidak Diketahui"
	}
}

type StatusTransaksi int

const (
	DitawarkanPetani  StatusTransaksi = iota // 0
	DisetujuiKoperasi                        // 1
	DitolakKoperasi                          // 2
	DisetujuiPks                             // 3
	DitolakPks                               // 4
	DikirimPetani                            // 5
	DiterimaKoperasi                         // 6
	DikirimKoperasi                          // 7
	DiterimaPks                              // 8
	DibayarPks                               // 9
	DibayarKoperasi                          // 10
	TransaksiSelesai                         // 11
)

func (status StatusTransaksi) String() string {
	switch status {
	case DitawarkanPetani:
		return "Ditawarkan Petani"
	case DisetujuiKoperasi:
		return "Disetujui Koperasi"
	case DitolakKoperasi:
		return "Ditolak Koperasi"
	case DisetujuiPks:
		return "Disetujui Pabrik Kelapa Sawit"
	case DitolakPks:
		return "Ditolak Pabrik Kelapa Sawit"
	case DikirimPetani:
		return "Dikirim Petani"
	case DiterimaKoperasi:
		return "Diterima Koperasi"
	case DikirimKoperasi:
		return "Dikirim Koperasi"
	case DiterimaPks:
		return "Diterima Pabrik Kelapa Sawit"
	case DibayarPks:
		return "Dibayar Pabrik Kelapa Sawit"
	case DibayarKoperasi:
		return "Dibayar Koperasi"
	case TransaksiSelesai:
		return "Selesai"
	default:
		return "Status Tidak Diketahui"
	}
}
