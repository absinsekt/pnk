package user

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"

	"github.com/absinsekt/pnk/lib/crypto"
	"github.com/absinsekt/pnk/lib/strings"
	"github.com/absinsekt/pnk/models"
)

// User type
type User struct {
	tableName   struct{}  `pg:"auth_user"`
	ID          int64     `pg:",pk" json:"id"`
	Password    string    `pg:"type:varchar(128),notnull" json:"-"`
	LastLogin   time.Time `json:"-"`
	Username    string    `pg:"type:varchar(150),notnull,unique" json:"username"`
	FirstName   string    `pg:"type:varchar(30),notnull" json:"-"`
	LastName    string    `pg:"type:varchar(30),notnull" json:"-"`
	Email       string    `pg:"type:varchar(254),notnull" json:"email"`
	IsSuperuser bool      `pg:",notnull" json:"-"`
	IsStaff     bool      `pg:",notnull" json:"-"`
	IsActive    bool      `pg:",notnull" json:"-"`
	DateJoined  time.Time `pg:",notnull" json:"-"`
}

// SessionData type to store user data in session
type SessionData struct {
	ID             int64
	Username       string
	Email          string
	IsStaff        bool
	SessionVersion string
}

// Auth authenticates user (checks if it exists with a given password)
func Auth(username, password string) (*User, error) {
	user, err := getActiveUser(username)
	if err != nil {
		return nil, err
	}

	if crypto.DjangoPasswordEquals(password, user.Password) {
		updateLastLogin(user)
		return user, nil
	}

	return nil, pg.ErrNoRows
}

func getActiveUser(username string) (*User, error) {
	// NO CACHE (isActive, permissions)
	user := new(User)

	if err := models.DB.
		Model(user).
		Where("username = ?", username).
		Where("is_active = ?", true).
		Select(); err != nil {

		return nil, err
	}

	return user, nil
}

func updateLastLogin(user *User) error {
	user.LastLogin = time.Now()
	return models.DB.Update(user)
}

// CreateUser creates a new active user with a given password
func CreateUser(username, password, firstName, lastName, email string, isStaff, isSuperuser bool) error {
	salt := strings.GenerateRandomString(12)
	iterations := 12000

	if isStaff {
		iterations = 24000
	}

	if isSuperuser {
		iterations = 30000
	}

	user := &User{
		Username:    username,
		Password:    crypto.GetDjangoPasswordHash(password, salt, iterations),
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		DateJoined:  time.Now(),
		IsSuperuser: isSuperuser,
		IsStaff:     isStaff,
		IsActive:    true,
	}

	return models.DB.Insert(user)
}

func GetList() []User {
	users := make([]User, 1)

	models.DB.
		Model(users).
		Where("is_active = ?", true).
		Select()

	return users
}

// maintenance functions to be run with go test
func createUserTable() error {
	var (
		user = new(User)
		err  error
	)

	if err = models.DB.CreateTable(user, &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}); err != nil {
		return err
	}

	_, err = models.DB.Exec(`CREATE INDEX auth_user_username_like
		ON public.auth_user USING btree
		(username COLLATE pg_catalog."default" varchar_pattern_ops ASC NULLS LAST)
		TABLESPACE pg_default;`)

	return err
}

func dropUserTable() error {
	var user = new(User)

	return models.DB.DropTable(user, &orm.DropTableOptions{
		Cascade:  true,
		IfExists: true,
	})
}
