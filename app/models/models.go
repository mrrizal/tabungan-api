package models

import "time"

type DaftarResponseOk struct {
	NoRekening string `json:"no_rekening"`
}

type ErrorResponse struct {
	Remark string `json:"remark"`
}

type Nasabah struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type Rekening struct {
	NoRekening string  `json:"no_rekening"`
	NasabahID  int     `json:"nasabah_id"`
	Saldo      float64 `json:"saldo"`
}

type TransaksiRequest struct {
	NoRekening   string  `json:"no_rekening"`
	Nominal      float64 `json:"nominal"`
	CurrentSaldo float64 `json:"current_saldo"`
	Type         string  `json:"type"`
}

type TransaksiResponseOk struct {
	Saldo float64 `json:"saldo"`
}

type MutasiTransaksi struct {
	Waktu         time.Time `json:"waktu"`
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       float64   `json:"nominal"`
}

type MutasiResp struct {
	Results []MutasiTransaksi `json:"results"`
}
