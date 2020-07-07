package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type dbRow map[string]interface{}



func scanRow(rows *sql.Rows) (dbRow, error) {
	columns, _ := rows.Columns()

	vals := make([]interface{}, len(columns))
	valsPtr := make([]interface{}, len(columns))

	for i := range vals {
		valsPtr[i] = &vals[i]
	}

	err := rows.Scan(valsPtr...)

	if err != nil {
		return nil ,err
	}

	r := make(dbRow)

	for i, v := range columns {
		if va, ok := vals[i].([]byte); ok {
			r[v] = string(va)
		} else {
			r[v] = vals[i]
		}
	}

	return r, nil

}
