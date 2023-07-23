package domain

import "rantai-pasok-chaincode/constant"

type RantaiPasok struct {
	Id      string                     `json:"id"`
	Kontrak Kontrak                    `json:"kontrak"`
	Status  constant.StatusRantaiPasok `json:"status"`
}
