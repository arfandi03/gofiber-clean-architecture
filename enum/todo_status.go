package enum

import "database/sql/driver"

type TodoStatus string

const (
	OPEN      TodoStatus = "OPEN"
	WIP       TodoStatus = "WIP"
	COMPLETED TodoStatus = "COMPLETED"
)

func (status *TodoStatus) Scan(value interface{}) error {
	*status = TodoStatus(value.([]byte))
	return nil
}

func (status TodoStatus) Value() (driver.Value, error) {
	return string(status), nil
}

/*
ref:
- https://stackoverflow.com/questions/68637265/how-can-i-add-enum-in-gorm
*/
