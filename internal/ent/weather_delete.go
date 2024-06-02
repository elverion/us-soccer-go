// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"us-soccer-go-test/internal/ent/predicate"
	"us-soccer-go-test/internal/ent/weather"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WeatherDelete is the builder for deleting a Weather entity.
type WeatherDelete struct {
	config
	hooks    []Hook
	mutation *WeatherMutation
}

// Where appends a list predicates to the WeatherDelete builder.
func (wd *WeatherDelete) Where(ps ...predicate.Weather) *WeatherDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WeatherDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WeatherDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WeatherDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(weather.Table, sqlgraph.NewFieldSpec(weather.FieldID, field.TypeUUID))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WeatherDeleteOne is the builder for deleting a single Weather entity.
type WeatherDeleteOne struct {
	wd *WeatherDelete
}

// Where appends a list predicates to the WeatherDelete builder.
func (wdo *WeatherDeleteOne) Where(ps ...predicate.Weather) *WeatherDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WeatherDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{weather.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WeatherDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}