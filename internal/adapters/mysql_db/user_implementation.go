// connect to core and not to application

package mysqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/xHappyface/social/core/errors"
	"github.com/xHappyface/social/core/user"
	"github.com/xHappyface/social/internal/ports"
)

// create a structure that will fulfill the needs of the core port
// inject dependencies, such as the databse, here too

type UserRepoImpl struct {
	db         *sql.DB
	ctxTimeout time.Duration
}

// fulfilling the port interface will allow you to return the above structure
// in place of the core port type

func NewUserRepoImpl(db *sql.DB, ctxTimeout time.Duration) ports.UserRepository {
	return &UserRepoImpl{
		db:         db,
		ctxTimeout: ctxTimeout,
	}
}

// fulfill the port interface

func (repo *UserRepoImpl) Create(user *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.ctxTimeout)
	defer cancel()
	stmt := fmt.Sprintf("insert into users(id, username, password, date_created) values (%q, %q, %q, %q);", user.ID, user.Username,
		user.Password, user.DateCreated.Format("2006-01-02"))
	result, err := repo.db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.ErrNoRowsAffected
	}
	return nil
}

func (repo *UserRepoImpl) ReadByID(userID string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.ctxTimeout)
	defer cancel()
	stmt := fmt.Sprintf("select * from users where id=%q;", userID)
	rows, err := repo.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.ErrNoRowsRetrieved
	}
	var usr user.User
	if err = rows.Scan(&usr.ID, &usr.Username, &usr.Password, &usr.DateCreated); err != nil {
		return nil, err
	}
	if rows.NextResultSet() {
		return nil, errors.ErrUnexpectedNextResultSet
	}
	return &usr, nil
}

func (repo *UserRepoImpl) ReadByName(userName string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.ctxTimeout)
	defer cancel()
	stmt := fmt.Sprintf("select * from users where username=%q;", userName)
	rows, err := repo.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.ErrNoRowsRetrieved
	}
	var usr user.User
	if err = rows.Scan(&usr.ID, &usr.Username, &usr.Password, &usr.DateCreated); err != nil {
		return nil, err
	}
	if rows.NextResultSet() {
		return nil, errors.ErrUnexpectedNextResultSet
	}
	return &usr, nil
}

func (repo *UserRepoImpl) Update(usr *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.ctxTimeout)
	defer cancel()
	stmt := fmt.Sprintf("update users set username=%q, password=%q where id=%q;", usr.Username, usr.Password, usr.ID)
	result, err := repo.db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.ErrNoRowsAffected
	}
	return nil
}

func (repo *UserRepoImpl) DeleteByID(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.ctxTimeout)
	defer cancel()
	stmt := fmt.Sprintf("delete from users where id=%q;", userID)
	result, err := repo.db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.ErrNoRowsAffected
	}
	return nil
}
