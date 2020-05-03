package user

import (
	"time"

	"github.com/go-pg/pg/v9"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/models"
	"github.com/absinsekt/pnk/utils/core"
)

// User type
type User struct {
	tableName   struct{}  `pg:"auth_user"`
	ID          string    `pg:",pk" json:"id"`
	Password    string    `pg:",notnull" json:"-"`
	LastLogin   time.Time `json:"-"`
	IsSuperuser bool      `pg:",notnull" json:"-"`
	Username    string    `pg:",notnull,unique" json:"username"`
	FirstName   string    `pg:",notnull" json:"-"`
	LastName    string    `pg:",notnull" json:"-"`
	Email       string    `pg:",notnull" json:"email"`
	IsStaff     bool      `pg:",notnull" json:"-"`
	IsActive    bool      `pg:",notnull" json:"-"`
	DateJoined  time.Time `pg:",notnull" json:"-"`
}

// Auth authenticates user (checks if it exists with a given password)
func Auth(username, password string) (*User, error) {
	user, err := getUser(username)
	if err != nil {
		return nil, err
	}

	if core.DjangoPasswordEquals(password, user.Password) {
		return user, nil
	}

	return nil, pg.ErrNoRows
}

func getUser(username string) (*User, error) {
	method := func() (interface{}, error) {
		user := new(User)

		if err := models.DB.
			Model(user).
			Where("username = ?", username).
			Select(); err != nil {

			return nil, err
		}

		return user, nil
	}

	result, err := core.GetCached("GetUser", configuration.SecondsFrequently, method)
	if err != nil {
		return nil, err
	}

	return result.(*User), nil
}
