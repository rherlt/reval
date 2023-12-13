// Code generated by ent, DO NOT EDIT.

package configuration

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the configuration type in the database.
	Label = "configuration"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the configuration in the database.
	Table = "configurations"
)

// Columns holds all SQL columns for configuration fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Configuration queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}