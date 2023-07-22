package web

import "rantai-pasok-chaincode/helper"

type ConfirmContractRequest struct {
	IdPks          string        `json:"idPks"`
	IdKoperasi     string        `json:"idKoperasi"`
	IdKontrak      string        `json:"idKontrak"`
	Status         helper.Status `json:"status"`
	Pesan          string        `json:"pesan"`
	TanggalRespons string        `json:"tanggalRespons"`
}
