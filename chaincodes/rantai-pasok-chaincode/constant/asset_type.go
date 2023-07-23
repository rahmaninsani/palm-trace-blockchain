package constant

type AssetType int

const (
	AssetTypeKontrak       AssetType = iota // 0
	AssetTypeDeliveryOrder                  // 1
	AssetTypeTransaksi                      // 2
	AssetTypeTransaksiItem                  // 3
)

func (assetType AssetType) String() string {
	switch assetType {
	case AssetTypeKontrak:
		return "Kontrak"
	case AssetTypeDeliveryOrder:
		return "Delivery Order"
	case AssetTypeTransaksi:
		return "Transaksi"
	case AssetTypeTransaksiItem:
		return "Transaksi Item"
	default:
		return "Status Tidak Diketahui"
	}
}
