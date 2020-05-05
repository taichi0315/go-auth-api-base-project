package persistence

import (
	"database/sql"

	"github.com/nepp-tumsat/documents-api/infrastructure/model"
	"golang.org/x/xerrors"
)

type AuthRepository interface {
	InsertUser(model.User) error
	InsertUserAuth(model.UserAuth) error
	InsertAuthToken(model.AuthToken) error
	SelectUserByUserID(uint64) (*model.User, error)
	SelectUserByUserName(string) (*model.User, error)
	SelectUserAuthByEmail(string) (*model.UserAuth, error)
	SelectAuthTokenByToken(string) (*model.AuthToken, error)
	DeleteAuthTokenByToken(string) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) InsertUser(user model.User) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			users(
			  username
			)
		VALUES(?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(user.UserName)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}

func (a *authRepository) InsertUserAuth(userAuth model.UserAuth) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			user_auths(
			  user_id,
			  email,
				hash
			)
		VALUES(?,?,?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(userAuth.UserID, userAuth.Email, userAuth.Hash)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}

func (a *authRepository) InsertAuthToken(authToken model.AuthToken) error {
	stmt, err := a.db.Prepare(`
		INSERT INTO
			auth_tokens(
			  user_id,
				token
			)
		VALUES(?,?);
	`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(authToken.UserID, authToken.Token)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}

func (a *authRepository) SelectUserByUserName(userName string) (*model.User, error) {
	row := a.db.QueryRow(`
		SELECT
		  id,
			username
		FROM
		  users
		WHERE
		  username=?;
	`, userName)

	var user model.User
	err := row.Scan(&user.UserID, &user.UserName)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) SelectUserByUserID(userID uint64) (*model.User, error) {
	row := a.db.QueryRow(`
		SELECT
		  username
		FROM
		  users
		WHERE
		  id=?;
	`, userID)

	var user model.User
	err := row.Scan(&user.UserName)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) SelectUserAuthByEmail(email string) (*model.UserAuth, error) {
	row := a.db.QueryRow(`
		SELECT
		  user_id,
			email,
			hash
		FROM
		  user_auths
		WHERE
		  email=?;
	`, email)

	var userAuth model.UserAuth
	err := row.Scan(&userAuth.UserID, &userAuth.Email, &userAuth.Hash)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}
	return &userAuth, nil
}

func (a *authRepository) SelectAuthTokenByToken(token string) (*model.AuthToken, error) {
	row := a.db.QueryRow(`
		SELECT
		  user_id,
			token
		FROM
		  auth_tokens
		WHERE
		  token=?;
	`, token)

	var authToken model.AuthToken
	err := row.Scan(&authToken.UserID, &authToken.Token)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}
	return &authToken, nil
}

func (a *authRepository) DeleteAuthTokenByToken(token string) error {
	stmt, err := a.db.Prepare(`
		DELETE FROM
		  auth_tokens
		WHERE
			token=?;
		`)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}

	_, err = stmt.Exec(token)
	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return err
	}
	return nil
}
