// Code generated by ent, DO NOT EDIT.

package permission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActionID holds the string denoting the action_id field in the database.
	FieldActionID = "action_id"
	// FieldResourceID holds the string denoting the resource_id field in the database.
	FieldResourceID = "resource_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeAction holds the string denoting the action edge name in mutations.
	EdgeAction = "action"
	// EdgeResource holds the string denoting the resource edge name in mutations.
	EdgeResource = "resource"
	// EdgeUsersPermissions holds the string denoting the users_permissions edge name in mutations.
	EdgeUsersPermissions = "users_permissions"
	// Table holds the table name of the permission in the database.
	Table = "permission"
	// ActionTable is the table that holds the action relation/edge.
	ActionTable = "permission"
	// ActionInverseTable is the table name for the Action entity.
	// It exists in this package in order to avoid circular dependency with the "action" package.
	ActionInverseTable = "action"
	// ActionColumn is the table column denoting the action relation/edge.
	ActionColumn = "action_id"
	// ResourceTable is the table that holds the resource relation/edge.
	ResourceTable = "permission"
	// ResourceInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourceInverseTable = "resource"
	// ResourceColumn is the table column denoting the resource relation/edge.
	ResourceColumn = "resource_id"
	// UsersPermissionsTable is the table that holds the users_permissions relation/edge.
	UsersPermissionsTable = "users_permission"
	// UsersPermissionsInverseTable is the table name for the UsersPermission entity.
	// It exists in this package in order to avoid circular dependency with the "userspermission" package.
	UsersPermissionsInverseTable = "users_permission"
	// UsersPermissionsColumn is the table column denoting the users_permissions relation/edge.
	UsersPermissionsColumn = "permission_id"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldActionID,
	FieldResourceID,
	FieldDescription,
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

// OrderOption defines the ordering options for the Permission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByActionID orders the results by the action_id field.
func ByActionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionID, opts...).ToFunc()
}

// ByResourceID orders the results by the resource_id field.
func ByResourceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResourceID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByActionField orders the results by action field.
func ByActionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActionStep(), sql.OrderByField(field, opts...))
	}
}

// ByResourceField orders the results by resource field.
func ByResourceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourceStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersPermissionsCount orders the results by users_permissions count.
func ByUsersPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersPermissionsStep(), opts...)
	}
}

// ByUsersPermissions orders the results by users_permissions terms.
func ByUsersPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActionTable, ActionColumn),
	)
}
func newResourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ResourceTable, ResourceColumn),
	)
}
func newUsersPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersPermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersPermissionsTable, UsersPermissionsColumn),
	)
}
