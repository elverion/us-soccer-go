// Code generated by ent, DO NOT EDIT.

package weather

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the weather type in the database.
	Label = "weather"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTemperature holds the string denoting the temperature field in the database.
	FieldTemperature = "temperature"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeStadium holds the string denoting the stadium edge name in mutations.
	EdgeStadium = "stadium"
	// Table holds the table name of the weather in the database.
	Table = "weathers"
	// StadiumTable is the table that holds the stadium relation/edge.
	StadiumTable = "weathers"
	// StadiumInverseTable is the table name for the Stadium entity.
	// It exists in this package in order to avoid circular dependency with the "stadium" package.
	StadiumInverseTable = "stadia"
	// StadiumColumn is the table column denoting the stadium relation/edge.
	StadiumColumn = "stadium_weather"
)

// Columns holds all SQL columns for weather fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTemperature,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "weathers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"stadium_weather",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "us-soccer-go-test/internal/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Weather queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByTemperature orders the results by the temperature field.
func ByTemperature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemperature, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStadiumField orders the results by stadium field.
func ByStadiumField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStadiumStep(), sql.OrderByField(field, opts...))
	}
}
func newStadiumStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StadiumInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, StadiumTable, StadiumColumn),
	)
}
