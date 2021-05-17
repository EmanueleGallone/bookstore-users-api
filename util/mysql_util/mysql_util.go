package mysql_util

import (
	"bookstore-users-api/util/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	noRowFound         = "no rows"
	duplicatedKeyError = 1062
	incorrectDateTime  = 1292
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError) // parsing to mysql error
	if !ok {
		if strings.Contains(err.Error(), noRowFound) {
			return errors.NewBadRequestError("no ID matching found")
		}
		return errors.NewInternalServerError("Unable to parse database error")
	}
	switch sqlErr.Number {
	case duplicatedKeyError:
		return errors.NewBadRequestError("Invalid data")
	case incorrectDateTime:
		return errors.NewBadRequestError("Invalid datetime")
	}

	return errors.NewInternalServerError("Error processing request")
}
