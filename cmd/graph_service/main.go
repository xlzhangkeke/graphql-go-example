package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/schema"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(string(query))
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func main() {
	if err := model.InitDB("postgres", "postgres://postgres:pandora@10.95.84.99:25432/graphql?sslmode=disable"); err != nil {
		log.Fatal(err)
	}
	if err := schema.InitSchema(); err != nil {
		log.Fatal(err)
	}
	log.Println("service starting:", time.Now().Format("2006-01-02 15:04:05"))
	http.Handle("/graphql", handler(schema.Schema))
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
