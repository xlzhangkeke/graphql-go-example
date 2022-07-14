package model

import (
	"github.com/xlzhangkeke/graphql-go-example/model/comment"
	"github.com/xlzhangkeke/graphql-go-example/model/post"
	"github.com/xlzhangkeke/graphql-go-example/model/user"
)

type Backend interface {
	user.BackendUser
	post.BackendPost
	comment.BackendComment
}
