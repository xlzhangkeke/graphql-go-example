package comment

import "database/sql"

var _ BackendComment = new(BackendCommentObject)

type BackendCommentObject struct {
	db   *sql.DB
	meta *Comment
}

func NewBackendCommentObject(db *sql.DB) *BackendCommentObject {
	return &BackendCommentObject{
		db:   db,
		meta: new(Comment),
	}
}
func (object *BackendCommentObject) InsertComment(comment *Comment) error {
	var id int
	err := object.db.QueryRow(`
		INSERT INTO comments(user_id, post_id, title, body)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, comment.UserID, comment.PostID, comment.Title, comment.Body).Scan(&id)
	if err != nil {
		return err
	}
	comment.ID = id
	return nil
}

func (object *BackendCommentObject) RemoveCommentByID(id int) error {
	_, err := object.db.Exec("DELETE FROM comments WHERE id=$1", id)
	return err
}

func (object *BackendCommentObject) GetCommentByIDAndPost(id int, postID int) (*Comment, error) {
	var (
		userID      int
		title, body string
	)
	err := object.db.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=$1
		AND post_id=$2
	`, id, postID).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &Comment{
		ID:     id,
		UserID: userID,
		PostID: postID,
		Title:  title,
		Body:   body,
	}, nil
}

func (object *BackendCommentObject) GetCommentsForPost(id int) ([]*Comment, error) {
	rows, err := object.db.Query(`
		SELECT c.id, c.user_id, c.title, c.body
		FROM comments AS c
		WHERE c.post_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		comments    = []*Comment{}
		cid, userID int
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&cid, &userID, &title, &body); err != nil {
			return nil, err
		}
		comments = append(comments, &Comment{
			ID:     cid,
			UserID: userID,
			PostID: id,
			Title:  title,
			Body:   body,
		})
	}
	return comments, nil
}
