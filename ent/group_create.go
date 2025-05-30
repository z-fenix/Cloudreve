// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudreve/Cloudreve/v4/ent/group"
	"github.com/cloudreve/Cloudreve/v4/ent/storagepolicy"
	"github.com/cloudreve/Cloudreve/v4/ent/user"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
	"github.com/cloudreve/Cloudreve/v4/pkg/boolset"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GroupCreate) SetDeletedAt(t time.Time) *GroupCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetMaxStorage sets the "max_storage" field.
func (gc *GroupCreate) SetMaxStorage(i int64) *GroupCreate {
	gc.mutation.SetMaxStorage(i)
	return gc
}

// SetNillableMaxStorage sets the "max_storage" field if the given value is not nil.
func (gc *GroupCreate) SetNillableMaxStorage(i *int64) *GroupCreate {
	if i != nil {
		gc.SetMaxStorage(*i)
	}
	return gc
}

// SetSpeedLimit sets the "speed_limit" field.
func (gc *GroupCreate) SetSpeedLimit(i int) *GroupCreate {
	gc.mutation.SetSpeedLimit(i)
	return gc
}

// SetNillableSpeedLimit sets the "speed_limit" field if the given value is not nil.
func (gc *GroupCreate) SetNillableSpeedLimit(i *int) *GroupCreate {
	if i != nil {
		gc.SetSpeedLimit(*i)
	}
	return gc
}

// SetPermissions sets the "permissions" field.
func (gc *GroupCreate) SetPermissions(bs *boolset.BooleanSet) *GroupCreate {
	gc.mutation.SetPermissions(bs)
	return gc
}

// SetSettings sets the "settings" field.
func (gc *GroupCreate) SetSettings(ts *types.GroupSetting) *GroupCreate {
	gc.mutation.SetSettings(ts)
	return gc
}

// SetStoragePolicyID sets the "storage_policy_id" field.
func (gc *GroupCreate) SetStoragePolicyID(i int) *GroupCreate {
	gc.mutation.SetStoragePolicyID(i)
	return gc
}

// SetNillableStoragePolicyID sets the "storage_policy_id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableStoragePolicyID(i *int) *GroupCreate {
	if i != nil {
		gc.SetStoragePolicyID(*i)
	}
	return gc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gc *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the "users" edges to the User entity.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// SetStoragePoliciesID sets the "storage_policies" edge to the StoragePolicy entity by ID.
func (gc *GroupCreate) SetStoragePoliciesID(id int) *GroupCreate {
	gc.mutation.SetStoragePoliciesID(id)
	return gc
}

// SetNillableStoragePoliciesID sets the "storage_policies" edge to the StoragePolicy entity by ID if the given value is not nil.
func (gc *GroupCreate) SetNillableStoragePoliciesID(id *int) *GroupCreate {
	if id != nil {
		gc = gc.SetStoragePoliciesID(*id)
	}
	return gc
}

