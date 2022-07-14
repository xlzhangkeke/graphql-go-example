package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/model/comment"
	"github.com/xlzhangkeke/graphql-go-example/model/post"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Post",
	Interfaces: nil,
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*post.Post); ok {
					return meta.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*post.Post); ok {
					return meta.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*post.Post); ok {
					return meta.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*post.Post); ok {
				return model.Object.GetUserByID(meta.UserID)
			}
			return nil, nil
		},
	})

	PostType.AddFieldConfig("comment", &graphql.Field{
		Type: CommentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*post.Post); ok {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return model.Object.GetCommentByIDAndPost(id, meta.ID)
			}
			return nil, nil
		},
	})

	PostType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CommentType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*post.Post); ok {
				return model.Object.GetCommentsForPost(meta.ID)
			}
			return []comment.Comment{}, nil
		},
	})
}
