package handlers

import (
	"context"
	"testing"

	"github.com/machinebox/graphql"
)

func TestCreateUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
			mutation {
				createUser(input : {
					nik:           "135",
					nama:          "ronda",
					alamat:        "taman anggrek",
					jenisKelamin:  "male",
					tanggalLahir:  "2011-01-02 15:04:05",
					agama:         "buddha"
				}){
					nik
				nama
				alamat
				jenisKelamin
				tanggalLahir
				agama
				}
			}
		`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

}

func TestUpdateUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation {
			updateUser(id : "2", input : {
				nik:           "13545",
				nama:          "asdfronda",
				alamat:        "taman anggrek",
				jenisKelamin:  "male",
				tanggalLahir:  "2011-01-02 15:04:05",
				agama:         "buddha"
			}){
				nik
				nama
				alamat
				jenisKelamin
				tanggalLahir
				agama
			}
		}
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation deleteUser {
			deleteUser(id : "1")
		}
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}
