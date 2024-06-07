package errors

import (
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func GormErrHandler(err error) error {
	if Equal(err, gorm.ErrRecordNotFound) {
		return NewNotFoundError("record not found")
	} else {
		switch v := err.(type) {
		case *pgconn.PgError:
			return postgresErrorMap[v.Code]
		default:
			return err
		}
	}
}
