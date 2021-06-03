package models

import "time"

type User struct {
	ID            int64
	Nama          string
	NIK           string
	Jenis_kelamin string
	Tanggal_lahir string
	Alamat        string
	Agama         string
	Created_at    time.Time
	Updated_at    time.Time
}

// Tablename is name of the table
func (User) TableName() string {
	return "user"
}
