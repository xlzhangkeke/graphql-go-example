package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/relay/examples/starwars"
)

func main() {

	// simplest relay-compliant graphql server HTTP handler
	// using Starwars schema from `graphql-relay-go` examples
	h := handler.New(&handler.Config{
		Schema: &starwars.Schema,
		Pretty: true,
	})
	wrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//query, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	log.Println(err)
		//}
		//log.Println(string(query))
		log.Println("req:", r.RequestURI, time.Now().Format("2006-01-02 15:04:05"))
		h.ServeHTTP(w, r)
	})
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(workDir)
	staticPath := "static"
	if !strings.Contains(workDir, "playground") {
		staticPath = filepath.Join(workDir, "cmd/playground/static")
	}
	log.Println(staticPath)
	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)
	http.Handle("/graphql", wrapper)
	http.ListenAndServe(":8080", nil)
}
