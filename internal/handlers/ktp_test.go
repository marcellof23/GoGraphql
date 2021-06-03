package handlers

// import (
// 	"context"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/matryer/is"
// )

// func TestCreateUser(t *testing.T) {
// 	is := is.New(t)

// 	var calls int
// 	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		calls++
// 		query := r.FormValue("query")
// 		is.Equal(query, "query {}")
// 		is.Equal(r.FormValue("variables"), `{"username":"matryer"}`+"\n")
// 		_, err := io.WriteString(w, `{"data":{"value":"some data"}}`)
// 		is.NoErr(err)
// 	}))
// 	defer srv.Close()
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	client := NewClient(srv.URL, UseMultipartForm())

// 	req := NewRequest("query {}")
// 	req.Var("username", "matryer")

// 	// check variables
// 	is.True(req != nil)
// 	is.Equal(req.vars["username"], "matryer")

// 	var resp struct {
// 		Value string
// 	}
// 	err := client.Run(ctx, req, &resp)
// 	is.NoErr(err)
// 	is.Equal(calls, 1)

// 	is.Equal(resp.Value, "some data")
// }