// SetStoragePolicies sets the "storage_policies" edge to the StoragePolicy entity.
func (gc *GroupCreate) SetStoragePolicies(s *StoragePolicy) *GroupCreate {
	return gc.SetStoragePoliciesID(s.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if err := gc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		if group.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		if group.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := group.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.Settings(); !ok {
		v := group.DefaultSettings
		gc.mutation.SetSettings(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Group.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Group.updated_at"`)}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if _, ok := gc.mutation.Permissions(); !ok {
		return &ValidationError{Name: "permissions", err: errors.New(`ent: missing required field "Group.permissions"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	)

	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		id64 := int64(id)
		_spec.ID.Value = id64
	}

	_spec.OnConflict = gc.conflict
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(group.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.MaxStorage(); ok {
		_spec.SetField(group.FieldMaxStorage, field.TypeInt64, value)
		_node.MaxStorage = value
	}
	if value, ok := gc.mutation.SpeedLimit(); ok {
		_spec.SetField(group.FieldSpeedLimit, field.TypeInt, value)
		_node.SpeedLimit = value
	}
	if value, ok := gc.mutation.Permissions(); ok {
		_spec.SetField(group.FieldPermissions, field.TypeBytes, value)
		_node.Permissions = value
	}
	if value, ok := gc.mutation.Settings(); ok {
		_spec.SetField(group.FieldSettings, field.TypeJSON, value)
		_node.Settings = value
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.StoragePoliciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.StoragePoliciesTable,
			Columns: []string{group.StoragePoliciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StoragePolicyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gc *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	gc.conflict = opts
	return &GroupUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: gc,
	}
}

type (
	// GroupUpsertOne is the builder for "upsert"-ing
	//  one Group node.
	GroupUpsertOne struct {
		create *GroupCreate
	}

	// GroupUpsert is the "OnConflict" setter.
	GroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *GroupUpsert) SetUpdatedAt(v time.Time) *GroupUpsert {
	u.Set(group.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GroupUpsert) UpdateUpdatedAt() *GroupUpsert {
	u.SetExcluded(group.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GroupUpsert) SetDeletedAt(v time.Time) *GroupUpsert {
	u.Set(group.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GroupUpsert) UpdateDeletedAt() *GroupUpsert {
	u.SetExcluded(group.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GroupUpsert) ClearDeletedAt() *GroupUpsert {
	u.SetNull(group.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsert) SetName(v string) *GroupUpsert {
	u.Set(group.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsert) UpdateName() *GroupUpsert {
	u.SetExcluded(group.FieldName)
	return u
}

// SetMaxStorage sets the "max_storage" field.
func (u *GroupUpsert) SetMaxStorage(v int64) *GroupUpsert {
	u.Set(group.FieldMaxStorage, v)
	return u
}

// UpdateMaxStorage sets the "max_storage" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMaxStorage() *GroupUpsert {
	u.SetExcluded(group.FieldMaxStorage)
	return u
}

// AddMaxStorage adds v to the "max_storage" field.
func (u *GroupUpsert) AddMaxStorage(v int64) *GroupUpsert {
	u.Add(group.FieldMaxStorage, v)
	return u
}

// ClearMaxStorage clears the value of the "max_storage" field.
func (u *GroupUpsert) ClearMaxStorage() *GroupUpsert {
	u.SetNull(group.FieldMaxStorage)
	return u
}

// SetSpeedLimit sets the "speed_limit" field.
func (u *GroupUpsert) SetSpeedLimit(v int) *GroupUpsert {
	u.Set(group.FieldSpeedLimit, v)
	return u
}

// UpdateSpeedLimit sets the "speed_limit" field to the value that was provided on create.
func (u *GroupUpsert) UpdateSpeedLimit() *GroupUpsert {
	u.SetExcluded(group.FieldSpeedLimit)
	return u
}

// AddSpeedLimit adds v to the "speed_limit" field.
func (u *GroupUpsert) AddSpeedLimit(v int) *GroupUpsert {
	u.Add(group.FieldSpeedLimit, v)
	return u
}

// ClearSpeedLimit clears the value of the "speed_limit" field.
func (u *GroupUpsert) ClearSpeedLimit() *GroupUpsert {
	u.SetNull(group.FieldSpeedLimit)
	return u
}

// SetPermissions sets the "permissions" field.
func (u *GroupUpsert) SetPermissions(v *boolset.BooleanSet) *GroupUpsert {
	u.Set(group.FieldPermissions, v)
	return u
}

// UpdatePermissions sets the "permissions" field to the value that was provided on create.
func (u *GroupUpsert) UpdatePermissions() *GroupUpsert {
	u.SetExcluded(group.FieldPermissions)
	return u
}

// SetSettings sets the "settings" field.
func (u *GroupUpsert) SetSettings(v *types.GroupSetting) *GroupUpsert {
	u.Set(group.FieldSettings, v)
	return u
}

// UpdateSettings sets the "settings" field to the value that was provided on create.
func (u *GroupUpsert) UpdateSettings() *GroupUpsert {
	u.SetExcluded(group.FieldSettings)
	return u
}

// ClearSettings clears the value of the "settings" field.
func (u *GroupUpsert) ClearSettings() *GroupUpsert {
	u.SetNull(group.FieldSettings)
	return u
}

// SetStoragePolicyID sets the "storage_policy_id" field.
func (u *GroupUpsert) SetStoragePolicyID(v int) *GroupUpsert {
	u.Set(group.FieldStoragePolicyID, v)
	return u
}

// UpdateStoragePolicyID sets the "storage_policy_id" field to the value that was provided on create.
func (u *GroupUpsert) UpdateStoragePolicyID() *GroupUpsert {
	u.SetExcluded(group.FieldStoragePolicyID)
	return u
}

// ClearStoragePolicyID clears the value of the "storage_policy_id" field.
func (u *GroupUpsert) ClearStoragePolicyID() *GroupUpsert {
	u.SetNull(group.FieldStoragePolicyID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(group.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupUpsertOne) Ignore() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertOne) DoNothing() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreate.OnConflict
// documentation for more info.
func (u *GroupUpsertOne) Update(set func(*GroupUpsert)) *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GroupUpsertOne) SetUpdatedAt(v time.Time) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateUpdatedAt() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GroupUpsertOne) SetDeletedAt(v time.Time) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateDeletedAt() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GroupUpsertOne) ClearDeletedAt() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *GroupUpsertOne) SetName(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateName() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetMaxStorage sets the "max_storage" field.
func (u *GroupUpsertOne) SetMaxStorage(v int64) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMaxStorage(v)
	})
}

// AddMaxStorage adds v to the "max_storage" field.
func (u *GroupUpsertOne) AddMaxStorage(v int64) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMaxStorage(v)
	})
}

// UpdateMaxStorage sets the "max_storage" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMaxStorage() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMaxStorage()
	})
}

// ClearMaxStorage clears the value of the "max_storage" field.
func (u *GroupUpsertOne) ClearMaxStorage() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearMaxStorage()
	})
}

// SetSpeedLimit sets the "speed_limit" field.
func (u *GroupUpsertOne) SetSpeedLimit(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetSpeedLimit(v)
	})
}

// AddSpeedLimit adds v to the "speed_limit" field.
func (u *GroupUpsertOne) AddSpeedLimit(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddSpeedLimit(v)
	})
}

// UpdateSpeedLimit sets the "speed_limit" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateSpeedLimit() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSpeedLimit()
	})
}

// ClearSpeedLimit clears the value of the "speed_limit" field.
func (u *GroupUpsertOne) ClearSpeedLimit() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearSpeedLimit()
	})
}

// SetPermissions sets the "permissions" field.
func (u *GroupUpsertOne) SetPermissions(v *boolset.BooleanSet) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetPermissions(v)
	})
}

// UpdatePermissions sets the "permissions" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdatePermissions() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePermissions()
	})
}

// SetSettings sets the "settings" field.
func (u *GroupUpsertOne) SetSettings(v *types.GroupSetting) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetSettings(v)
	})
}

// UpdateSettings sets the "settings" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateSettings() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSettings()
	})
}

// ClearSettings clears the value of the "settings" field.
func (u *GroupUpsertOne) ClearSettings() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearSettings()
	})
}

// SetStoragePolicyID sets the "storage_policy_id" field.
func (u *GroupUpsertOne) SetStoragePolicyID(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetStoragePolicyID(v)
	})
}

// UpdateStoragePolicyID sets the "storage_policy_id" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateStoragePolicyID() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateStoragePolicyID()
	})
}

// ClearStoragePolicyID clears the value of the "storage_policy_id" field.
func (u *GroupUpsertOne) ClearStoragePolicyID() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearStoragePolicyID()
	})
}

// Exec executes the query.
func (u *GroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (m *GroupCreate) SetRawID(t int) *GroupCreate {
	m.mutation.SetRawID(t)
	return m
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	gcb.conflict = opts
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// GroupUpsertBulk is the builder for "upsert"-ing
// a bulk of Group nodes.
type GroupUpsertBulk struct {
	create *GroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertBulk) UpdateNewValues() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(group.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupUpsertBulk) Ignore() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertBulk) DoNothing() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreateBulk.OnConflict
// documentation for more info.
func (u *GroupUpsertBulk) Update(set func(*GroupUpsert)) *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GroupUpsertBulk) SetUpdatedAt(v time.Time) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateUpdatedAt() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GroupUpsertBulk) SetDeletedAt(v time.Time) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateDeletedAt() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *GroupUpsertBulk) ClearDeletedAt() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *GroupUpsertBulk) SetName(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateName() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetMaxStorage sets the "max_storage" field.
func (u *GroupUpsertBulk) SetMaxStorage(v int64) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMaxStorage(v)
	})
}

// AddMaxStorage adds v to the "max_storage" field.
func (u *GroupUpsertBulk) AddMaxStorage(v int64) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMaxStorage(v)
	})
}

// UpdateMaxStorage sets the "max_storage" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMaxStorage() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMaxStorage()
	})
}

// ClearMaxStorage clears the value of the "max_storage" field.
func (u *GroupUpsertBulk) ClearMaxStorage() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearMaxStorage()
	})
}

// SetSpeedLimit sets the "speed_limit" field.
func (u *GroupUpsertBulk) SetSpeedLimit(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetSpeedLimit(v)
	})
}

// AddSpeedLimit adds v to the "speed_limit" field.
func (u *GroupUpsertBulk) AddSpeedLimit(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddSpeedLimit(v)
	})
}

// UpdateSpeedLimit sets the "speed_limit" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateSpeedLimit() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSpeedLimit()
	})
}

// ClearSpeedLimit clears the value of the "speed_limit" field.
func (u *GroupUpsertBulk) ClearSpeedLimit() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearSpeedLimit()
	})
}

// SetPermissions sets the "permissions" field.
func (u *GroupUpsertBulk) SetPermissions(v *boolset.BooleanSet) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetPermissions(v)
	})
}

// UpdatePermissions sets the "permissions" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdatePermissions() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePermissions()
	})
}

// SetSettings sets the "settings" field.
func (u *GroupUpsertBulk) SetSettings(v *types.GroupSetting) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetSettings(v)
	})
}

// UpdateSettings sets the "settings" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateSettings() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSettings()
	})
}

// ClearSettings clears the value of the "settings" field.
func (u *GroupUpsertBulk) ClearSettings() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearSettings()
	})
}

// SetStoragePolicyID sets the "storage_policy_id" field.
func (u *GroupUpsertBulk) SetStoragePolicyID(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetStoragePolicyID(v)
	})
}

// UpdateStoragePolicyID sets the "storage_policy_id" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateStoragePolicyID() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateStoragePolicyID()
	})
}

// ClearStoragePolicyID clears the value of the "storage_policy_id" field.
func (u *GroupUpsertBulk) ClearStoragePolicyID() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearStoragePolicyID()
	})
}

// Exec executes the query.
func (u *GroupUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
