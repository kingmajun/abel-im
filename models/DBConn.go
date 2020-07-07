package models

import (
	"database/sql"
	"fmt"
)

var (
	db  *sql.DB
	err error
)
//cwei@~^ha3
func init() {
	db, err = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/im?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		panic(err.Error())
	}
}

type DBConn struct {
	db *sql.DB
}


// 获取一行记录
func (this *DBConn) GetOne(sql string, args ...interface{}) (dbRow, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	result, err := scanRow(rows)
	return result, err
}

// 获取多行记录
func (this *DBConn) GetAll(sql string, args ...interface{}) ([]dbRow, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]dbRow, 0)
	for rows.Next() {
		r, err := scanRow(rows)
		if err != nil {
			continue
		}
		result = append(result, r)
	}
	return result, nil

}