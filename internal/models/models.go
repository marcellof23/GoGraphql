package models

import "time"

type User struct {
	ID           int64
	Nama         string
	NIK          string
	JenisKelamin string
	TanggalLahir string
	Alamat       string
	Agama        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Tablename is name of the table
func (User) TableName() string {
	return "user"
}
