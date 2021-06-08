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

type PaginationResultUser struct {
	totalCount int64
	edges      []PaginationEdge
	pageInfo   PaginationInfo
}

type PaginationEdge struct {
	node   User
	cursor int64
}

type PaginationInfo struct {
	endCursor   int64
	hasNextPage bool
}

// Tablename is name of the table
func (User) TableName() string {
	return "user"
}
