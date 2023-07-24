package constant

type AssetType int

const (
	AssetTypeKebun         AssetType = iota // 0
	AssetTypeKontrak                        // 1
	AssetTypeDeliveryOrder                  // 2
	AssetTypeTransaksi                      // 3
	AssetTypeTransaksiItem                  // 4
)

func (assetType AssetType) String() string {
	switch assetType {
	case AssetTypeKebun:
		return "Kebun"
	case AssetTypeKontrak:
		return "Kontrak"
	case AssetTypeDeliveryOrder:
		return "DeliveryOrder"
	case AssetTypeTransaksi:
		return "Transaksi"
	case AssetTypeTransaksiItem:
		return "TransaksiItem"
	default:
		return "Unknown"
	}
}
