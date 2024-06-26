// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"us-soccer-go-test/internal/ent/stadium"
	"us-soccer-go-test/internal/ent/weather"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StadiumCreate is the builder for creating a Stadium entity.
type StadiumCreate struct {
	config
	mutation *StadiumMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTeam sets the "team" field.
func (sc *StadiumCreate) SetTeam(s string) *StadiumCreate {
	sc.mutation.SetTeam(s)
	return sc
}

// SetFdcouk sets the "fdcouk" field.
func (sc *StadiumCreate) SetFdcouk(s string) *StadiumCreate {
	sc.mutation.SetFdcouk(s)
	return sc
}

// SetCity sets the "city" field.
func (sc *StadiumCreate) SetCity(s string) *StadiumCreate {
	sc.mutation.SetCity(s)
	return sc
}

// SetStadium sets the "stadium" field.
func (sc *StadiumCreate) SetStadium(s string) *StadiumCreate {
	sc.mutation.SetStadium(s)
	return sc
}

// SetCapacity sets the "capacity" field.
func (sc *StadiumCreate) SetCapacity(i int) *StadiumCreate {
	sc.mutation.SetCapacity(i)
	return sc
}

// SetLatitude sets the "latitude" field.
func (sc *StadiumCreate) SetLatitude(f float64) *StadiumCreate {
	sc.mutation.SetLatitude(f)
	return sc
}

// SetLongitude sets the "longitude" field.
func (sc *StadiumCreate) SetLongitude(f float64) *StadiumCreate {
	sc.mutation.SetLongitude(f)
	return sc
}

// SetCountry sets the "country" field.
func (sc *StadiumCreate) SetCountry(s string) *StadiumCreate {
	sc.mutation.SetCountry(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StadiumCreate) SetID(u uuid.UUID) *StadiumCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StadiumCreate) SetNillableID(u *uuid.UUID) *StadiumCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetWeatherID sets the "weather" edge to the Weather entity by ID.
func (sc *StadiumCreate) SetWeatherID(id uuid.UUID) *StadiumCreate {
	sc.mutation.SetWeatherID(id)
	return sc
}

// SetNillableWeatherID sets the "weather" edge to the Weather entity by ID if the given value is not nil.
func (sc *StadiumCreate) SetNillableWeatherID(id *uuid.UUID) *StadiumCreate {
	if id != nil {
		sc = sc.SetWeatherID(*id)
	}
	return sc
}

// SetWeather sets the "weather" edge to the Weather entity.
func (sc *StadiumCreate) SetWeather(w *Weather) *StadiumCreate {
	return sc.SetWeatherID(w.ID)
}

// Mutation returns the StadiumMutation object of the builder.
func (sc *StadiumCreate) Mutation() *StadiumMutation {
	return sc.mutation
}

// Save creates the Stadium in the database.
func (sc *StadiumCreate) Save(ctx context.Context) (*Stadium, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StadiumCreate) SaveX(ctx context.Context) *Stadium {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StadiumCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StadiumCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StadiumCreate) defaults() error {
	if _, ok := sc.mutation.ID(); !ok {
		if stadium.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized stadium.DefaultID (forgotten import ent/runtime?)")
		}
		v := stadium.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *StadiumCreate) check() error {
	if _, ok := sc.mutation.Team(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`ent: missing required field "Stadium.team"`)}
	}
	if v, ok := sc.mutation.Team(); ok {
		if err := stadium.TeamValidator(v); err != nil {
			return &ValidationError{Name: "team", err: fmt.Errorf(`ent: validator failed for field "Stadium.team": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Fdcouk(); !ok {
		return &ValidationError{Name: "fdcouk", err: errors.New(`ent: missing required field "Stadium.fdcouk"`)}
	}
	if v, ok := sc.mutation.Fdcouk(); ok {
		if err := stadium.FdcoukValidator(v); err != nil {
			return &ValidationError{Name: "fdcouk", err: fmt.Errorf(`ent: validator failed for field "Stadium.fdcouk": %w`, err)}
		}
	}
	if _, ok := sc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "Stadium.city"`)}
	}
	if v, ok := sc.mutation.City(); ok {
		if err := stadium.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "Stadium.city": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Stadium(); !ok {
		return &ValidationError{Name: "stadium", err: errors.New(`ent: missing required field "Stadium.stadium"`)}
	}
	if v, ok := sc.mutation.Stadium(); ok {
		if err := stadium.StadiumValidator(v); err != nil {
			return &ValidationError{Name: "stadium", err: fmt.Errorf(`ent: validator failed for field "Stadium.stadium": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "Stadium.capacity"`)}
	}
	if _, ok := sc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Stadium.latitude"`)}
	}
	if _, ok := sc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Stadium.longitude"`)}
	}
	if _, ok := sc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "Stadium.country"`)}
	}
	if v, ok := sc.mutation.Country(); ok {
		if err := stadium.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "Stadium.country": %w`, err)}
		}
	}
	return nil
}

