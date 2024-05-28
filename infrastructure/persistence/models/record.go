package models

import (
	"fmt"
	"go-complaint/erros"
	"strings"
)

type Columns []string

type Values []interface{}

type Record interface {
	Columns() Columns
	Values() Values
	Table() string
	Args() string
}

func Alias(columns Columns, alias string) Columns {
	newColumns := make(Columns, 0)
	for _, c := range columns {
		newColumns = append(newColumns, alias+"."+c)
	}
	return newColumns
}

func StringKeyArgs(columns Columns) string {
	var args string
	for i, c := range columns {
		args += c + " = $" + fmt.Sprint(i+1)
		if i != len(columns)-1 {
			args += ", "
		}
	}
	return args

}

func StringColumns(columns Columns) string {
	return strings.Join(columns, ", ")
}

func FindColumn(columns Columns, column ...string) (Columns, error) {
	newColumns := make(Columns, 0)
	for _, c := range column {
		for _, v := range columns {
			if c == v {
				newColumns = append(newColumns, c)
			}
		}
	}
	if len(newColumns) == 0 {
		return nil, &erros.ColumnNotFoundError{}
	}
	return newColumns, nil
}

func Args(columns Columns) string {
	var args string
	for i := range columns {
		args += "$" + fmt.Sprint(i+1)
		if i != len(columns)-1 {
			args += ","
		}
	}
	return args
}
