package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type Transaction struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	NomorKontrak       string `json:"nomorKontrak"`
	NomorDeliveryOrder string `json:"nomorDeliveryOrder"`
	NomorTransaksi     string `json:"nomorTransaksi"`
	IdPks              string `json:"idPks"`
	IdKoperasi         string `json:"idKoperasi"`
	IdPetani           string `json:"idPetani"`
	Kuantitas          int    `json:"kuantitas"`
	Harga              int    `json:"harga"`
	TanggalPengiriman  string `json:"tanggalPengiriman"`
	TanggalPenerimaan  string `json:"tanggalPenerimaan"`
	Status             string `json:"status"`
}

// CreateAsset issues a new asset to the world state with given details.
func (t *Transaction) CreateAsset(ctx contractapi.TransactionContextInterface, nomorKontrak, nomorDeliveryOrder, nomorTransaksi, idPks, idKoperasi, idPetani string, kuantitas, harga int, tanggalPengiriman, tanggalPenerimaan, status string) error {
	exists, err := t.AssetExists(ctx, nomorTransaksi)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", nomorTransaksi)
	}

	asset := Asset{
		NomorKontrak:       nomorKontrak,
		NomorDeliveryOrder: nomorDeliveryOrder,
		NomorTransaksi:     nomorTransaksi,
		IdPks:              idPks,
		IdKoperasi:         idKoperasi,
		IdPetani:           idPetani,
		Kuantitas:          kuantitas,
		Harga:              harga,
		TanggalPengiriman:  tanggalPengiriman,
		TanggalPenerimaan:  tanggalPenerimaan,
		Status:             status,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(nomorTransaksi, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (t *Transaction) ReadAsset(ctx contractapi.TransactionContextInterface, nomorTransaksi string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(nomorTransaksi)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", nomorTransaksi)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (t *Transaction) AssetExists(ctx contractapi.TransactionContextInterface, nomorTransaksi string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(nomorTransaksi)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetAllAssets returns all assets found in world state
func (t *Transaction) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
