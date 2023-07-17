package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ReferensiHargaChaincode struct {
	contractapi.Contract
}

type ReferensiHarga struct {
	Id               string  `json:"id"`
	IdDinas          []byte  `json:"idDinas"`
	UmurTanam        int     `json:"umurTanam"`
	Harga            float64 `json:"harga"`
	TanggalPembaruan string  `json:"tanggalPembaruan"`
}

type ReferensiHargaHistory struct {
	Timestamp string         `json:"timestamp"`
	Asset     ReferensiHarga `json:"asset"`
}

func (c *ReferensiHargaChaincode) Create(ctx contractapi.TransactionContextInterface, payload string) (string, error) {

	err := c.checkAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return "", fmt.Errorf("submitting client not authorized to create asset, does not have dinas.user affiliation/role")
	}

	var referensiHarga ReferensiHarga
	err = json.Unmarshal([]byte(payload), &referensiHarga)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal object: %v", err)
	}

	exists, err := c.assetExists(ctx, referensiHarga.Id)
	if err != nil {
		return "", fmt.Errorf("failed to get asset: %v", err)
	}

	if exists {
		return "", fmt.Errorf("the asset %s already exists", referensiHarga.Id)
	}

	referensiHargaJSON, err := json.Marshal(referensiHarga)
	if err != nil {
		return "", fmt.Errorf("failed to marshal asset: %v", err)
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(referensiHarga.Id, referensiHargaJSON)
}

func (c *ReferensiHargaChaincode) Update(ctx contractapi.TransactionContextInterface, id string, payload string) error {
	err := c.checkAffiliation(ctx, []string{"dinas.user"})
	if err != nil {
		return fmt.Errorf("submitting client not authorized to update asset, does not have dinas.user affiliation/role")
	}

	exists, err := c.assetExists(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}

	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	var updatedReferensiHargaAsset ReferensiHarga
	err = json.Unmarshal([]byte(payload), &updatedReferensiHargaAsset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal object: %v", err)
	}

	updatedReferensiHargaAssetJSON, err := json.Marshal(updatedReferensiHargaAsset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated asset: %v", err)
	}

	return ctx.GetStub().PutState(id, updatedReferensiHargaAssetJSON)
}

func (c *ReferensiHargaChaincode) Get(ctx contractapi.TransactionContextInterface, id string) (*ReferensiHarga, error) {
	err := c.checkAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	referensiHargaJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if referensiHargaJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var referensiHarga ReferensiHarga
	err = json.Unmarshal(referensiHargaJSON, &referensiHarga)
	if err != nil {
		return nil, err
	}

	return &referensiHarga, nil
}

func (c *ReferensiHargaChaincode) GetAll(ctx contractapi.TransactionContextInterface) ([]*ReferensiHarga, error) {
	err := c.checkAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var referensiHargaAssets []*ReferensiHarga
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var referensiHarga ReferensiHarga
		err = json.Unmarshal(queryResponse.Value, &referensiHarga)
		if err != nil {
			return nil, err
		}
		referensiHargaAssets = append(referensiHargaAssets, &referensiHarga)
	}

	return referensiHargaAssets, nil
}

func (c *ReferensiHargaChaincode) assetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (c *ReferensiHargaChaincode) checkAffiliation(ctx contractapi.TransactionContextInterface, allowedAffiliations []string) error {
	affiliation, isExist, err := ctx.GetClientIdentity().GetAttributeValue("hf.Affiliation")
	if !isExist || err != nil {
		return fmt.Errorf("failed to get client affiliation: %v", err)
	}

	isAllowed := false
	for _, allowedAffiliation := range allowedAffiliations {
		if affiliation == allowedAffiliation {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	return nil
}

func (c *ReferensiHargaChaincode) GetWithHistory(ctx contractapi.TransactionContextInterface, id string) ([]ReferensiHargaHistory, error) {
	err := c.checkAffiliation(ctx, []string{"dinas.user", "petani.user", "koperasi.user", "pabrikkelapasawit.user"})
	if err != nil {
		return nil, fmt.Errorf("submitting client not authorized to create asset, does not have the required affiliation/role")
	}

	referensiHargaJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if referensiHargaJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var referensiHargaAssets []ReferensiHargaHistory

	var latestReferensiHargaAsset ReferensiHarga
	err = json.Unmarshal(referensiHargaJSON, &latestReferensiHargaAsset)
	if err != nil {
		return nil, err
	}

	referensiHargaHistory := ReferensiHargaHistory{
		Timestamp: time.Now().Format(time.RFC3339),
		Asset:     latestReferensiHargaAsset,
	}
	referensiHargaAssets = append(referensiHargaAssets, referensiHargaHistory)

	// Get the history of previous versions
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for key %s: %v", id, err)
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate history for key %s: %v", id, err)
		}

		var previousReferensiHargaAsset ReferensiHarga
		err = json.Unmarshal(response.Value, &previousReferensiHargaAsset)
		if err != nil {
			return nil, err
		}

		referensiHargaHistory := ReferensiHargaHistory{
			Timestamp: response.Timestamp.String(),
			Asset:     previousReferensiHargaAsset,
		}
		referensiHargaAssets = append(referensiHargaAssets, referensiHargaHistory)
	}

	return referensiHargaAssets, nil
}
