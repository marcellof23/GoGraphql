package handlers

import (
	"context"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/matryer/is"
)

// TestCreateUser create dummy user testing
func TestCreateUser(t *testing.T) {
	is := is.New(t)

	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
			mutation {
				createUser(input : {
					nik:           "131125",
					nama:          "aliando",
					alamat:        "taman anggrek",
					jenis_kelamin: "male",
					tanggal_lahir:  "2011-01-02 15:04:05",
					agama:         "buddha"
				}){
				nik
				nama
				alamat
				jenis_kelamin
				tanggal_lahir
				agama
				}
			}
		`)

	is.True(req != nil)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

}

// TestUpdateUser update dummy user testing
func TestUpdateUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation {
			updateUser(id : "5", input : {
				nik:           "1354521",
				nama:          "wewe",
				alamat:        "taman anggrek",
				jenis_kelamin:  "male",
				tanggal_lahir:  "2011-01-02 15:04:05",
				agama:         "islam"
			}){
				nik
				nama
				alamat
				jenis_kelamin
				tanggal_lahir
				agama
			}
		}
	`)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

// TestDeleteUser delete dummy user testing
func TestDeleteUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation deleteUser {
			deleteUser(id : "1")
		}
	`)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

// TestPaginationUser test the user pagination
func TestPaginationUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8082/query")

	// success test with output shown
	var req = graphql.NewRequest(`
		query {
			getPagination(
				input: { first: 2, offset: 2, query: "ali", sort : ["-id"] }
			  ) {
				edges {
				  node {
					id
					nik
					nama
					alamat
					jenis_kelamin
					tanggal_lahir
					agama
				  }
				  cursor
				}
				totalCount
				pageInfo {
				  endCursor
				  hasNextPage
				}
			  }
		}	  
	`)

	//failed test with output empty
	var req2 = graphql.NewRequest(`
		query {
			getPagination(
				input: { first: 2, offset: 2, query: "ali", sort : ["-id"] }
			  ) {
				edges {
				  node {
					id
					nik
					nama
					alamat
					jenis_kelamin
					tanggal_lahir
					agama
				  }
				  cursor
				}
				totalCount
				pageInfo {
				  endCursor
				  hasNextPage
				}
			  }
		}
	`)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
	if err := client.Run(ctx, req2, &respData); err != nil {
		t.Error(err)
	}
}
