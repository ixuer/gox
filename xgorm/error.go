package xgorm

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsUniqueConflict Determine if it is a unique key conflict
func IsUniqueConflict(err error) bool {
	if IsPostgresUniqueConflict(err) {
		return true
	}
	if IsMysqlUniqueConflict(err) {
		return true
	}
	return false
}

// IsPostgresUniqueConflict Determine if it is a unique key conflict
func IsPostgresUniqueConflict(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}

// IsMysqlUniqueConflict Determine if it is a unique key conflict
func IsMysqlUniqueConflict(err error) bool {
	var mySQLErr *mysql.MySQLError
	if errors.As(err, &mySQLErr) && mySQLErr.Number == 1062 {
		return true
	}
	return false
}
