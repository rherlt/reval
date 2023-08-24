// Code generated by ent, DO NOT EDIT.

package scenario

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the scenario type in the database.
	Label = "scenario"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldExternalId holds the string denoting the externalid field in the database.
	FieldExternalId = "external_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSystemprompt holds the string denoting the systemprompt field in the database.
	FieldSystemprompt = "systemprompt"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// EdgeResponses holds the string denoting the responses edge name in mutations.
	EdgeResponses = "responses"
	// Table holds the table name of the scenario in the database.
	Table = "scenarios"
	// ResponsesTable is the table that holds the responses relation/edge.
	ResponsesTable = "responses"
	// ResponsesInverseTable is the table name for the Response entity.
	// It exists in this package in order to avoid circular dependency with the "response" package.
	ResponsesInverseTable = "responses"
	// ResponsesColumn is the table column denoting the responses relation/edge.
	ResponsesColumn = "scenario_id"
)

// Columns holds all SQL columns for scenario fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldExternalId,
	FieldDescription,
	FieldSystemprompt,
	FieldDate,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Scenario queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByExternalId orders the results by the externalId field.
func ByExternalId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalId, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySystemprompt orders the results by the systemprompt field.
func BySystemprompt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemprompt, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByResponsesCount orders the results by responses count.
func ByResponsesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResponsesStep(), opts...)
	}
}

// ByResponses orders the results by responses terms.
func ByResponses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResponsesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newResponsesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResponsesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResponsesTable, ResponsesColumn),
	)
}
