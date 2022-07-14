package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

var Schema graphql.Schema

func InitSchema() error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		return errors.Wrap(err, "create schema failed")
	}
	Schema = schema
	return nil
}
