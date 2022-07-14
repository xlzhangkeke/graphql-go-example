package model_test

import (
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/xlzhangkeke/graphql-go-example/model"
	"github.com/xlzhangkeke/graphql-go-example/model/user"
)

func TestMain(m *testing.M) {
	err := model.InitDB("postgres", "postgres://postgres:pandora@10.95.84.99:25432/graphql?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}

func TestBackendUser(t *testing.T) {
	t.Run("InsertUser", func(t *testing.T) {
		meta := new(user.User)
		//meta.ID = 1
		meta.Email = "email-xxxx-1@qq.com"
		err := model.Object.InsertUser(meta)
		assert.Nil(t, err)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		meta, err := model.Object.GetUserByID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, meta.ID)
	})
}

func TestBackendPost(t *testing.T) {
	t.Run("", func(t *testing.T) {

	})
}

func TestBackendComment(t *testing.T) {
	t.Run("", func(t *testing.T) {

	})
}
