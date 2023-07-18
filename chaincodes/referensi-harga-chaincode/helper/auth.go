package helper

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func CheckAffiliation(ctx contractapi.TransactionContextInterface, allowedAffiliations []string) error {
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
