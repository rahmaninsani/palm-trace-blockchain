package constant

type AssetType int

const (
	AssetTypeKebun         AssetType = iota // 0
	AssetTypeKontrak                        // 1
	AssetTypeDeliveryOrder                  // 2
	AssetTypeTransaksi                      // 3
	AssetTypeTransaksiItem                  // 4
	AssetTypePengiriman                     // 5
	AssetTypePenerimaan                     // 6
	AssetTypePembayaran                     // 7
)

func (assetType AssetType) String() string {
	switch assetType {
	case AssetTypeKebun:
		return "Kebun"
	case AssetTypeKontrak:
		return "Kontrak"
	case AssetTypeDeliveryOrder:
		return "Delivery Order"
	case AssetTypeTransaksi:
		return "Transaksi"
	case AssetTypeTransaksiItem:
		return "Transaksi Item"
	case AssetTypePengiriman:
		return "Pengiriman"
	case AssetTypePenerimaan:
		return "Penerimaan"
	case AssetTypePembayaran:
		return "Pembayaran"
	default:
		return "Unknown"
	}
}
