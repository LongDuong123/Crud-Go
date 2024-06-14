package helpers

import "database/sql"

func NullInt64ToInt(nullInt sql.NullInt64) int {
	if nullInt.Valid {
		return int(nullInt.Int64)
	}
	return -1
}

func NullStringToString(nullString sql.NullString) string {
	if nullString.Valid {
		return string(nullString.String)
	}
	return ""
}

func IntToInt64Null(IntCanNull int64) sql.NullInt64 {
	if IntCanNull == 0 {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: IntCanNull, Valid: true}
}

func StringToNullString(StringCanNull string) sql.NullString {
	if StringCanNull == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: StringCanNull, Valid: true}
}
