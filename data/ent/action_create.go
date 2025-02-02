// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/action"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/permission"
)

// ActionCreate is the builder for creating a Action entity.
type ActionCreate struct {
	config
	mutation *ActionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *ActionCreate) SetName(s string) *ActionCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *ActionCreate) SetDescription(s string) *ActionCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ActionCreate) SetNillableDescription(s *string) *ActionCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ActionCreate) SetID(u uuid.UUID) *ActionCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ActionCreate) SetNillableID(u *uuid.UUID) *ActionCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ac *ActionCreate) AddPermissionIDs(ids ...uuid.UUID) *ActionCreate {
	ac.mutation.AddPermissionIDs(ids...)
	return ac
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ac *ActionCreate) AddPermissions(p ...*Permission) *ActionCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPermissionIDs(ids...)
}

// Mutation returns the ActionMutation object of the builder.
func (ac *ActionCreate) Mutation() *ActionMutation {
	return ac.mutation
}

// Save creates the Action in the database.
func (ac *ActionCreate) Save(ctx context.Context) (*Action, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActionCreate) SaveX(ctx context.Context) *Action {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ActionCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ActionCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ActionCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := action.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActionCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Action.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := action.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Action.name": %w`, err)}
		}
	}
	return nil
}

func (ac *ActionCreate) sqlSave(ctx context.Context) (*Action, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ActionCreate) createSpec() (*Action, *sqlgraph.CreateSpec) {
	var (
		_node = &Action{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(action.Table, sqlgraph.NewFieldSpec(action.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(action.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(action.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if nodes := ac.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   action.PermissionsTable,
			Columns: []string{action.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActionCreateBulk is the builder for creating many Action entities in bulk.
type ActionCreateBulk struct {
	config
	err      error
	builders []*ActionCreate
}

// Save creates the Action entities in the database.
func (acb *ActionCreateBulk) Save(ctx context.Context) ([]*Action, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Action, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActionMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ActionCreateBulk) SaveX(ctx context.Context) []*Action {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ActionCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ActionCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
