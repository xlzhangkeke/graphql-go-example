package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/model/comment"
	"strconv"
)

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*comment.Comment); ok {
					return meta.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*comment.Comment); ok {
					return meta.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*comment.Comment); ok {
					return meta.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CommentType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*comment.Comment); ok {
				return model.Object.GetUserByID(meta.UserID)
			}
			return nil, nil
		},
	})

	CommentType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "Post ID",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if postID, ok := p.Args["id"].(string); ok {
				id, err := strconv.Atoi(postID)
				if err != nil {
					return nil, err
				}
				return model.Object.GetPostByID(id)
			}
			return nil, nil
		},
	})
}
