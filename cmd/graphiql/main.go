package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/schema"
)

func main() {
	if err := model.InitDB("postgres", "postgres://postgres:pandora@10.95.84.99:25432/graphql?sslmode=disable"); err != nil {
		log.Fatal(err)
	}
	if err := schema.InitSchema(); err != nil {
		log.Fatal(err)
	}
	log.Println("service starting:", time.Now().Format("2006-01-02 15:04:05"))
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	wrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//query, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	log.Println(err)
		//}
		//log.Println(string(query))
		log.Println("req:", time.Now().Format("2006-01-02 15:04:05"))
		h.ServeHTTP(w, r)
	})

	http.Handle("/graphql", wrapper)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
