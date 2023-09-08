// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khanakia/entautoslow/ent/predicate"
	"github.com/khanakia/entautoslow/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetAPIKey sets the "api_key" field.
func (uu *UserUpdate) SetAPIKey(s string) *UserUpdate {
	uu.mutation.SetAPIKey(s)
	return uu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAPIKey(s *string) *UserUpdate {
	if s != nil {
		uu.SetAPIKey(*s)
	}
	return uu
}

// ClearAPIKey clears the value of the "api_key" field.
func (uu *UserUpdate) ClearAPIKey() *UserUpdate {
	uu.mutation.ClearAPIKey()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(b bool) *UserUpdate {
	uu.mutation.SetStatus(b)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(b *bool) *UserUpdate {
	if b != nil {
		uu.SetStatus(*b)
	}
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.APIKey(); ok {
		_spec.SetField(user.FieldAPIKey, field.TypeString, value)
	}
	if uu.mutation.APIKeyCleared() {
		_spec.ClearField(user.FieldAPIKey, field.TypeString)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeBool, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeBool)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetAPIKey sets the "api_key" field.
func (uuo *UserUpdateOne) SetAPIKey(s string) *UserUpdateOne {
	uuo.mutation.SetAPIKey(s)
	return uuo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAPIKey(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAPIKey(*s)
	}
	return uuo
}

// ClearAPIKey clears the value of the "api_key" field.
func (uuo *UserUpdateOne) ClearAPIKey() *UserUpdateOne {
	uuo.mutation.ClearAPIKey()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(b bool) *UserUpdateOne {
	uuo.mutation.SetStatus(b)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetStatus(*b)
	}
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.APIKey(); ok {
		_spec.SetField(user.FieldAPIKey, field.TypeString, value)
	}
	if uuo.mutation.APIKeyCleared() {
		_spec.ClearField(user.FieldAPIKey, field.TypeString)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeBool, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeBool)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
