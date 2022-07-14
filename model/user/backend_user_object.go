package user

import "database/sql"

var _ BackendUser = new(BackendUserObject)

type BackendUserObject struct {
	db   *sql.DB
	meta *User
}

func NewBackendUserObject(db *sql.DB) *BackendUserObject {
	return &BackendUserObject{
		db:   db,
		meta: new(User),
	}
}

func (object *BackendUserObject) GetUserByID(id int) (*User, error) {
	var email string
	err := object.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func (object *BackendUserObject) InsertUser(user *User) error {
	var id int
	err := object.db.QueryRow(`
		INSERT INTO users(email)
		VALUES ($1)
		RETURNING id
	`, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (object *BackendUserObject) RemoveUserByID(id int) error {
	_, err := object.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func (object *BackendUserObject) Follow(followerID, followeeID int) error {
	_, err := object.db.Exec(`
		INSERT INTO followers(follower_id, followee_id)
		VALUES ($1, $2)
	`, followerID, followeeID)
	return err
}

func (object *BackendUserObject) Unfollow(followerID, followeeID int) error {
	_, err := object.db.Exec(`
		DELETE FROM followers
		WHERE follower_id=$1
		AND followee_id=$2
	`, followerID, followeeID)
	return err
}

func (object *BackendUserObject) GetFollowerByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := object.db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.follower_id
		AND f.follower_id=$1
		AND f.followee_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func (object *BackendUserObject) GetFollowersForUser(id int) ([]*User, error) {
	rows, err := object.db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.followee_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: uid, Email: email})
	}
	return users, nil
}

func (object *BackendUserObject) GetFolloweeByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := object.db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.followee_id
		AND f.followee_id=$1
		AND f.follower_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func (object *BackendUserObject) GetFolloweesForUser(id int) ([]*User, error) {
	rows, err := object.db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.followee_id
		AND f.follower_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}
