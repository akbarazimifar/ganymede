// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/twitchcategory"
)

// TwitchCategoryDelete is the builder for deleting a TwitchCategory entity.
type TwitchCategoryDelete struct {
	config
	hooks    []Hook
	mutation *TwitchCategoryMutation
}

// Where appends a list predicates to the TwitchCategoryDelete builder.
func (tcd *TwitchCategoryDelete) Where(ps ...predicate.TwitchCategory) *TwitchCategoryDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TwitchCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TwitchCategoryMutation](ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TwitchCategoryDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TwitchCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(twitchcategory.Table, sqlgraph.NewFieldSpec(twitchcategory.FieldID, field.TypeString))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TwitchCategoryDeleteOne is the builder for deleting a single TwitchCategory entity.
type TwitchCategoryDeleteOne struct {
	tcd *TwitchCategoryDelete
}

// Where appends a list predicates to the TwitchCategoryDelete builder.
func (tcdo *TwitchCategoryDeleteOne) Where(ps ...predicate.TwitchCategory) *TwitchCategoryDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TwitchCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{twitchcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TwitchCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
