package models

type DaftarResponseOk struct {
	NoRekening string `json:"no_rekening"`
}

type DaftarResponseError struct {
	Remark string `json:"remark"`
}
