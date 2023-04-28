package models

type Nasabah struct {
	Nama       string `json:"nama"`
	Nik        string `json:"nik"`
	NoHp       string `json:"no_hp"`
	NoRekening string `json:"no_rekening"`
}
