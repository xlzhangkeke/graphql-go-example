package comment

type BackendComment interface {
	InsertComment(comment *Comment) error
	RemoveCommentByID(id int) error
	GetCommentByIDAndPost(id int, postID int) (*Comment, error)
	GetCommentsForPost(id int) ([]*Comment, error)
}
