package model

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/xlzhangkeke/graphql-go-example/model/comment"
	"github.com/xlzhangkeke/graphql-go-example/model/post"
	"github.com/xlzhangkeke/graphql-go-example/model/user"
)

var Object Backend = new(BackendObject)

type BackendObject struct {
	*user.BackendUserObject
	*post.BackendPostObject
	*comment.BackendCommentObject
}

func NewBackendObject(db *sql.DB) *BackendObject {
	object := new(BackendObject)
	object.BackendUserObject = user.NewBackendUserObject(db)
	object.BackendPostObject = post.NewBackendPostObject(db)
	object.BackendCommentObject = comment.NewBackendCommentObject(db)
	return object
}

func InitDB(driverName, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return errors.Wrap(err, "open source failed")
	}
	Object = NewBackendObject(db)
	return nil
}
