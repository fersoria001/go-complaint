package persistence

import (
	"reflect"
)

//A way to query using pgx preserving self encapsulation of the domain models

type COMMAND int

const (
	SELECT COMMAND = iota
	INSERT
	UPDATE
	DELETE
)

// ConstructQuery is a function that constructs a query based on the command and the struct passed
// The struct passed is the domain model that will be used to construct the query
// The command is the type of query that will be constructed
// The table name will be arbitratry and will be the same as the struct name lowercased in plural
// The function uses reflection to get the fields of the struct and construct the query
// The function returns a string with the query constructed
func ConstructQuery(c COMMAND, s interface{}) {
	interfaceType := reflect.TypeOf(s)
	if interfaceType.Kind() != reflect.Struct {
		return
	}
	switch c {
	case SELECT:
		//Construct a select query
	case INSERT:
		//Construct an insert query
	case UPDATE:
		//Construct an update query
	case DELETE:
		//Construct a delete query
	}
}

// Fields is a function that gets the fields of a struct
// And returns them in a db friendly format(snake_case)
// idea 1 convert to snake case
// idea 2 get the tag of the field
// can I get private the value trough reflection?
func Fields(structType reflect.Type, reflected map[string]interface{}) {
	//to make it better we need to put a condition
	//to finish a recursion, what can it be?
	//we need to know if its already traversed
	//we need to know if something has left to traverse
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == "Entity" && field.Type == reflect.TypeOf(COMMAND(0)) {
			reflected["table_name"] = field.Tag.Get("table_name")
		}
		if field.Type.Kind() == reflect.Struct {
			Fields(field.Type, reflected)
		}
		reflected[ToSnakeCase(field.Name)] = field.Type
	}
}

func ToSnakeCase(s string) string {
	var snake string
	for i, c := range s {
		if i > 0 && c >= 'A' && c <= 'Z' {
			snake += "_"
		}
		snake += string(c)
	}
	return snake
}

// ReturnType should return a new struct that reflects the value types from the database
// So that you can use the values from the database to construct the domain model
// trough a constructor/factory method
func ReturnType() {}
