package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/model/post"
	"github.com/xlzhangkeke/graphql-go-example/model/user"
	"strconv"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*user.User); ok {
					return meta.ID, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if meta, ok := p.Source.(*user.User); ok {
					return meta.Email, nil
				}
				return nil, nil
			}},
	},
})

func init() {

	UserType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return model.Object.GetPostByIDAndUser(id, meta.ID)
			}
			return nil, nil
		},
	})

	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PostType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok {
				return model.Object.GetPostsForUser(meta.ID)
			}
			return []post.Post{}, nil
		},
	})

	UserType.AddFieldConfig("follower", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Follower ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return model.Object.GetFollowerByIDAndUser(id, meta.ID)
			}
			return nil, nil
		},
	})

	UserType.AddFieldConfig("followers", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok == true {
				return model.Object.GetFollowersForUser(meta.ID)
			}
			return []user.User{}, nil
		},
	})

	UserType.AddFieldConfig("followee", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Followee ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return model.Object.GetFolloweeByIDAndUser(id, meta.ID)
			}
			return nil, nil
		},
	})

	UserType.AddFieldConfig("followees", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if meta, ok := p.Source.(*user.User); ok {
				return model.Object.GetFolloweesForUser(meta.ID)
			}
			return []user.User{}, nil
		},
	})
}
