package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
	id integer primary key not null auto_increment,
	title varchar(20) not null,
	description varchar(250) not null,
	status bool not null
	);`
)
