// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/permission"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/predicate"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/resource"
)

// ResourceQuery is the builder for querying Resource entities.
type ResourceQuery struct {
	config
	ctx             *QueryContext
	order           []resource.OrderOption
	inters          []Interceptor
	predicates      []predicate.Resource
	withPermissions *PermissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResourceQuery builder.
func (rq *ResourceQuery) Where(ps ...predicate.Resource) *ResourceQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ResourceQuery) Limit(limit int) *ResourceQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ResourceQuery) Offset(offset int) *ResourceQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ResourceQuery) Unique(unique bool) *ResourceQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ResourceQuery) Order(o ...resource.OrderOption) *ResourceQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryPermissions chains the current query on the "permissions" edge.
func (rq *ResourceQuery) QueryPermissions() *PermissionQuery {
	query := (&PermissionClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, resource.PermissionsTable, resource.PermissionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Resource entity from the query.
// Returns a *NotFoundError when no Resource was found.
func (rq *ResourceQuery) First(ctx context.Context) (*Resource, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ResourceQuery) FirstX(ctx context.Context) *Resource {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Resource ID from the query.
// Returns a *NotFoundError when no Resource ID was found.
func (rq *ResourceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ResourceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Resource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Resource entity is found.
// Returns a *NotFoundError when no Resource entities are found.
func (rq *ResourceQuery) Only(ctx context.Context) (*Resource, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resource.Label}
	default:
		return nil, &NotSingularError{resource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ResourceQuery) OnlyX(ctx context.Context) *Resource {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Resource ID in the query.
// Returns a *NotSingularError when more than one Resource ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ResourceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = &NotSingularError{resource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ResourceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Resources.
func (rq *ResourceQuery) All(ctx context.Context) ([]*Resource, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryAll)
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Resource, *ResourceQuery]()
	return withInterceptors[[]*Resource](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ResourceQuery) AllX(ctx context.Context) []*Resource {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Resource IDs.
func (rq *ResourceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryIDs)
	if err = rq.Select(resource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ResourceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ResourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryCount)
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ResourceQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ResourceQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ResourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryExist)
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ResourceQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ResourceQuery) Clone() *ResourceQuery {
	if rq == nil {
		return nil
	}
	return &ResourceQuery{
		config:          rq.config,
		ctx:             rq.ctx.Clone(),
		order:           append([]resource.OrderOption{}, rq.order...),
		inters:          append([]Interceptor{}, rq.inters...),
		predicates:      append([]predicate.Resource{}, rq.predicates...),
		withPermissions: rq.withPermissions.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithPermissions tells the query-builder to eager-load the nodes that are connected to
// the "permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ResourceQuery) WithPermissions(opts ...func(*PermissionQuery)) *ResourceQuery {
	query := (&PermissionClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withPermissions = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Resource.Query().
//		GroupBy(resource.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ResourceQuery) GroupBy(field string, fields ...string) *ResourceGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ResourceGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = resource.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Resource.Query().
//		Select(resource.FieldName).
//		Scan(ctx, &v)
func (rq *ResourceQuery) Select(fields ...string) *ResourceSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ResourceSelect{ResourceQuery: rq}
	sbuild.label = resource.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ResourceSelect configured with the given aggregations.
func (rq *ResourceQuery) Aggregate(fns ...AggregateFunc) *ResourceSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ResourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !resource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ResourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Resource, error) {
	var (
		nodes       = []*Resource{}
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withPermissions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Resource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Resource{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withPermissions; query != nil {
		if err := rq.loadPermissions(ctx, query, nodes,
			func(n *Resource) { n.Edges.Permissions = []*Permission{} },
			func(n *Resource, e *Permission) { n.Edges.Permissions = append(n.Edges.Permissions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ResourceQuery) loadPermissions(ctx context.Context, query *PermissionQuery, nodes []*Resource, init func(*Resource), assign func(*Resource, *Permission)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Resource)
	nids := make(map[uuid.UUID]map[*Resource]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(resource.PermissionsTable)
		s.Join(joinT).On(s.C(permission.FieldID), joinT.C(resource.PermissionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(resource.PermissionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(resource.PermissionsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Resource]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Permission](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "permissions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rq *ResourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ResourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for i := range fields {
			if fields[i] != resource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ResourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(resource.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = resource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResourceGroupBy is the group-by builder for Resource entities.
type ResourceGroupBy struct {
	selector
	build *ResourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ResourceGroupBy) Aggregate(fns ...AggregateFunc) *ResourceGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ResourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceQuery, *ResourceGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ResourceGroupBy) sqlScan(ctx context.Context, root *ResourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ResourceSelect is the builder for selecting fields of Resource entities.
type ResourceSelect struct {
	*ResourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ResourceSelect) Aggregate(fns ...AggregateFunc) *ResourceSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ResourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceQuery, *ResourceSelect](ctx, rs.ResourceQuery, rs, rs.inters, v)
}

func (rs *ResourceSelect) sqlScan(ctx context.Context, root *ResourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