func (sc *StadiumCreate) sqlSave(ctx context.Context) (*Stadium, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StadiumCreate) createSpec() (*Stadium, *sqlgraph.CreateSpec) {
	var (
		_node = &Stadium{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stadium.Table, sqlgraph.NewFieldSpec(stadium.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Team(); ok {
		_spec.SetField(stadium.FieldTeam, field.TypeString, value)
		_node.Team = value
	}
	if value, ok := sc.mutation.Fdcouk(); ok {
		_spec.SetField(stadium.FieldFdcouk, field.TypeString, value)
		_node.Fdcouk = value
	}
	if value, ok := sc.mutation.City(); ok {
		_spec.SetField(stadium.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := sc.mutation.Stadium(); ok {
		_spec.SetField(stadium.FieldStadium, field.TypeString, value)
		_node.Stadium = value
	}
	if value, ok := sc.mutation.Capacity(); ok {
		_spec.SetField(stadium.FieldCapacity, field.TypeInt, value)
		_node.Capacity = value
	}
	if value, ok := sc.mutation.Latitude(); ok {
		_spec.SetField(stadium.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := sc.mutation.Longitude(); ok {
		_spec.SetField(stadium.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if value, ok := sc.mutation.Country(); ok {
		_spec.SetField(stadium.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if nodes := sc.mutation.WeatherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   stadium.WeatherTable,
			Columns: []string{stadium.WeatherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weather.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stadium.Create().
//		SetTeam(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StadiumUpsert) {
//			SetTeam(v+v).
//		}).
//		Exec(ctx)
func (sc *StadiumCreate) OnConflict(opts ...sql.ConflictOption) *StadiumUpsertOne {
	sc.conflict = opts
	return &StadiumUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stadium.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *StadiumCreate) OnConflictColumns(columns ...string) *StadiumUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StadiumUpsertOne{
		create: sc,
	}
}

type (
	// StadiumUpsertOne is the builder for "upsert"-ing
	//  one Stadium node.
	StadiumUpsertOne struct {
		create *StadiumCreate
	}

	// StadiumUpsert is the "OnConflict" setter.
	StadiumUpsert struct {
		*sql.UpdateSet
	}
)

// SetTeam sets the "team" field.
func (u *StadiumUpsert) SetTeam(v string) *StadiumUpsert {
	u.Set(stadium.FieldTeam, v)
	return u
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateTeam() *StadiumUpsert {
	u.SetExcluded(stadium.FieldTeam)
	return u
}

// SetFdcouk sets the "fdcouk" field.
func (u *StadiumUpsert) SetFdcouk(v string) *StadiumUpsert {
	u.Set(stadium.FieldFdcouk, v)
	return u
}

// UpdateFdcouk sets the "fdcouk" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateFdcouk() *StadiumUpsert {
	u.SetExcluded(stadium.FieldFdcouk)
	return u
}

// SetCity sets the "city" field.
func (u *StadiumUpsert) SetCity(v string) *StadiumUpsert {
	u.Set(stadium.FieldCity, v)
	return u
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateCity() *StadiumUpsert {
	u.SetExcluded(stadium.FieldCity)
	return u
}

// SetStadium sets the "stadium" field.
func (u *StadiumUpsert) SetStadium(v string) *StadiumUpsert {
	u.Set(stadium.FieldStadium, v)
	return u
}

// UpdateStadium sets the "stadium" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateStadium() *StadiumUpsert {
	u.SetExcluded(stadium.FieldStadium)
	return u
}

// SetCapacity sets the "capacity" field.
func (u *StadiumUpsert) SetCapacity(v int) *StadiumUpsert {
	u.Set(stadium.FieldCapacity, v)
	return u
}

// UpdateCapacity sets the "capacity" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateCapacity() *StadiumUpsert {
	u.SetExcluded(stadium.FieldCapacity)
	return u
}

// AddCapacity adds v to the "capacity" field.
func (u *StadiumUpsert) AddCapacity(v int) *StadiumUpsert {
	u.Add(stadium.FieldCapacity, v)
	return u
}

// SetLatitude sets the "latitude" field.
func (u *StadiumUpsert) SetLatitude(v float64) *StadiumUpsert {
	u.Set(stadium.FieldLatitude, v)
	return u
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateLatitude() *StadiumUpsert {
	u.SetExcluded(stadium.FieldLatitude)
	return u
}

// AddLatitude adds v to the "latitude" field.
func (u *StadiumUpsert) AddLatitude(v float64) *StadiumUpsert {
	u.Add(stadium.FieldLatitude, v)
	return u
}

// SetLongitude sets the "longitude" field.
func (u *StadiumUpsert) SetLongitude(v float64) *StadiumUpsert {
	u.Set(stadium.FieldLongitude, v)
	return u
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateLongitude() *StadiumUpsert {
	u.SetExcluded(stadium.FieldLongitude)
	return u
}

// AddLongitude adds v to the "longitude" field.
func (u *StadiumUpsert) AddLongitude(v float64) *StadiumUpsert {
	u.Add(stadium.FieldLongitude, v)
	return u
}

// SetCountry sets the "country" field.
func (u *StadiumUpsert) SetCountry(v string) *StadiumUpsert {
	u.Set(stadium.FieldCountry, v)
	return u
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *StadiumUpsert) UpdateCountry() *StadiumUpsert {
	u.SetExcluded(stadium.FieldCountry)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Stadium.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stadium.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StadiumUpsertOne) UpdateNewValues() *StadiumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(stadium.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Stadium.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StadiumUpsertOne) Ignore() *StadiumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StadiumUpsertOne) DoNothing() *StadiumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StadiumCreate.OnConflict
// documentation for more info.
func (u *StadiumUpsertOne) Update(set func(*StadiumUpsert)) *StadiumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StadiumUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeam sets the "team" field.
func (u *StadiumUpsertOne) SetTeam(v string) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetTeam(v)
	})
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateTeam() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateTeam()
	})
}

// SetFdcouk sets the "fdcouk" field.
func (u *StadiumUpsertOne) SetFdcouk(v string) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetFdcouk(v)
	})
}

// UpdateFdcouk sets the "fdcouk" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateFdcouk() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateFdcouk()
	})
}

// SetCity sets the "city" field.
func (u *StadiumUpsertOne) SetCity(v string) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCity(v)
	})
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateCity() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCity()
	})
}

