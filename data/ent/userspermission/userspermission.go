// Code generated by ent, DO NOT EDIT.

package userspermission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userspermission type in the database.
	Label = "users_permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPermissionID holds the string denoting the permission_id field in the database.
	FieldPermissionID = "permission_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePermission holds the string denoting the permission edge name in mutations.
	EdgePermission = "permission"
	// Table holds the table name of the userspermission in the database.
	Table = "users_permission"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_users_permissions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// PermissionTable is the table that holds the permission relation/edge. The primary key declared below.
	PermissionTable = "permission_users_permissions"
	// PermissionInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionInverseTable = "permission"
)

// Columns holds all SQL columns for userspermission fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldPermissionID,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "users_permission_id"}
	// PermissionPrimaryKey and PermissionColumn2 are the table columns denoting the
	// primary key for the permission relation (M2M).
	PermissionPrimaryKey = []string{"permission_id", "users_permission_id"}
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the UsersPermission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPermissionID orders the results by the permission_id field.
func ByPermissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermissionID, opts...).ToFunc()
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPermissionCount orders the results by permission count.
func ByPermissionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionStep(), opts...)
	}
}

// ByPermission orders the results by permission terms.
func ByPermission(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
func newPermissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PermissionTable, PermissionPrimaryKey...),
	)
}
