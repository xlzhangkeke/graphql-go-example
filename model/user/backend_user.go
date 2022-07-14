package user

type BackendUser interface {
	InsertUser(user *User) error
	GetUserByID(id int) (*User, error)
	RemoveUserByID(id int) error
	Follow(followerID, followeeID int) error
	Unfollow(followerID, followeeID int) error
	GetFollowerByIDAndUser(id int, userID int) (*User, error)
	GetFollowersForUser(id int) ([]*User, error)
	GetFolloweeByIDAndUser(id int, userID int) (*User, error)
	GetFolloweesForUser(id int) ([]*User, error)
}