// SetStadium sets the "stadium" field.
func (u *StadiumUpsertOne) SetStadium(v string) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetStadium(v)
	})
}

// UpdateStadium sets the "stadium" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateStadium() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateStadium()
	})
}

// SetCapacity sets the "capacity" field.
func (u *StadiumUpsertOne) SetCapacity(v int) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCapacity(v)
	})
}

// AddCapacity adds v to the "capacity" field.
func (u *StadiumUpsertOne) AddCapacity(v int) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.AddCapacity(v)
	})
}

// UpdateCapacity sets the "capacity" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateCapacity() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCapacity()
	})
}

// SetLatitude sets the "latitude" field.
func (u *StadiumUpsertOne) SetLatitude(v float64) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *StadiumUpsertOne) AddLatitude(v float64) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateLatitude() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *StadiumUpsertOne) SetLongitude(v float64) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *StadiumUpsertOne) AddLongitude(v float64) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateLongitude() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateLongitude()
	})
}

// SetCountry sets the "country" field.
func (u *StadiumUpsertOne) SetCountry(v string) *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *StadiumUpsertOne) UpdateCountry() *StadiumUpsertOne {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCountry()
	})
}

// Exec executes the query.
func (u *StadiumUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StadiumCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StadiumUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StadiumUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: StadiumUpsertOne.ID is not supported by MySQL driver. Use StadiumUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StadiumUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StadiumCreateBulk is the builder for creating many Stadium entities in bulk.
type StadiumCreateBulk struct {
	config
	err      error
	builders []*StadiumCreate
	conflict []sql.ConflictOption
}

// Save creates the Stadium entities in the database.
func (scb *StadiumCreateBulk) Save(ctx context.Context) ([]*Stadium, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stadium, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StadiumMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StadiumCreateBulk) SaveX(ctx context.Context) []*Stadium {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StadiumCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StadiumCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stadium.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StadiumUpsert) {
//			SetTeam(v+v).
//		}).
//		Exec(ctx)
func (scb *StadiumCreateBulk) OnConflict(opts ...sql.ConflictOption) *StadiumUpsertBulk {
	scb.conflict = opts
	return &StadiumUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stadium.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *StadiumCreateBulk) OnConflictColumns(columns ...string) *StadiumUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StadiumUpsertBulk{
		create: scb,
	}
}

// StadiumUpsertBulk is the builder for "upsert"-ing
// a bulk of Stadium nodes.
type StadiumUpsertBulk struct {
	create *StadiumCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Stadium.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stadium.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StadiumUpsertBulk) UpdateNewValues() *StadiumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(stadium.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Stadium.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StadiumUpsertBulk) Ignore() *StadiumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StadiumUpsertBulk) DoNothing() *StadiumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StadiumCreateBulk.OnConflict
