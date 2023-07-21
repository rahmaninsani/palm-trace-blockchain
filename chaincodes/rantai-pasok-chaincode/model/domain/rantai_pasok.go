package domain

import "rantai-pasok-chaincode/helper"

type RantaiPasok struct {
	Id      string                   `json:"id"`
	Kontrak Kontrak                  `json:"kontrak"`
	Status  helper.StatusRantaiPasok `json:"status"`
}
