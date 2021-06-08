package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/marcellof23/GoGraphql/configs/database"
	"github.com/marcellof23/GoGraphql/internal/models"
	"github.com/stretchr/testify/assert"
)

const defaultUrl = "http://localhost:8082/query"

// TestCreateUser create dummy user testing
func TestCreateUser(t *testing.T) {

	var client = graphql.NewClient(fmt.Sprintf("%s", defaultUrl))

	//Success test with output created as from input
	var req = graphql.NewRequest(`
			mutation {
				createUser(input : {
					nik:           "13012",
					nama:          "antares",
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

	//failed test with output empty by not filling nik
	var req2 = graphql.NewRequest(`
		mutation {
			createUser(input : {
				nama:          "antares",
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

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	err2 := client.Run(ctx, req2, &respData)
	assert.NotNil(t, err2)

	nik := respData["createUser"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "13012", "nik should be equal")

	nama := respData["createUser"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "antares", "nama should be equal")

	agama := respData["createUser"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "buddha", "agama should be equal")

}

// TestUpdateUser update dummy user testing
func TestUpdateUser(t *testing.T) {
	var client = graphql.NewClient(fmt.Sprintf("%s", defaultUrl))

	//Success test with databsae updated refer from input
	var req = graphql.NewRequest(`
		mutation {
			updateUser(id : "5", input : {
				nik:           "1354521",
				nama:          "wews",
				alamat:        "taman anggrek",
				jenis_kelamin:  "male",
				tanggal_lahir:  "2011-01-02 15:04:05",
				agama:         "katolik"
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

	//failed test with output empty by not filling nama and alamat
	var req2 = graphql.NewRequest(`
		mutation {
			updateUser(id : "6", input : {
				nik:           "1300",
				jenis_kelamin:  "male",
				tanggal_lahir:  "2011-01-02 15:04:05",
				agama:         "katolik"
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

	err2 := client.Run(ctx, req2, &respData)
	assert.NotNil(t, err2)

	nik := respData["updateUser"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "1354521", "nik should be equal")

	nama := respData["updateUser"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "wews", "nama should be equal")

	agama := respData["updateUser"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "katolik", "agama should be equal")

}

// TestSuccessDeleteUser delete dummy user testing
func TestSuccessDeleteUser(t *testing.T) {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	var client = graphql.NewClient("http://localhost:8082/query")

	user := models.User{}
	rs := database.DB.First(&user)
	assert.Nil(t, rs.Error)
	id := user.ID

	var req = graphql.NewRequest(`
		mutation {
			deleteUser(id:"` + strconv.Itoa(int(id)) + `")
		}
	`)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
	fmt.Print(respData)

	isDeleted := respData["deleteUser"]
	assert.Equal(t, isDeleted, true, "Must be deleted")
}

// TestFailDeleteUser delete dummy user testing
func TestFailDeleteUser(t *testing.T) {
	var client = graphql.NewClient(fmt.Sprintf("%s", defaultUrl))

	var req = graphql.NewRequest(`
		mutation deleteUser {
			deleteUser(id : "20")
		}
	`)

	ctx := context.Background()

	// err to run and catch the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		isDeleted := respData["deleteUser"]
		assert.Nil(t, isDeleted)
	}

}

// TestPaginationUser test the user pagination
func TestPaginationUser(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8082/query")

	// success test to create user with output shown
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

	//failed test with output empty by add offset more than total data
	var req2 = graphql.NewRequest(`
		query {
			getPagination(
				input: { first: 2, offset: 100, query: "ali", sort : ["-id"] }
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
	err2 := client.Run(ctx, req2, &respData)
	assert.NotNil(t, err2)
}
