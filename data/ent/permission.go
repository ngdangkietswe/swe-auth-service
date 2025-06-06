// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/action"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/permission"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/resource"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ActionID holds the value of the "action_id" field.
	ActionID uuid.UUID `json:"action_id,omitempty"`
	// ResourceID holds the value of the "resource_id" field.
	ResourceID uuid.UUID `json:"resource_id,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Action holds the value of the action edge.
	Action *Action `json:"action,omitempty"`
	// Resource holds the value of the resource edge.
	Resource *Resource `json:"resource,omitempty"`
	// UsersPermissions holds the value of the users_permissions edge.
	UsersPermissions []*UsersPermission `json:"users_permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ActionOrErr returns the Action value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) ActionOrErr() (*Action, error) {
	if e.Action != nil {
		return e.Action, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: action.Label}
	}
	return nil, &NotLoadedError{edge: "action"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) ResourceOrErr() (*Resource, error) {
	if e.Resource != nil {
		return e.Resource, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: resource.Label}
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// UsersPermissionsOrErr returns the UsersPermissions value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) UsersPermissionsOrErr() ([]*UsersPermission, error) {
	if e.loadedTypes[2] {
		return e.UsersPermissions, nil
	}
	return nil, &NotLoadedError{edge: "users_permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldDescription:
			values[i] = new(sql.NullString)
		case permission.FieldID, permission.FieldActionID, permission.FieldResourceID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case permission.FieldActionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field action_id", values[i])
			} else if value != nil {
				pe.ActionID = *value
			}
		case permission.FieldResourceID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value != nil {
				pe.ResourceID = *value
			}
		case permission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = new(string)
				*pe.Description = value.String
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryAction queries the "action" edge of the Permission entity.
func (pe *Permission) QueryAction() *ActionQuery {
	return NewPermissionClient(pe.config).QueryAction(pe)
}

// QueryResource queries the "resource" edge of the Permission entity.
func (pe *Permission) QueryResource() *ResourceQuery {
	return NewPermissionClient(pe.config).QueryResource(pe)
}

// QueryUsersPermissions queries the "users_permissions" edge of the Permission entity.
func (pe *Permission) QueryUsersPermissions() *UsersPermissionQuery {
	return NewPermissionClient(pe.config).QueryUsersPermissions(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("action_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.ActionID))
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.ResourceID))
	builder.WriteString(", ")
	if v := pe.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
