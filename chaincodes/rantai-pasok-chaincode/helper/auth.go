package helper

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func CheckAffiliation(ctx contractapi.TransactionContextInterface, allowedAffiliations []string) error {
	affiliation, isExist, err := ctx.GetClientIdentity().GetAttributeValue("hf.Affiliation")
	if err != nil {
		return fmt.Errorf("failed to get client affiliation: %v", err)
	}

	if !isExist {
		return fmt.Errorf("client affiliation is not found")
	}

	isAllowed := false
	for _, allowedAffiliation := range allowedAffiliations {
		if affiliation == allowedAffiliation {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("client affiliation %s is not allowed", affiliation)
	}

	return nil
}
