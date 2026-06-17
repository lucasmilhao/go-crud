package service

import (
	"database/sql"
	"fmt"
	"strings"
)

func FindAll(db *sql.DB, tableName string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
}

func FindById(db *sql.DB, tableName string, id any) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName), id)
}

func Save(db *sql.DB, tableName string, rows []string, values []any) (sql.Result, error) {
	params := strings.Join(rows, ", ")
	var question string
	for i, _ := range rows {
		if i == len(rows)-1 {
			question += "?"
			break
		}
		question += "?, "
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) values (%s)", tableName, params, question)
	return db.Exec(query, values...)
}

func DeleteById(db *sql.DB, tableName string, id any) (sql.Result, error) {
	return db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName), id)
}

func UpdateById(db *sql.DB, tableName string, id any, fields []string, params []any) (sql.Result, error) {

	var query string

	for _, field := range fields {
		query += field + " = ? "
	}
	fmt.Printf("UPDATE %s SET %s WHERE id = ?", tableName, query)
	return db.Exec(fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", tableName, query), id)

}
