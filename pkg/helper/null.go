package helper

import (
	"database/sql"

	"github.com/google/uuid"
)

func UUIDToNullString(id uuid.UUID) (ns sql.NullString) {
	if id != uuid.Nil {
		ns.String = id.String()
		ns.Valid = true
	}
	return ns
}

// EmptyStringToNull also handles the case when user explicitly wants to delete photo
func EmptyStringToNull(s string) (res sql.NullString) {
	if s == "" {
		res.Valid = false
		return res
	}

	res.String = s
	res.Valid = true
	return res
}

// Float64ToNullFloat64 ...
func Float64ToNullFloat64(f float64) (nf sql.NullFloat64) {
	if f != 0 {
		nf.Float64 = f
		nf.Valid = true
	}
	return nf
}

func Int64ToNullInt64(number int64) (ni sql.NullInt64) {
	if number != 0 {
		ni.Int64 = number
		ni.Valid = true
	}
	return ni
}

// StringToNullString ...
func StringToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}
