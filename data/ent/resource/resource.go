// Code generated by ent, DO NOT EDIT.

package resource

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the resource type in the database.
	Label = "resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// Table holds the table name of the resource in the database.
	Table = "resource"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "resource_permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permission"
)

// Columns holds all SQL columns for resource fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

var (
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"resource_id", "permission_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Resource queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
	)
}