// documentation for more info.
func (u *StadiumUpsertBulk) Update(set func(*StadiumUpsert)) *StadiumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StadiumUpsert{UpdateSet: update})
	}))
	return u
}

// SetTeam sets the "team" field.
func (u *StadiumUpsertBulk) SetTeam(v string) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetTeam(v)
	})
}

// UpdateTeam sets the "team" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateTeam() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateTeam()
	})
}

// SetFdcouk sets the "fdcouk" field.
func (u *StadiumUpsertBulk) SetFdcouk(v string) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetFdcouk(v)
	})
}

// UpdateFdcouk sets the "fdcouk" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateFdcouk() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateFdcouk()
	})
}

// SetCity sets the "city" field.
func (u *StadiumUpsertBulk) SetCity(v string) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCity(v)
	})
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateCity() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCity()
	})
}

// SetStadium sets the "stadium" field.
func (u *StadiumUpsertBulk) SetStadium(v string) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetStadium(v)
	})
}

// UpdateStadium sets the "stadium" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateStadium() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateStadium()
	})
}

// SetCapacity sets the "capacity" field.
func (u *StadiumUpsertBulk) SetCapacity(v int) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCapacity(v)
	})
}

// AddCapacity adds v to the "capacity" field.
func (u *StadiumUpsertBulk) AddCapacity(v int) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.AddCapacity(v)
	})
}

// UpdateCapacity sets the "capacity" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateCapacity() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCapacity()
	})
}

// SetLatitude sets the "latitude" field.
func (u *StadiumUpsertBulk) SetLatitude(v float64) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *StadiumUpsertBulk) AddLatitude(v float64) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateLatitude() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *StadiumUpsertBulk) SetLongitude(v float64) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *StadiumUpsertBulk) AddLongitude(v float64) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateLongitude() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateLongitude()
	})
}

// SetCountry sets the "country" field.
func (u *StadiumUpsertBulk) SetCountry(v string) *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *StadiumUpsertBulk) UpdateCountry() *StadiumUpsertBulk {
	return u.Update(func(s *StadiumUpsert) {
		s.UpdateCountry()
	})
}

// Exec executes the query.
func (u *StadiumUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StadiumCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StadiumCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StadiumUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
