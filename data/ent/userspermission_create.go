// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/permission"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/user"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/userspermission"
)

// UsersPermissionCreate is the builder for creating a UsersPermission entity.
type UsersPermissionCreate struct {
	config
	mutation *UsersPermissionMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (upc *UsersPermissionCreate) SetUserID(u uuid.UUID) *UsersPermissionCreate {
	upc.mutation.SetUserID(u)
	return upc
}

// SetPermissionID sets the "permission_id" field.
func (upc *UsersPermissionCreate) SetPermissionID(u uuid.UUID) *UsersPermissionCreate {
	upc.mutation.SetPermissionID(u)
	return upc
}

// SetID sets the "id" field.
func (upc *UsersPermissionCreate) SetID(u uuid.UUID) *UsersPermissionCreate {
	upc.mutation.SetID(u)
	return upc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (upc *UsersPermissionCreate) SetNillableID(u *uuid.UUID) *UsersPermissionCreate {
	if u != nil {
		upc.SetID(*u)
	}
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UsersPermissionCreate) SetUser(u *User) *UsersPermissionCreate {
	return upc.SetUserID(u.ID)
}

// SetPermission sets the "permission" edge to the Permission entity.
func (upc *UsersPermissionCreate) SetPermission(p *Permission) *UsersPermissionCreate {
	return upc.SetPermissionID(p.ID)
}

// Mutation returns the UsersPermissionMutation object of the builder.
func (upc *UsersPermissionCreate) Mutation() *UsersPermissionMutation {
	return upc.mutation
}

// Save creates the UsersPermission in the database.
func (upc *UsersPermissionCreate) Save(ctx context.Context) (*UsersPermission, error) {
	upc.defaults()
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UsersPermissionCreate) SaveX(ctx context.Context) *UsersPermission {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UsersPermissionCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UsersPermissionCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UsersPermissionCreate) defaults() {
	if _, ok := upc.mutation.ID(); !ok {
		v := userspermission.DefaultID()
		upc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UsersPermissionCreate) check() error {
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UsersPermission.user_id"`)}
	}
	if _, ok := upc.mutation.PermissionID(); !ok {
		return &ValidationError{Name: "permission_id", err: errors.New(`ent: missing required field "UsersPermission.permission_id"`)}
	}
	if len(upc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UsersPermission.user"`)}
	}
	if len(upc.mutation.PermissionIDs()) == 0 {
		return &ValidationError{Name: "permission", err: errors.New(`ent: missing required edge "UsersPermission.permission"`)}
	}
	return nil
}

func (upc *UsersPermissionCreate) sqlSave(ctx context.Context) (*UsersPermission, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UsersPermissionCreate) createSpec() (*UsersPermission, *sqlgraph.CreateSpec) {
	var (
		_node = &UsersPermission{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userspermission.Table, sqlgraph.NewFieldSpec(userspermission.FieldID, field.TypeUUID))
	)
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userspermission.UserTable,
			Columns: []string{userspermission.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := upc.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userspermission.PermissionTable,
			Columns: []string{userspermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PermissionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UsersPermissionCreateBulk is the builder for creating many UsersPermission entities in bulk.
type UsersPermissionCreateBulk struct {
	config
	err      error
	builders []*UsersPermissionCreate
}

// Save creates the UsersPermission entities in the database.
func (upcb *UsersPermissionCreateBulk) Save(ctx context.Context) ([]*UsersPermission, error) {
	if upcb.err != nil {
		return nil, upcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UsersPermission, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersPermissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UsersPermissionCreateBulk) SaveX(ctx context.Context) []*UsersPermission {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UsersPermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UsersPermissionCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
