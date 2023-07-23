package web

import "rantai-pasok-chaincode/constant"

type ConfirmContractRequest struct {
	IdPks          string                          `json:"idPks"`
	IdKoperasi     string                          `json:"idKoperasi"`
	IdKontrak      string                          `json:"idKontrak"`
	Status         constant.StatusPenawaranKontrak `json:"status"`
	Pesan          string                          `json:"pesan"`
	TanggalRespons string                          `json:"tanggalRespons"`
}
