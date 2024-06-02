// Code generated by ent, DO NOT EDIT.

package weather

import (
	"time"
	"us-soccer-go-test/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Weather {
	return predicate.Weather(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldUpdateTime, v))
}

// Temperature applies equality check predicate on the "temperature" field. It's identical to TemperatureEQ.
func Temperature(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldTemperature, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldDescription, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Weather {
	return predicate.Weather(sql.FieldLTE(FieldUpdateTime, v))
}

// TemperatureEQ applies the EQ predicate on the "temperature" field.
func TemperatureEQ(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldTemperature, v))
}

// TemperatureNEQ applies the NEQ predicate on the "temperature" field.
func TemperatureNEQ(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldNEQ(FieldTemperature, v))
}

// TemperatureIn applies the In predicate on the "temperature" field.
func TemperatureIn(vs ...float64) predicate.Weather {
	return predicate.Weather(sql.FieldIn(FieldTemperature, vs...))
}

// TemperatureNotIn applies the NotIn predicate on the "temperature" field.
func TemperatureNotIn(vs ...float64) predicate.Weather {
	return predicate.Weather(sql.FieldNotIn(FieldTemperature, vs...))
}

// TemperatureGT applies the GT predicate on the "temperature" field.
func TemperatureGT(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldGT(FieldTemperature, v))
}

// TemperatureGTE applies the GTE predicate on the "temperature" field.
func TemperatureGTE(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldGTE(FieldTemperature, v))
}

// TemperatureLT applies the LT predicate on the "temperature" field.
func TemperatureLT(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldLT(FieldTemperature, v))
}

// TemperatureLTE applies the LTE predicate on the "temperature" field.
func TemperatureLTE(v float64) predicate.Weather {
	return predicate.Weather(sql.FieldLTE(FieldTemperature, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Weather {
	return predicate.Weather(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Weather {
	return predicate.Weather(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Weather {
	return predicate.Weather(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Weather {
	return predicate.Weather(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Weather {
	return predicate.Weather(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Weather {
	return predicate.Weather(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Weather {
	return predicate.Weather(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Weather {
	return predicate.Weather(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Weather {
	return predicate.Weather(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Weather {
	return predicate.Weather(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Weather {
	return predicate.Weather(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Weather {
	return predicate.Weather(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Weather {
	return predicate.Weather(sql.FieldContainsFold(FieldDescription, v))
}

// HasStadium applies the HasEdge predicate on the "stadium" edge.
func HasStadium() predicate.Weather {
	return predicate.Weather(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StadiumTable, StadiumColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStadiumWith applies the HasEdge predicate on the "stadium" edge with a given conditions (other predicates).
func HasStadiumWith(preds ...predicate.Stadium) predicate.Weather {
	return predicate.Weather(func(s *sql.Selector) {
		step := newStadiumStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Weather) predicate.Weather {
	return predicate.Weather(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Weather) predicate.Weather {
	return predicate.Weather(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Weather) predicate.Weather {
	return predicate.Weather(sql.NotPredicates(p))
}