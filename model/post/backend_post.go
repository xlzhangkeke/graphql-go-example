package post

type BackendPost interface {
	InsertPost(post *Post) error
	RemovePostByID(id int) error
	GetPostByID(id int) (*Post, error)
	GetPostByIDAndUser(id, userID int) (*Post, error)
	GetPostsForUser(id int) ([]*Post, error)
}
