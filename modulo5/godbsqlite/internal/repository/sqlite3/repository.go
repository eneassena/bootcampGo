package sqlite3

import (
	"database/sql"
	"go-db-sqlite/internal/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertData(data domain.User) error {
	sqlStmt := `
	insert into users(userid, username, email) values(?, ?, ?);
	`
	stmt, err := r.db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.UserID, data.UserName, data.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectOneData(user *domain.User) (*domain.User, error) {
	var dataResult *domain.User

	rows := r.db.QueryRow("SELECT userid, username,email FROM users WHERE userid=?;", user.UserID)
	if err := rows.Scan(dataResult.UserID, dataResult.UserName, dataResult.Email); err != nil {
		return &domain.User{}, err
	}

	return &domain.User{}, nil
}

func (r *repository) SelectData() ([]*domain.User, error) {
	rows, err := r.db.Query("SELECT userid, username,email FROM users;")
	if err != nil {
		return []*domain.User{}, err
	}
	defer rows.Close()

	var listUser []*domain.User

	for rows.Next() {
		var user domain.User

		if err := rows.Scan(&user.UserID, &user.UserName, &user.Email); err != nil {
			return []*domain.User{}, err
		}
		listUser = append(listUser, &user)
	}
	return listUser, nil
}

func (r *repository) UpdateData(oldUser, newUser domain.User) error {
	query := "SELECT userid,username,email FROM users WHERE userid=?;"
	rows := r.db.QueryRow(query, oldUser.UserID)
	var userObject domain.User

	if rows.Err() != nil {
		return rows.Err()
	}
	if err := rows.Scan(&userObject.UserID, &userObject.UserName, &userObject.Email); err != nil {
		return err
	}

	validateDataUpdateUser(&oldUser, &newUser)

	query = `
	UPDATE users SET username=?, email=? WHERE userid=?;
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&userObject.UserName, userObject.Email, &userObject.UserID)
	if err != nil {
		return err
	}
	return nil
}

func validateDataUpdateUser(oldUser, newUser *domain.User) {
	if newUser.UserName != "" {
		oldUser.UserName = newUser.UserName
	}
	if newUser.Email != "" {
		oldUser.Email = newUser.Email
	}
}
