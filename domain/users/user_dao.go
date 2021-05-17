/*
Data Access Object.
Logic related to persistent data (DB)
*/
package users

import (
	"bookstore-users-api/datasources/mysql_userDB"
	"bookstore-users-api/logger"
	"bookstore-users-api/util/errors"
	"bookstore-users-api/util/mysql_util"
	"fmt"
)

var (
	db = mysql_userDB.DBClient
)

const (
	selectByIDQuery        = "SELECT * FROM users_db.user WHERE ID = ?"
	insertQuery            = "INSERT INTO users_db.user (FirstName, LastName, Email, DateCreated, Status, Password) VALUES (?,?,?,?,?,?) "
	updateQuery            = "UPDATE users_db.user SET FirstName=?, LastName=?, Email=? WHERE users_db.user.ID=?"
	setNotActiveQuery      = "UPDATE users_db.user SET Status = 'Inactive' WHERE users_db.user.ID=?"
	findUsersByStatusQuery = "SELECT FirstName, LastName, Email, DateCreated, Status FROM users_db.user WHERE Status = ?"
)

/*
retrieves user data from database.
*/
func (user *User) GetByID() *errors.RestError {
	stmt, err := db.Prepare(selectByIDQuery)
	if err != nil {
		logger.Error("Error when trying to GetByID", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
		&user.Password,
	)

	if err != nil {
		logger.Error("Error when trying to GetByID", err)
		return mysql_util.ParseError(err)
	}
	return nil
}

/*
Method to save object to DB
*/
func (user *User) SaveToDB() *errors.RestError {
	if err := db.Ping(); err != nil {
		panic("Connection to DB lost!")
	}

	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		logger.Error("Error when trying to SaveToDB", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)

	if err != nil {
		logger.Error("Error when trying to GetByID", err)
		return mysql_util.ParseError(err)
	}

	user.Id, err = insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to GetByID", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) UpdateUser() *errors.RestError {
	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		logger.Error("Error when trying to GetByID", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		return mysql_util.ParseError(err)
	}

	return nil
}

func (user *User) DeleteUser() *errors.RestError {
	stmt, err := db.Prepare(setNotActiveQuery)
	if err != nil {
		return mysql_util.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_util.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := db.Prepare(findUsersByStatusQuery)
	if err != nil {
		logger.Error("Error when trying to find by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to stmt.Query FindByStatus", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_util.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}

	return results, nil
}
